package notion

// import (
// 	"context"
// 	"testing"

// 	"github.com/cmskitdev/client"
// 	"github.com/cmskitdev/test"
// )

// func TestDefaultNotionSourceConfig(t *testing.T) {
// 	config := sources.DefaultNotionSourceConfig()

// 	if !config.IncludeDatabases {
// 		t.Error("Expected IncludeDatabases to be true")
// 	}

// 	if !config.IncludePages {
// 		t.Error("Expected IncludePages to be true")
// 	}

// 	if !config.IncludeBlocks {
// 		t.Error("Expected IncludeBlocks to be true")
// 	}

// 	if !config.IncludeComments {
// 		t.Error("Expected IncludeComments to be true")
// 	}

// 	if config.IncludeUsers {
// 		t.Error("Expected IncludeUsers to be false")
// 	}

// 	if !config.IncludeFiles {
// 		t.Error("Expected IncludeFiles to be true")
// 	}

// 	if config.MaxDepth != 0 {
// 		t.Errorf("Expected MaxDepth to be 0, got %d", config.MaxDepth)
// 	}

// 	if config.IncludeArchived {
// 		t.Error("Expected IncludeArchived to be false")
// 	}

// 	if config.RequestsPerSecond != 3.0 {
// 		t.Errorf("Expected RequestsPerSecond to be 3.0, got %f", config.RequestsPerSecond)
// 	}

// 	if config.PageSize != 100 {
// 		t.Errorf("Expected PageSize to be 100, got %d", config.PageSize)
// 	}
// }

// func TestNewNotionSource(t *testing.T) {
// 	// Create a mock client for testing
// 	clientConfig := client.DefaultConfig()
// 	clientConfig.APIKey = "test-api-key"

// 	notionClient, err := client.NewClient(clientConfig)
// 	if err != nil {
// 		t.Fatalf("Failed to create client: %v", err)
// 	}
// 	defer notionClient.Close()

// 	config := sources.DefaultNotionSourceConfig()
// 	source := sources.NewNotionSource(notionClient, config)

// 	if source == nil {
// 		t.Fatal("Expected source to be created")
// 	}
// }

// func TestNotionSource_Config(t *testing.T) {
// 	// Create a mock client
// 	clientConfig := client.DefaultConfig()
// 	clientConfig.APIKey = "test-api-key"

// 	notionClient, err := client.NewClient(clientConfig)
// 	if err != nil {
// 		t.Fatalf("Failed to create client: %v", err)
// 	}
// 	defer notionClient.Close()

// 	sourceConfig := sources.NotionSourceConfig{
// 		WorkspaceID:       "workspace-123",
// 		IncludeDatabases:  true,
// 		IncludePages:      true,
// 		IncludeBlocks:     false,
// 		IncludeComments:   false,
// 		IncludeUsers:      false,
// 		IncludeFiles:      true,
// 		MaxDepth:          2,
// 		IncludeArchived:   false,
// 		RequestsPerSecond: 2.5,
// 		PageSize:          50,
// 	}

// 	source := sources.NewNotionSource(notionClient, sourceConfig)
// 	config := source.Config()

// 	if config.Type != "notion" {
// 		t.Errorf("Expected Type to be notion, got %s", config.Type)
// 	}

// 	if config.Name != "Notion Data Source" {
// 		t.Errorf("Expected Name to be Notion Data Source, got %s", config.Name)
// 	}

// 	// Test properties
// 	if config.Properties["workspace_id"] != "workspace-123" {
// 		t.Errorf("Expected workspace_id to be workspace-123, got %v", config.Properties["workspace_id"])
// 	}

// 	if config.Properties["include_databases"] != true {
// 		t.Errorf("Expected include_databases to be true, got %v", config.Properties["include_databases"])
// 	}

// 	if config.Properties["include_blocks"] != false {
// 		t.Errorf("Expected include_blocks to be false, got %v", config.Properties["include_blocks"])
// 	}

// 	if config.Properties["max_depth"] != 2 {
// 		t.Errorf("Expected max_depth to be 2, got %v", config.Properties["max_depth"])
// 	}

// 	if config.Properties["requests_per_second"] != 2.5 {
// 		t.Errorf("Expected requests_per_second to be 2.5, got %v", config.Properties["requests_per_second"])
// 	}

// 	if config.Properties["page_size"] != 50 {
// 		t.Errorf("Expected page_size to be 50, got %v", config.Properties["page_size"])
// 	}
// }

// func TestNotionSource_SupportsType(t *testing.T) {
// 	// Create a mock client
// 	clientConfig := client.DefaultConfig()
// 	clientConfig.APIKey = "test-api-key"

// 	notionClient, err := client.NewClient(clientConfig)
// 	if err != nil {
// 		t.Fatalf("Failed to create client: %v", err)
// 	}
// 	defer notionClient.Close()

// 	config := sources.NotionSourceConfig{
// 		IncludeDatabases: true,
// 		IncludePages:     true,
// 		IncludeBlocks:    false,
// 		IncludeComments:  true,
// 		IncludeUsers:     false,
// 		IncludeFiles:     true,
// 	}

// 	source := sources.NewNotionSource(notionClient, config)

// 	// Test supported types
// 	if !source.SupportsType(engine.ObjectTypeDatabase) {
// 		t.Error("Expected to support database type")
// 	}

// 	if !source.SupportsType(engine.ObjectTypePage) {
// 		t.Error("Expected to support page type")
// 	}

// 	if !source.SupportsType(engine.ObjectTypeComment) {
// 		t.Error("Expected to support comment type")
// 	}

// 	if !source.SupportsType(engine.ObjectTypeFile) {
// 		t.Error("Expected to support file type")
// 	}

// 	// Test unsupported types
// 	if source.SupportsType(engine.ObjectTypeBlock) {
// 		t.Error("Expected to not support block type (disabled in config)")
// 	}

// 	if source.SupportsType(engine.ObjectTypeUser) {
// 		t.Error("Expected to not support user type (disabled in config)")
// 	}

// 	if source.SupportsType(engine.ObjectTypeGeneric) {
// 		t.Error("Expected to not support generic type")
// 	}
// }

// func TestNotionSource_Validate(t *testing.T) {
// 	// Create a mock client
// 	clientConfig := client.DefaultConfig()
// 	clientConfig.APIKey = "test-api-key"

// 	notionClient, err := client.NewClient(clientConfig)
// 	if err != nil {
// 		t.Fatalf("Failed to create client: %v", err)
// 	}
// 	defer notionClient.Close()

// 	// Test valid configuration
// 	validConfig := sources.NotionSourceConfig{
// 		IncludeDatabases: true,
// 		IncludePages:     false,
// 		IncludeBlocks:    false,
// 		IncludeComments:  false,
// 		IncludeUsers:     false,
// 		IncludeFiles:     false,
// 	}

// 	source := sources.NewNotionSource(notionClient, validConfig)
// 	err = source.Validate(&engine.ReadRequest{})
// 	if err != nil {
// 		t.Errorf("Expected no error for valid config, got %v", err)
// 	}

// 	// Test invalid configuration - no content types enabled
// 	invalidConfig := sources.NotionSourceConfig{
// 		IncludeDatabases: false,
// 		IncludePages:     false,
// 		IncludeBlocks:    false,
// 		IncludeComments:  false,
// 		IncludeUsers:     false,
// 		IncludeFiles:     false,
// 	}

// 	invalidSource := sources.NewNotionSource(notionClient, invalidConfig)
// 	err = invalidSource.Validate(&engine.ReadRequest{})
// 	if err == nil {
// 		t.Error("Expected error for invalid config with no content types")
// 	}

// 	// Test invalid configuration - block IDs without blocks enabled
// 	blockIDsConfig := sources.NotionSourceConfig{
// 		IncludeDatabases: true,
// 		IncludePages:     true,
// 		IncludeBlocks:    false,                          // Blocks disabled
// 		BlockIDs:         []string{"block-1", "block-2"}, // But block IDs specified
// 	}

// 	blockIDsSource := sources.NewNotionSource(notionClient, blockIDsConfig)
// 	err = blockIDsSource.Validate(&engine.ReadRequest{})
// 	if err == nil {
// 		t.Error("Expected error for block IDs specified but blocks not enabled")
// 	}
// }

// func TestNotionSource_Close(t *testing.T) {
// 	// Create a mock client
// 	clientConfig := client.DefaultConfig()
// 	clientConfig.APIKey = "test-api-key"

// 	notionClient, err := client.NewClient(clientConfig)
// 	if err != nil {
// 		t.Fatalf("Failed to create client: %v", err)
// 	}
// 	defer notionClient.Close()

// 	config := sources.DefaultNotionSourceConfig()
// 	source := sources.NewNotionSource(notionClient, config)

// 	err = source.Close()
// 	if err != nil {
// 		t.Errorf("Expected no error closing source, got %v", err)
// 	}
// }

// // Integration tests (require actual Notion API key)

// func TestNotionSource_Read_Integration(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping integration test in short mode")
// 	}

// 	// Only run if we have actual credentials
// 	if test.TestConfig.NotionAPIKey == "" {
// 		t.Skip("Skipping integration test: no Notion API key configured")
// 	}

// 	// Create real client
// 	clientConfig := client.DefaultConfig()
// 	clientConfig.APIKey = test.TestConfig.NotionAPIKey

// 	notionClient, err := client.NewClient(clientConfig)
// 	if err != nil {
// 		t.Fatalf("Failed to create client: %v", err)
// 	}
// 	defer notionClient.Close()

// 	// Create source with limited scope for testing
// 	config := sources.NotionSourceConfig{
// 		IncludeDatabases: false, // Avoid large reads
// 		IncludePages:     true,
// 		IncludeBlocks:    false,
// 		IncludeComments:  false,
// 		IncludeUsers:     false,
// 		// IncludeFiles:      false,
// 		MaxDepth:          1,
// 		RequestsPerSecond: 1.0, // Be conservative
// 		PageSize:          10,  // Small page size for testing
// 	}

// 	source := sources.NewNotionSource(notionClient, config)

// 	ctx := context.Background()
// 	readReq := &engine.ReadRequest{
// 		Types: []engine.ObjectType{engine.ObjectTypePage},
// 		Pagination: &engine.PaginationOptions{
// 			PageSize: 5, // Limit results for testing
// 		},
// 	}

// 	results, err := source.Read(ctx, readReq)
// 	if err != nil {
// 		t.Fatalf("Failed to read from source: %v", err)
// 	}

// 	if results == nil {
// 		t.Fatal("Expected results channel")
// 	}

// 	// Collect a few results
// 	var items []engine.DataItem
// 	count := 0
// 	for item := range results {
// 		items = append(items, item)
// 		count++
// 		if count >= 3 { // Only collect a few items for testing
// 			break
// 		}
// 	}

// 	// We should get at least some results (unless the workspace is empty)
// 	// Note: This test might fail if the workspace has no pages
// 	t.Logf("Retrieved %d items from Notion", len(items))

// 	// Validate the items we got
// 	for i, item := range items {
// 		if item.Type != engine.ObjectTypePage {
// 			t.Errorf("Expected item %d to be page type, got %s", i, item.Type)
// 		}

// 		if item.ID == "" {
// 			t.Errorf("Expected item %d to have non-empty ID", i)
// 		}

// 		if item.Data == nil {
// 			t.Errorf("Expected item %d to have data", i)
// 		}

// 		// Check that source metadata is set
// 		if sourceType, exists := item.GetProperty("source"); !exists || sourceType != "notion" {
// 			t.Errorf("Expected item %d to have source property set to notion", i)
// 		}
// 	}
// }

// func TestNotionSource_ReadSpecificPages_Integration(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping integration test in short mode")
// 	}

// 	if test.TestConfig.NotionAPIKey == "" {
// 		t.Skip("Skipping integration test: no Notion API key configured")
// 	}

// 	// Create real client
// 	clientConfig := client.DefaultConfig()
// 	clientConfig.APIKey = test.TestConfig.NotionAPIKey

// 	notionClient, err := client.NewClient(clientConfig)
// 	if err != nil {
// 		t.Fatalf("Failed to create client: %v", err)
// 	}
// 	defer notionClient.Close()

// 	// Test with specific page IDs (you would need actual page IDs from your workspace)
// 	// This test will be skipped unless you have specific test page IDs
// 	testPageIDs := []string{
// 		// Add actual page IDs from your test workspace here
// 		// "page-id-1",
// 		// "page-id-2",
// 	}

// 	if len(testPageIDs) == 0 {
// 		t.Skip("Skipping specific pages test: no test page IDs configured")
// 	}

// 	config := sources.NotionSourceConfig{
// 		IncludeDatabases:  false,
// 		IncludePages:      true,
// 		IncludeBlocks:     false,
// 		IncludeComments:   false,
// 		PageIDs:           testPageIDs,
// 		RequestsPerSecond: 1.0,
// 	}

// 	source := sources.NewNotionSource(notionClient, config)

// 	results, err := source.Read(context.Background(), &engine.ReadRequest{})
// 	if err != nil {
// 		t.Fatalf("Failed to read specific pages: %v", err)
// 	}

// 	var items []engine.DataItem
// 	for item := range results {
// 		items = append(items, item)
// 	}

// 	if len(items) != len(testPageIDs) {
// 		t.Errorf("Expected %d items, got %d", len(testPageIDs), len(items))
// 	}

// 	for i, item := range items {
// 		if item.Type != engine.ObjectTypePage {
// 			t.Errorf("Expected item %d to be page type, got %s", i, item.Type)
// 		}
// 	}
// }
