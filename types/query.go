package types

// QueryFilter represents a filter for database query results.
// See https://developers.notion.com/reference/post-database-query-filter.
type QueryFilter struct {
	Property *string `json:"property,omitempty"`

	// Compound filters
	And []*QueryFilter `json:"and,omitempty"`
	Or  []*QueryFilter `json:"or,omitempty"`

	// Property-specific filters
	Checkbox  *CheckboxFilter  `json:"checkbox,omitempty"`
	Date      *DateFilter      `json:"date,omitempty"`
	Files     *FilesFilter     `json:"files,omitempty"`
	Number    *NumberFilter    `json:"number,omitempty"`
	RichText  *RichTextFilter  `json:"rich_text,omitempty"`
	Select    *SelectFilter    `json:"select,omitempty"`
	Status    *StatusFilter    `json:"status,omitempty"`
	Timestamp *TimestampFilter `json:"timestamp,omitempty"`
}

// QuerySort represents a sort order for database query results.
// Note: Only one of PropertySort or TimestampSort can be used.
// See: https://developers.notion.com/reference/post-database-query-sort.
// TODO: Investigate if this is mutually exclusive or not (notion documentation is unclear).
type QuerySort struct {
	// Sort by property(s).
	PropertySort *PropertySort `json:"property,omitempty"`
	// Sort by timestamp.
	TimestampSort *TimestampSort `json:"timestamp,omitempty"`
}

// Query represents a request to query a Notion database.
// See https://developers.notion.com/reference/post-database-query.
type Query struct {
	Filter      *QueryFilter `json:"filter,omitempty"`
	Sorts       []*QuerySort `json:"sorts,omitempty"`
	StartCursor *string      `json:"start_cursor,omitempty"`
	PageSize    *int         `json:"page_size,omitempty"`
}
