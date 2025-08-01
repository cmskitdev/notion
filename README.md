# Notion Plugin for cmskit

This package provides comprehensive, type-safe Go definitions for all Notion API entities including pages, databases, blocks, users, and properties. The types are designed for high performance, excellent developer experience, and seamless JSON serialization for API communication.

## Features

- **Complete Type Coverage**: All Notion API types with 20+ property types and 25+ block types
- **Type Safety**: Branded ID types prevent mixing different entity identifiers  
- **Comprehensive Validation**: Built-in validation ensures data integrity
- **Rich Documentation**: Extensive documentation with examples following Go best practices
- **JSON Serialization**: Optimized for Notion API communication
- **Goroutine Ready**: Designed for concurrent operations and channel-based architectures

## Quick Start

```go
import "github.com/cmskitdev/notion/types"

// Create a page with properties
parent := &types.Parent{
    Type: types.ParentTypeDatabase, 
    DatabaseID: &types.DatabaseID("550e8400-e29b-41d4-a716-446655440000"),
}

properties := map[string]types.Property{
    "Name": *types.NewTitleProperty("My New Page"),
    "Status": *types.NewSelectProperty(&"In Progress", &types.ColorYellow),
    "Completed": *types.NewCheckboxProperty(false),
}

page := types.NewPage(parent, properties)
```

## Core Types

### Identity Types

- `PageID`, `DatabaseID`, `BlockID`, `UserID`, `PropertyID` - Type-safe identifiers with validation
- `ObjectType` - Fundamental object classification system
- `Timestamp` - RFC3339 timestamp handling with JSON marshaling

### Content Types  

- `RichText` - Rich text with formatting, mentions, equations
- `Property` - Database properties with 20+ supported types
- `Block` - Page content blocks with 25+ supported types

### Entity Types

- `Page` - Notion pages with properties and metadata
- `Database` - Database schemas with property configurations  
- `User` - Person and bot users with full type discrimination
- `Parent` - Hierarchical relationships between entities

## Advanced Usage

### Creating Rich Content

```go
// Rich text with formatting
richText := types.NewTextRichText("Important text", &types.Annotations{
    Bold: true,
    Color: types.ColorRed,
})

// User mentions
mention := types.NewUserMention("550e8400-e29b-41d4-a716-446655440000")

// Links with custom text
link := types.NewLinkRichText("Visit our site", "https://example.com", nil)
```

### Building Complex Blocks

```go
// Paragraph with rich content
paragraph := types.NewParagraphBlock([]types.RichText{
    *types.NewTextRichText("This is ", nil),
    *types.NewTextRichText("bold text", &types.Annotations{Bold: true}),
    *types.NewTextRichText(" with a ", nil),
    *types.NewLinkRichText("link", "https://example.com", nil),
})

// To-do items
todo := types.NewToDoBlock([]types.RichText{
    *types.NewTextRichText("Complete the project", nil),
}, false)

// Code blocks with syntax highlighting
code := types.NewCodeBlock([]types.RichText{
    *types.NewTextRichText(`func main() {\n\tfmt.Println("Hello, World!")\n}`, nil),
}, "go")
```

### Database Schema Design

```go
properties := map[string]types.DatabaseProperty{
    "Name": {
        Type: types.PropertyTypeTitle,
        Title: &types.TitleConfig{},
    },
    "Status": {
        Type: types.PropertyTypeSelect,
        Select: &types.SelectConfig{
            Options: []types.SelectOption{
                {Name: "Not Started", Color: types.ColorRed},
                {Name: "In Progress", Color: types.ColorYellow}, 
                {Name: "Complete", Color: types.ColorGreen},
            },
        },
    },
    "Priority": {
        Type: types.PropertyTypeNumber,
        Number: &types.NumberConfig{Format: types.NumberFormatNumber},
    },
}

database := types.NewDatabase(
    []types.RichText{*types.NewTextRichText("Project Tasks", nil)},
    properties,
)
```

## Validation and Error Handling

All types include comprehensive validation:

```go
page := types.NewPage(parent, properties)
if err := page.Validate(); err != nil {
    log.Printf("Invalid page: %v", err)
    return
}

// ID validation
pageID := types.PageID("invalid-id")
if err := pageID.Validate(); err != nil {
    log.Printf("Invalid page ID: %v", err)
    return
}
```

## JSON Serialization

Types are optimized for JSON marshaling/unmarshaling:

```go
// Marshal to JSON for API requests
data, err := json.Marshal(page)
if err != nil {
    log.Printf("Failed to marshal: %v", err)
    return
}

// Unmarshal API responses
var response types.Page
if err := json.Unmarshal(data, &response); err != nil {
    log.Printf("Failed to unmarshal: %v", err)
    return  
}
```

## Type Safety Examples

The package prevents common mistakes through type safety:

```go
// This won't compile - prevents mixing ID types
var pageID types.PageID = "550e8400-e29b-41d4-a716-446655440000"  
var dbID types.DatabaseID = pageID // ✗ Compile error

// This works - explicit conversion when needed
var dbID types.DatabaseID = types.DatabaseID(pageID) // ✓ Explicit conversion
```

## Property Type Coverage

**Basic Properties:**

- Title, Rich Text, Number, Select, Multi-select
- Date, People, Files, Checkbox, URL, Email, Phone

**Advanced Properties:**  

- Formula, Relation, Rollup, Status, Unique ID, Verification
- Created Time/By, Last Edited Time/By

**Property Configurations:**

- Select options with colors
- Formula expressions
- Relation and rollup configurations
- Number formatting options (currency, percentage, etc.)

## Block Type Coverage

**Text Blocks:**

- Paragraph, Headings (H1-H3), Lists (bulleted/numbered)
- To-do, Toggle, Quote, Callout, Code

**Media Blocks:**

- Image, Video, Audio, File, PDF, Bookmark, Embed

**Structure Blocks:**

- Child Page/Database, Divider, Table of Contents
- Column List, Synced Block, Template, Link to Page, Table

## Testing

Run the comprehensive test suite:

```bash
go test -v ./...
```

Tests cover:

- ID validation for all types
- JSON marshaling/unmarshaling round trips  
- Property creation and manipulation
- Block construction and validation
- Page and database operations
- User type discrimination
- Error conditions and edge cases

## Performance Considerations

- **Memory Efficient**: Minimal allocations through careful struct design
- **JSON Optimized**: Fast marshaling/unmarshaling with proper field tags
- **Validation Caching**: Expensive validations can be cached at application level
- **Goroutine Safe**: All types are safe for concurrent read access
- **Channel Ready**: Designed for streaming and concurrent processing

## Architecture Integration

This package is designed to integrate with:

- **HTTP Clients**: Direct JSON serialization to/from Notion API
- **Goroutine Pools**: Safe concurrent processing of entities
- **Channel Pipelines**: Streaming processing of pages/blocks/properties  
- **Validation Middleware**: Request/response validation layers
- **Caching Systems**: Efficient serialization for Redis/other caches

## Migration from TypeScript

Key differences from the TypeScript SDK:

- **Branded Types**: Go uses distinct types instead of branded strings
- **Explicit Validation**: Manual validation calls instead of runtime checking
- **Pointer Semantics**: Optional fields use pointers for nil representation
- **Interface Design**: Go interfaces for polymorphic behavior instead of union types

## Contributing

When adding new types:

1. Follow the existing documentation patterns
2. Add comprehensive validation
3. Include constructor functions for common use cases
4. Add test coverage for all validation paths
5. Update this README with usage examples

## License

This package is designed to be used with the Notion API in accordance with Notion's Terms of Service.
