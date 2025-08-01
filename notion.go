package notion

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/notioncodes/client"
	"github.com/notioncodes/engine"
	"github.com/notioncodes/plugin-notion/types"
)

// NotionSource implements DataSource for reading data from Notion
type NotionSource struct {
	client  *client.Client
	config  NotionSourceConfig
	metrics *NotionSourceMetrics
	mu      sync.RWMutex
}

// NotionSourceConfig configures the Notion data source
type NotionSourceConfig struct {
	// Object types to read
	IncludeDatabases bool `json:"include_databases"`
	IncludePages     bool `json:"include_pages"`
	IncludeBlocks    bool `json:"include_blocks"`
	IncludeComments  bool `json:"include_comments"`
	IncludeUsers     bool `json:"include_users"`

	// Depth control for hierarchical objects
	MaxDepth int `json:"max_depth"`

	// Pagination and performance
	PageSize          int     `json:"page_size"`
	RequestsPerSecond float64 `json:"requests_per_second"`
	MaxConcurrent     int     `json:"max_concurrent"`

	// Pagination-specific configuration
	EnablePagination   bool          `json:"enable_pagination"`
	MaxPagesPerQuery   int           `json:"max_pages_per_query"`
	PaginationTimeout  time.Duration `json:"pagination_timeout"`
	StreamingBatchSize int           `json:"streaming_batch_size"`

	// Filtering
	DatabaseIDs []string           `json:"database_ids,omitempty"`
	PageIDs     []string           `json:"page_ids,omitempty"`
	Filters     []NotionTypeFilter `json:"filters,omitempty"`
	SearchQuery string             `json:"search_query,omitempty"`
	DateRange   *NotionDateRange   `json:"date_range,omitempty"`
	Properties  []string           `json:"properties,omitempty"` // Specific properties to include
}

// NotionTypeFilter filters objects by type and properties
type NotionTypeFilter struct {
	ObjectType engine.ObjectType `json:"object_type"`
	Conditions []FilterCondition `json:"conditions"`
}

// FilterCondition represents a single filter condition
type FilterCondition struct {
	Property string      `json:"property"`
	Operator string      `json:"operator"` // equals, contains, starts_with, etc.
	Value    interface{} `json:"value"`
}

// NotionDateRange filters objects by creation/modification date
type NotionDateRange struct {
	Start    *time.Time `json:"start,omitempty"`
	End      *time.Time `json:"end,omitempty"`
	DateType string     `json:"date_type"` // created_time, last_edited_time
}

// NotionSourceMetrics tracks source performance
type NotionSourceMetrics struct {
	ObjectsRead       int64         `json:"objects_read"`
	PagesRead         int64         `json:"pages_read"`
	BlocksRead        int64         `json:"blocks_read"`
	CommentsRead      int64         `json:"comments_read"`
	DatabasesRead     int64         `json:"databases_read"`
	UsersRead         int64         `json:"users_read"`
	RequestsMade      int64         `json:"requests_made"`
	ErrorsEncountered int64         `json:"errors_encountered"`
	TotalDuration     time.Duration `json:"total_duration"`
	StartTime         time.Time     `json:"start_time"`
	EndTime           *time.Time    `json:"end_time,omitempty"`
	mu                sync.RWMutex
}

// DefaultNotionSourceConfig returns sensible defaults
func DefaultNotionSourceConfig() NotionSourceConfig {
	return NotionSourceConfig{
		IncludeDatabases:   true,
		IncludePages:       true,
		IncludeBlocks:      true,
		IncludeComments:    false, // Comments are expensive to fetch
		IncludeUsers:       false, // Users are typically not needed in exports
		MaxDepth:           3,
		PageSize:           100,
		RequestsPerSecond:  3.0, // Notion API rate limit
		MaxConcurrent:      5,
		EnablePagination:   true,
		MaxPagesPerQuery:   100, // Reasonable limit
		PaginationTimeout:  5 * time.Minute,
		StreamingBatchSize: 10, // Process 10 items at a time
		DatabaseIDs:        []string{},
		PageIDs:            []string{},
		Filters:            []NotionTypeFilter{},
		SearchQuery:        "",
		DateRange:          nil,
		Properties:         []string{},
	}
}

// NewNotionSource creates a new Notion data source
func NewNotionSource(client *client.Client, config NotionSourceConfig) *NotionSource {
	return &NotionSource{
		client: client,
		config: config,
		metrics: &NotionSourceMetrics{
			StartTime: time.Now(),
		},
	}
}

func (ns *NotionSource) Read(ctx context.Context, req *engine.ReadRequest) (<-chan engine.DataItemContainer, error) {
	// Validate request
	if err := ns.Validate(req); err != nil {
		return nil, fmt.Errorf("invalid read request: %w", err)
	}

	// Make a larger buffer to handle concurrent processing burst
	bufferSize := ns.config.PageSize * ns.config.MaxConcurrent * 10
	results := make(chan engine.DataItemContainer, bufferSize)

	go func() {
		defer close(results)
		defer ns.updateEndTime()
		ns.readData(ctx, req, results)
	}()

	return results, nil
}

// Validate implements DataSource.Validate
func (ns *NotionSource) Validate(req *engine.ReadRequest) error {
	if ns.client == nil {
		return fmt.Errorf("notion client is required")
	}

	// Validate that requested types are supported
	for _, objType := range req.Types {
		if !ns.SupportsType(objType) {
			return fmt.Errorf("object type %s is not supported by Notion source", objType)
		}
	}

	return nil
}

// SupportsType implements DataSource.SupportsType
func (ns *NotionSource) SupportsType(objType engine.ObjectType) bool {
	switch objType {
	case engine.ObjectTypePage, engine.ObjectTypeDatabase, engine.ObjectTypeBlock,
		engine.ObjectTypeComment, engine.ObjectTypeUser:
		return true
	default:
		return false
	}
}

// Config implements DataSource.Config
func (ns *NotionSource) Config() engine.SourceConfig {
	return engine.SourceConfig{
		Type: "notion",
		Name: "Notion API Source",
		Properties: map[string]interface{}{
			"include_databases":   ns.config.IncludeDatabases,
			"include_pages":       ns.config.IncludePages,
			"include_blocks":      ns.config.IncludeBlocks,
			"include_comments":    ns.config.IncludeComments,
			"include_users":       ns.config.IncludeUsers,
			"max_depth":           ns.config.MaxDepth,
			"page_size":           ns.config.PageSize,
			"requests_per_second": ns.config.RequestsPerSecond,
			"max_concurrent":      ns.config.MaxConcurrent,
		},
	}
}

// Close implements DataSource.Close
func (ns *NotionSource) Close() error {
	ns.updateEndTime()
	return nil
}

// Helper methods for data processing

// readData orchestrates the data reading process
func (ns *NotionSource) readData(ctx context.Context, req *engine.ReadRequest, results chan<- engine.DataItemContainer) {
	var wg sync.WaitGroup

	// Handle each object type requested
	for _, objType := range req.Types {
		switch objType {
		case engine.ObjectTypePage:
			if ns.config.IncludePages {
				wg.Add(1)
				go func() {
					defer wg.Done()
					ns.searchPages(ctx, results)
				}()
			}
		case engine.ObjectTypeDatabase:
			if ns.config.IncludeDatabases {
				wg.Add(1)
				go func() {
					defer wg.Done()
					ns.readDatabases(ctx, results)
				}()
			}
		}
	}

	wg.Wait()
}

// searchPages performs paginated search for pages with streaming
func (ns *NotionSource) searchPages(ctx context.Context, results chan<- engine.DataItemContainer) {
	searchReq := types.SearchRequest{
		Query: ns.config.SearchQuery,
		Filter: &types.SearchFilter{
			Value:    "page",
			Property: "object",
		},
		PageSize: &ns.config.PageSize,
	}

	ns.incrementRequestCount()
	// Use streaming query for pagination
	searchResults := ns.client.Registry.Search().Stream(ctx, searchReq)

	// Process paginated results as they stream in
	for result := range searchResults {
		if result.IsError() {
			ns.incrementErrorCount()
			continue
		}

		if result.Data.Page != nil {
			// Process each page concurrently
			go ns.processPageConcurrently(ctx, result.Data.Page, results)
		}
	}
}

// processPageConcurrently processes a single page with concurrency control
func (ns *NotionSource) processPageConcurrently(ctx context.Context, page *types.Page, results chan<- engine.DataItemContainer) {
	// Simple concurrency control using a semaphore pattern
	semaphore := make(chan struct{}, ns.config.MaxConcurrent)
	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	select {
	case <-ctx.Done():
		return
	default:
	}

	// Get full page details
	opts := client.GetPageOptions{
		IncludeBlocks:   ns.config.IncludeBlocks,
		IncludeChildren: ns.config.MaxDepth > 1,
		IncludeComments: ns.config.IncludeComments,
	}

	ns.incrementRequestCount()
	result := ns.client.Pages().Get(ctx, page.ID, opts)

	if result.IsError() {
		ns.incrementErrorCount()
		return
	}

	// Convert and send immediately
	item := ns.convertPageToDataItem(&result.Data)

	select {
	case results <- *item:
		ns.incrementPageCount()
	case <-ctx.Done():
		return
	}

	// Process blocks if included
	if ns.config.IncludeBlocks && len(result.Data.Blocks) > 0 {
		for _, block := range result.Data.Blocks {
			blockItem := ns.convertBlockToDataItem(block, string(page.ID))
			select {
			case results <- *blockItem:
				ns.incrementBlockCount()
			case <-ctx.Done():
				return
			}
		}
	}
}

// readDatabases reads databases from the workspace
func (ns *NotionSource) readDatabases(ctx context.Context, results chan<- engine.DataItemContainer) {
	// Search for databases
	searchReq := types.SearchRequest{
		Filter: &types.SearchFilter{
			Value:    "database",
			Property: "object",
		},
		PageSize: &ns.config.PageSize,
	}

	ns.incrementRequestCount()
	searchResults := ns.client.Registry.Search().Stream(ctx, searchReq)

	for result := range searchResults {
		if result.IsError() {
			ns.incrementErrorCount()
			continue
		}

		if result.Data.Database != nil {
			item := ns.convertDatabaseToDataItem(result.Data.Database)
			select {
			case results <- *item:
				ns.incrementDatabaseCount()
			case <-ctx.Done():
				return
			}
		}
	}
}

// Data conversion methods
func (ns *NotionSource) convertPageToDataItem(page *client.GetPageResult) *engine.DataItemContainer {
	now := time.Now()
	return &engine.DataItemContainer{
		ID:   string(page.Page.ID),
		Type: engine.ObjectTypePage,
		Data: page,
		Metadata: engine.ItemMetadata{
			SourceType:  "notion",
			OriginalID:  string(page.Page.ID),
			CreatedAt:   page.Page.CreatedTime.Time,
			ModifiedAt:  page.Page.LastEditedTime.Time,
			ProcessedAt: &now,
			ValidationState: engine.ValidationState{
				IsValid:     false,
				ValidatedAt: now,
			},
			TransformState: engine.TransformState{
				IsTransformed: false,
			},
			ProcessingState: engine.ProcessingState{
				Phase:     engine.PhaseRead,
				Status:    engine.ProcessingStatusPending,
				StartedAt: now,
			},
			Properties: map[string]interface{}{
				"archived":     page.Page.Archived,
				"blocks_count": len(page.Blocks),
				"created_by":   page.Page.CreatedBy,
				"edited_by":    page.Page.LastEditedBy,
			},
		},
	}
}

func (ns *NotionSource) convertBlockToDataItem(block *types.Block, parentID string) *engine.DataItemContainer {
	now := time.Now()
	return &engine.DataItemContainer{
		ID:   string(block.ID),
		Type: engine.ObjectTypeBlock,
		Data: block,
		Metadata: engine.ItemMetadata{
			SourceType:  "notion",
			OriginalID:  string(block.ID),
			CreatedAt:   block.CreatedTime,
			ModifiedAt:  block.LastEditedTime,
			ProcessedAt: &now,
			ValidationState: engine.ValidationState{
				IsValid:     false,
				ValidatedAt: now,
			},
			TransformState: engine.TransformState{
				IsTransformed: false,
			},
			ProcessingState: engine.ProcessingState{
				Phase:     engine.PhaseRead,
				Status:    engine.ProcessingStatusPending,
				StartedAt: now,
			},
			Properties: map[string]interface{}{
				"parent_id":  parentID,
				"created_by": block.CreatedBy,
				"edited_by":  block.LastEditedBy,
			},
		},
	}
}

func (ns *NotionSource) convertDatabaseToDataItem(database *types.Database) *engine.DataItemContainer {
	now := time.Now()
	return &engine.DataItemContainer{
		ID:   string(database.ID),
		Type: engine.ObjectTypeDatabase,
		Data: database,
		Metadata: engine.ItemMetadata{
			SourceType:  "notion",
			OriginalID:  string(database.ID),
			CreatedAt:   database.CreatedTime,
			ModifiedAt:  database.LastEditedTime,
			ProcessedAt: &now,
			ValidationState: engine.ValidationState{
				IsValid:     false,
				ValidatedAt: now,
			},
			TransformState: engine.TransformState{
				IsTransformed: false,
			},
			ProcessingState: engine.ProcessingState{
				Phase:     engine.PhaseRead,
				Status:    engine.ProcessingStatusPending,
				StartedAt: now,
			},
			Properties: map[string]interface{}{
				"archived":   database.Archived,
				"created_by": database.CreatedBy,
				"edited_by":  database.LastEditedBy,
			},
		},
	}
}
