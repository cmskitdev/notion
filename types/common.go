// package types provides comprehensive type definitions for the Notion API.
// This package implements a type-safe, well-documented interface for interacting
// with all Notion API entities including pages, databases, blocks, users, and properties.
//
// The package follows Go best practices with extensive documentation, proper error
// handling, and comprehensive validation. All types are designed to be easily
// serializable to/from JSON for API communication.
//
// Example:
//
//	pageID := PageID("550e8400-e29b-41d4-a716-446655440000")
//	if err := pageID.Validate(); err != nil {
//	    log.Fatal(err)
//	}
//
//	page := &Page{
//	    ID:     pageID,
//	    Object: ObjectTypePage,
//	    Title:  []RichText{{Type: RichTextTypeText, Text: &TextContent{Content: "My Page"}}},
//	}
package types

import (
	"encoding/json"
	"fmt"
	"time"
)

// ObjectType represents the fundamental object types in the Notion API.
// All Notion entities implement this basic classification system.
type ObjectType string

const (
	ObjectTypePage        ObjectType = "page"
	ObjectTypeDatabase    ObjectType = "database"
	ObjectTypeBlock       ObjectType = "block"
	ObjectTypeUser        ObjectType = "user"
	ObjectTypeComment     ObjectType = "comment"
	ObjectTypePropertyRow ObjectType = "property_item"
	ObjectTypeList        ObjectType = "list"
)

// Object represents the base interface that all Notion API entities must implement.
// This provides a common structure for identification and type safety across
// all API operations.
type Object interface {
	GetID() string
	GetObject() ObjectType
	GetCreatedTime() time.Time
	GetLastEditedTime() time.Time
}

// Color represents the available colors in Notion for various elements.
// These colors are used for text formatting, select options, and other UI elements.
type Color string

const (
	ColorDefault Color = "default"
	ColorGray    Color = "gray"
	ColorBrown   Color = "brown"
	ColorOrange  Color = "orange"
	ColorYellow  Color = "yellow"
	ColorGreen   Color = "green"
	ColorBlue    Color = "blue"
	ColorPurple  Color = "purple"
	ColorPink    Color = "pink"
	ColorRed     Color = "red"
)

// BackgroundColor represents background color variants available in Notion.
type BackgroundColor string

const (
	BackgroundColorDefault BackgroundColor = "default"
	BackgroundColorGray    BackgroundColor = "gray_background"
	BackgroundColorBrown   BackgroundColor = "brown_background"
	BackgroundColorOrange  BackgroundColor = "orange_background"
	BackgroundColorYellow  BackgroundColor = "yellow_background"
	BackgroundColorGreen   BackgroundColor = "green_background"
	BackgroundColorBlue    BackgroundColor = "blue_background"
	BackgroundColorPurple  BackgroundColor = "purple_background"
	BackgroundColorPink    BackgroundColor = "pink_background"
	BackgroundColorRed     BackgroundColor = "red_background"
)

// Timestamp represents a Notion API timestamp with proper JSON marshaling.
// This type handles the ISO 8601 format used by the Notion API.
type Timestamp struct {
	time.Time
}

// MarshalJSON converts the Timestamp to JSON in ISO 8601 format.
//
// Returns:
// - []byte: JSON-encoded timestamp in ISO 8601 format.
// - error: Marshaling error, if any.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format(time.RFC3339))
}

// UnmarshalJSON parses a JSON timestamp in ISO 8601 format.
//
// Arguments:
// - data: Raw JSON bytes containing the timestamp string.
//
// Returns:
// - error: Parsing error if the timestamp format is invalid, nil if successful.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return fmt.Errorf("invalid timestamp format: %w", err)
	}

	t.Time = parsed
	return nil
}

// BaseObject provides common fields that all Notion objects share.
// This struct can be embedded in specific object types to provide
// standard identification and metadata fields.
type BaseObject struct {
	Object         ObjectType `json:"object"`
	CreatedTime    time.Time  `json:"created_time"`
	LastEditedTime time.Time  `json:"last_edited_time"`
	CreatedBy      *User      `json:"created_by,omitempty"`
	LastEditedBy   *User      `json:"last_edited_by,omitempty"`
}

type Exportable interface {
	GetID() string
	ObjectType() ObjectType
}

// Icon represents an icon that can be applied to pages and databases.
// Icons can be either emoji characters or external file URLs.
type Icon struct {
	Type  IconType `json:"type"`
	Emoji *string  `json:"emoji,omitempty"`
	File  *File    `json:"file,omitempty"`
}

// IconType represents the type of icon being used.
type IconType string

const (
	IconTypeEmoji    IconType = "emoji"
	IconTypeFile     IconType = "file"
	IconTypeExternal IconType = "external"
)

// Cover represents a cover image that can be applied to pages and databases.
// Covers are always external file URLs in the current API.
type Cover struct {
	Type     CoverType `json:"type"`
	File     *File     `json:"file,omitempty"`
	External *File     `json:"external,omitempty"`
}

// CoverType represents the type of cover being used.
type CoverType string

const (
	CoverTypeFile     CoverType = "file"
	CoverTypeExternal CoverType = "external"
)

// Parent represents the parent relationship for pages and blocks.
// This determines where an object is located within the Notion hierarchy.
type Parent struct {
	Type       ParentType  `json:"type"`
	PageID     *PageID     `json:"page_id,omitempty"`
	DatabaseID *DatabaseID `json:"database_id,omitempty"`
	BlockID    *BlockID    `json:"block_id,omitempty"`
	Workspace  *bool       `json:"workspace,omitempty"`
}

// ParentType represents the different types of parent relationships.
type ParentType string

const (
	ParentTypePage      ParentType = "page_id"
	ParentTypeDatabase  ParentType = "database_id"
	ParentTypeBlock     ParentType = "block_id"
	ParentTypeWorkspace ParentType = "workspace"
)

// GetParentID returns the appropriate parent ID based on the parent type.
//
// Returns:
// - string: The parent ID as a string, or empty string if not applicable.
func (p *Parent) GetParentID() string {
	switch p.Type {
	case ParentTypePage:
		if p.PageID != nil {
			return p.PageID.String()
		}
	case ParentTypeDatabase:
		if p.DatabaseID != nil {
			return p.DatabaseID.String()
		}
	case ParentTypeBlock:
		if p.BlockID != nil {
			return p.BlockID.String()
		}
	}
	return ""
}

// Validate ensures the Parent has the correct fields set for its type.
//
// Returns:
// - error: Validation error if required fields are missing, nil if valid.
func (p *Parent) Validate() error {
	switch p.Type {
	case ParentTypePage:
		if p.PageID == nil {
			return fmt.Errorf("page_id is required for parent type 'page_id'")
		}
	case ParentTypeDatabase:
		if p.DatabaseID == nil {
			return fmt.Errorf("database_id is required for parent type 'database_id'")
		}
	case ParentTypeBlock:
		if p.BlockID == nil {
			return fmt.Errorf("block_id is required for parent type 'block_id'")
		}
	case ParentTypeWorkspace:
		if p.Workspace == nil || !*p.Workspace {
			return fmt.Errorf("workspace must be true for parent type 'workspace'")
		}
		return nil
	default:
		return fmt.Errorf("unknown parent type: %s", p.Type)
	}
	return nil
}
