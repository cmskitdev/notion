package types

// ResourceType represents the type of Notion resource.
type ResourceType string

const (
	// ResourceTypePage represents a Notion page.
	ResourceTypePage ResourceType = "page"
	// ResourceTypeDatabase represents a Notion database.
	ResourceTypeDatabase ResourceType = "database"
	// ResourceTypeBlock represents a Notion block.
	ResourceTypeBlock ResourceType = "block"
	// ResourceTypeUser represents a Notion user.
	ResourceTypeUser ResourceType = "user"
	// ResourceTypeProperty represents a Notion property.
	ResourceTypeProperty ResourceType = "property"
	// ResourceTypeComment represents a Notion comment.
	ResourceTypeComment ResourceType = "comment"
)
