package types

import (
	"encoding/json"
	"net/url"
)

// SearchFilter represents a filter for search results.
// See https://developers.notion.com/reference/post-search.
type SearchFilter struct {
	// The property to filter against.
	Property string `json:"property"`
	// The value to filter against.
	Value string `json:"value"`
}

// SearchSort represents a sort order for search results.
// See https://developers.notion.com/reference/post-search.
type SearchSort struct {
	// The name of the timestamp to sort against. Possible values include last_edited_time.
	Timestamp string `json:"property,omitempty"`
	// The direction to sort in. Possible values include ascending and descending.
	Direction string `json:"direction,omitempty"`
}

// SearchRequest represents a request to search the Notion API.
// See https://developers.notion.com/reference/post-search.
type SearchRequest struct {
	// The query string to search for.
	Query string `json:"query,omitempty"`
	// The filter to apply to the search results.
	Filter *SearchFilter `json:"filter,omitempty"`
	// The sort order to apply to the search results.
	Sort *SearchSort `json:"sort,omitempty"`
	// The start cursor to use for pagination.
	StartCursor *string `json:"start_cursor,omitempty"`
	// The page size to use for pagination.
	PageSize *int `json:"page_size,omitempty"`
}

func (sr *SearchRequest) GetPath() string {
	return "/search"
}

func (sr *SearchRequest) GetMethod() string {
	return "POST"
}

func (sr *SearchRequest) GetBody() SearchRequest {
	return *sr
}

func (sr *SearchRequest) GetQuery() url.Values {
	return nil
}

// SearchResult represents a search result that can be either a page or database.
type SearchResult struct {
	BaseObject
	Object   string    `json:"object"`
	Page     *Page     `json:",omitempty"`
	Database *Database `json:",omitempty"`
	Block    *Block    `json:",omitempty"`
	User     *User     `json:",omitempty"`
}

// UnmarshalJSON implements custom JSON unmarshaling for SearchResult.
// This handles the union type where the response can be either a page or database.
func (sr *SearchResult) UnmarshalJSON(data []byte) error {
	// First unmarshal into a map to determine the object type
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Set the object type
	if obj, ok := raw["object"].(string); ok {
		sr.Object = obj
	}

	// Based on the object type, unmarshal into the appropriate field
	switch sr.Object {
	case "page":
		var page Page
		if err := json.Unmarshal(data, &page); err != nil {
			return err
		}
		sr.Page = &page
		sr.BaseObject = page.BaseObject
	case "database":
		var database Database
		if err := json.Unmarshal(data, &database); err != nil {
			return err
		}
		sr.Database = &database
		sr.BaseObject = database.BaseObject
	case "block":
		var block Block
		if err := json.Unmarshal(data, &block); err != nil {
			return err
		}
		sr.Block = &block
		sr.BaseObject = block.BaseObject
	case "user":
		var user User
		if err := json.Unmarshal(data, &user); err != nil {
			return err
		}
		sr.User = &user
		sr.BaseObject = user.BaseObject
	}

	return nil
}
