// Package notion implements a plugin for reading data from Notion.
package notion

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cmskitdev/client"
	"github.com/cmskitdev/common"
	"github.com/cmskitdev/engine"
	"github.com/cmskitdev/notion/types"
	"github.com/mateothegreat/go-multilog/multilog"
)

// Plugin implements DataSource for reading data from Notion
type Plugin struct {
	client  *client.Client
	config  NotionSourceConfig
	metrics *NotionSourceMetrics
	mu      sync.RWMutex
}

// NewNotionSource creates a new Notion data source
func NewNotionSource(client *client.Client, config NotionSourceConfig) *Plugin {
	return &Plugin{
		client: client,
		config: config,
		metrics: &NotionSourceMetrics{
			StartTime: time.Now(),
		},
	}
}

func (ns *Plugin) Read(ctx context.Context, req *engine.ReadRequest) (<-chan engine.DataItemContainer[any], error) {
	// Validate request
	if err := ns.Validate(req); err != nil {
		return nil, fmt.Errorf("invalid read request: %w", err)
	}

	// Make a larger buffer to handle concurrent processing burst
	bufferSize := ns.config.PageSize * ns.config.MaxConcurrent * 10
	results := make(chan engine.DataItemContainer[any], bufferSize)

	go func() {
		defer ns.updateEndTime()
		defer close(results)

		var wg sync.WaitGroup

		for _, objType := range req.Types {
			switch objType {
			case common.ObjectTypePage:
				if ns.config.IncludePages {
					wg.Add(1)
					go func() {
						defer wg.Done()
						ns.searchPages(ctx, results)
					}()
				}
			case common.ObjectTypeCollection:
				if ns.config.IncludeCollections {
					wg.Add(1)
					go func() {
						defer wg.Done()
						ns.readDatabases(ctx, results)
					}()
				}
			}
		}

		wg.Wait()
	}()

	return results, nil
}

// Validate implements DataSource.Validate
func (ns *Plugin) Validate(req *engine.ReadRequest) error {
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
func (ns *Plugin) SupportsType(objType common.ObjectType) bool {
	switch objType {
	case common.ObjectTypePage, common.ObjectTypeCollection, common.ObjectTypeBlock,
		common.ObjectTypeComment, common.ObjectTypeUser:
		return true
	default:
		return false
	}
}

// Config implements DataSource.Config
func (ns *Plugin) Config() engine.SourceConfig {
	return engine.SourceConfig{
		Type: "notion",
		Name: "Notion API Source",
		Properties: map[string]interface{}{
			"include_databases":   ns.config.IncludeCollections,
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
func (ns *Plugin) Close() error {
	ns.updateEndTime()
	return nil
}

// searchPages performs paginated search for pages with streaming.
//
// Arguments:
// - ctx: The context for the request.
// - results: The channel to send the results to.
//
// Returns:
// - The processed pages.
func (ns *Plugin) searchPages(ctx context.Context, results chan<- engine.DataItemContainer[any]) {
	searchReq := types.SearchRequest{
		Query: ns.config.SearchQuery,
		Filter: &types.SearchFilter{
			Value:    "page",
			Property: "object",
		},
		PageSize: &ns.config.PageSize,
	}

	multilog.Debug("notion.searchPages", "searchReq", map[string]interface{}{
		"searchReq": searchReq,
	})

	ns.incrementRequestCount()
	// Use streaming query for pagination.
	searchResults := ns.client.Registry.Search().Stream(ctx, searchReq)

	// Process paginated results as they stream in.
	for result := range searchResults {
		if result.IsError() {
			ns.incrementErrorCount()
			continue
		}

		if result.Data.Page != nil {
			item := ns.convertPageDataToDataItem(result.Data.Page)
			select {
			case results <- *item:
				ns.incrementPageCount()
			case <-ctx.Done():
				return
			}
		}
	}
}

// getPage processes a single page.
//
// Arguments:
// - ctx: The context for the request.
// - page: The page to process.
// - results: The channel to send the results to.
//
// Returns:
// - The processed page.
func (ns *Plugin) getPage(ctx context.Context, page *types.Page, results chan<- engine.DataItemContainer[any]) {
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
func (ns *Plugin) readDatabases(ctx context.Context, results chan<- engine.DataItemContainer[any]) {
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

// convertPageDataToDataItem converts a basic page from search results to a data item
func (ns *Plugin) convertPageDataToDataItem(page *types.Page) *engine.DataItemContainer[any] {
	now := time.Now()
	return &engine.DataItemContainer[any]{
		ID:   string(page.ID),
		Type: common.ObjectTypePage,
		Data: page,
		Metadata: engine.ItemMetadata{
			SourceType:  "notion",
			SourceID:    string(page.ID),
			OriginalID:  string(page.ID),
			CreatedAt:   page.CreatedTime,
			ModifiedAt:  page.LastEditedTime,
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
				"archived":   page.Archived,
				"created_by": page.CreatedBy,
				"edited_by":  page.LastEditedBy,
			},
		},
	}
}

func (ns *Plugin) convertPageToDataItem(page *client.GetPageResult) *engine.DataItemContainer[any] {
	now := time.Now()
	return &engine.DataItemContainer[any]{
		ID:   string(page.Page.ID),
		Type: common.ObjectTypePage,
		Data: page,
		Metadata: engine.ItemMetadata{
			SourceType:  "notion",
			SourceID:    string(page.Page.ID),
			OriginalID:  string(page.Page.ID),
			CreatedAt:   page.Page.CreatedTime,
			ModifiedAt:  page.Page.LastEditedTime,
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

func (ns *Plugin) convertBlockToDataItem(block *types.Block, parentID string) *engine.DataItemContainer[any] {
	now := time.Now()
	return &engine.DataItemContainer[any]{
		ID:   string(block.ID),
		Type: common.ObjectTypeBlock,
		Data: block,
		Metadata: engine.ItemMetadata{
			SourceType:  "notion",
			SourceID:    string(block.ID),
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

func (ns *Plugin) convertDatabaseToDataItem(database *types.Database) *engine.DataItemContainer[any] {
	now := time.Now()
	return &engine.DataItemContainer[any]{
		ID:   string(database.ID),
		Type: common.ObjectTypeCollection,
		Data: database,
		Metadata: engine.ItemMetadata{
			SourceType:  "notion",
			SourceID:    string(database.ID),
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
