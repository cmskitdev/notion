// package types provides types for Notion pages.
package types

import (
	"encoding/json"
	"fmt"

	"github.com/cmskitdev/notion/id"
)

// PageID represents a unique identifier for a Notion page.
type PageID id.ID

// String returns the string representation of the PageID.
//
// Returns:
// - string: The string representation of the PageID.
func (p *PageID) String() string {
	return string(*p)
}

// ParsePageID parses a string as a Notion PageID, accepting both dashed and undashed formats.
// It returns the PageID in canonical dashed UUID format (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).
//
// Arguments:
// - id: The string to parse as a PageID.
//
// Returns:
// - PageID: The canonical, dashed PageID.
// - error: If the input is not a valid 32-character hexadecimal string.
func ParsePageID(id string) (PageID, error) {
	pageID, err := IDParser.Parse(id)
	if err != nil {
		return "", err
	}
	return PageID(pageID), nil
}

// UnmarshalJSON implements custom JSON unmarshaling for PageID.
// It accepts both dashed and undashed UUID formats and converts them to canonical format.
//
// Arguments:
// - data: Raw JSON bytes containing the ID string.
//
// Returns:
// - error: Parsing error if the ID format is invalid, nil if successful.
func (p *PageID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := ParsePageID(s)
	if err != nil {
		return fmt.Errorf("invalid PageID in JSON: %w", err)
	}

	*p = parsed
	return nil
}

// MarshalJSON implements custom JSON marshaling for PageID.
//
// Returns:
// - []byte: JSON-encoded PageID as a string.
// - error: Marshaling error, if any.
func (p PageID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

// Validate ensures the PageID is in the correct UUID format.
//
// Returns:
// - error: Validation error if the ID format is invalid, nil if valid.
//
// Example:
//
//	if err := pageID.Validate(); err != nil {
//	    log.Printf("Invalid PageID: %v", err)
//	}
func (p PageID) Validate() error {
	_, err := ParsePageID(string(p))
	return err
}

// NewPage creates a new page with the specified parent and properties.
//
// Arguments:
// - parent: The parent container for the page (database, page, or workspace).
// - properties: A map of property names to property values.
//
// Returns:
// - *Page: A new page with the specified parent and properties.
//
// Example:
//
//	parent := &Parent{Type: ParentTypeDatabase, DatabaseID: &dbID}
//	properties := map[string]Property{
//	    "Name": *NewTitleProperty("My New Page"),
//	    "Status": *NewSelectProperty(&"In Progress", &ColorYellow),
//	}
//	page := NewPage(parent, properties)
func NewPage(parent *Parent, properties map[string]Property) *Page {
	return &Page{
		ID: "todo",
		BaseObject: BaseObject{
			Object: ObjectTypePage,
		},
		PropertyAccessor: PropertyAccessor[Property]{
			PropertyContainer: &PropertyContainer[Property]{
				Properties: properties,
			},
		},
		Parent: parent,
	}
}

// Page represents a Notion page, which is a structured document containing blocks and properties.
// Pages can exist independently or within databases as records.
type Page struct {
	BaseObject
	PropertyAccessor[Property]
	ID        PageID  `json:"id"`
	Parent    *Parent `json:"parent,omitempty"`
	Icon      *Icon   `json:"icon,omitempty"`
	Cover     *Cover  `json:"cover,omitempty"`
	URL       string  `json:"url,omitempty"`
	PublicURL *string `json:"public_url,omitempty"`
	Archived  bool    `json:"archived,omitempty"`
	InTrash   bool    `json:"in_trash,omitempty"`
}

// GetID returns the ID of the page.
//
// Returns:
// - string: The ID of the page.
func (p Page) GetID() string {
	return p.ID.String()
}

// ObjectType returns the type of the page.
//
// Returns:
// - ObjectType: The type of the page.
func (p Page) ObjectType() ObjectType {
	return ObjectTypePage
}

// Note: GetProperty, SetProperty, RemoveProperty, GetPropertyValue, and typed property accessors
// are now inherited from PropertyAccessor

// IsInDatabase returns true if the page is a child of a database.
//
// Returns:
// - bool: True if the page is in a database, false otherwise.
//
// Example:
//
//	if page.IsInDatabase() {
//	    fmt.Println("This page is a database record")
//	}
func (p *Page) IsInDatabase() bool {
	return p.Parent != nil && p.Parent.Type == ParentTypeDatabase
}

// IsTopLevel returns true if the page is at the workspace level.
//
// Returns:
// - bool: True if the page is at the workspace level, false otherwise.
//
// Example:
//
//	if page.IsTopLevel() {
//	    fmt.Println("This is a top-level page")
//	}
func (p *Page) IsTopLevel() bool {
	return p.Parent != nil && p.Parent.Type == ParentTypeWorkspace
}

// GetParentID returns the ID of the page's parent.
//
// Returns:
// - string: The parent ID, or empty string if no parent or workspace parent.
//
// Example:
//
//	parentID := page.GetParentID()
//	if parentID != "" {
//	    fmt.Printf("Parent ID: %s\n", parentID)
//	}
func (p *Page) GetParentID() string {
	if p.Parent == nil {
		return ""
	}
	return p.Parent.GetParentID()
}

// SetIcon sets the icon for the page.
//
// Arguments:
// - icon: The icon to set (can be emoji or file).
//
// Example:
//
//	page.SetIcon(&Icon{Type: IconTypeEmoji, Emoji: &"📄"})
func (p *Page) SetIcon(icon *Icon) {
	p.Icon = icon
}

// SetCover sets the cover image for the page.
//
// Arguments:
// - cover: The cover to set.
//
// Example:
//
//	cover := &Cover{
//	    Type: CoverTypeExternal,
//	    External: &File{URL: "https://example.com/image.jpg"},
//	}
//	page.SetCover(cover)
func (p *Page) SetCover(cover *Cover) {
	p.Cover = cover
}

// Archive marks the page as archived.
//
// Example:
//
//	page.Archive()
//	fmt.Println("Page archived:", page.Archived)
func (p *Page) Archive() {
	p.Archived = true
}

// Unarchive marks the page as not archived.
//
// Example:
//
//	page.Unarchive()
//	fmt.Println("Page unarchived:", page.Archived == false)
func (p *Page) Unarchive() {
	p.Archived = false
}

// Validate ensures the Page has valid required fields and structure.
//
// Returns:
// - error: Validation error if the page structure is invalid, nil if valid.
//
// Example:
//
//	if err := page.Validate(); err != nil {
//	    log.Printf("Invalid page: %v", err)
//	}
func (p *Page) Validate() error {
	if p.Object != ObjectTypePage {
		return fmt.Errorf("object type must be 'page', got: %s", p.Object)
	}

	if p.Parent != nil {
		if err := p.Parent.Validate(); err != nil {
			return fmt.Errorf("invalid parent: %w", err)
		}
	}

	// Validate properties using inherited method
	if err := p.ValidateProperties(); err != nil {
		return err
	}

	// Validate that database pages have required title property
	if p.IsInDatabase() {
		hasTitle := false
		for _, prop := range p.Properties {
			if prop.Type == PropertyTypeTitle {
				hasTitle = true
				break
			}
		}
		if !hasTitle {
			return fmt.Errorf("database pages must have a title property")
		}
	}

	// Validate icon if present
	if p.Icon != nil {
		if err := p.Icon.Validate(); err != nil {
			return fmt.Errorf("invalid icon: %w", err)
		}
	}

	// Validate cover if present
	if p.Cover != nil {
		if err := p.Cover.Validate(); err != nil {
			return fmt.Errorf("invalid cover: %w", err)
		}
	}

	return nil
}

// Validate ensures the Icon has valid configuration.
//
// Returns:
// - error: Validation error if the icon configuration is invalid, nil if valid.
func (i *Icon) Validate() error {
	switch i.Type {
	case IconTypeEmoji:
		if i.Emoji == nil || *i.Emoji == "" {
			return fmt.Errorf("emoji field is required for icon type 'emoji'")
		}
		if i.File != nil {
			return fmt.Errorf("file field should be nil for icon type 'emoji'")
		}
	case IconTypeFile, IconTypeExternal:
		if i.File == nil {
			return fmt.Errorf("file field is required for icon type '%s'", i.Type)
		}
		if i.File.File.URL == "" {
			return fmt.Errorf("file URL cannot be empty")
		}
		if i.Emoji != nil {
			return fmt.Errorf("emoji field should be nil for icon type '%s'", i.Type)
		}
	default:
		return fmt.Errorf("unknown icon type: %s", i.Type)
	}
	return nil
}

// Validate ensures the Cover has valid configuration.
//
// Returns:
// - error: Validation error if the cover configuration is invalid, nil if valid.
func (c *Cover) Validate() error {
	switch c.Type {
	case CoverTypeFile:
		if c.File == nil {
			return fmt.Errorf("file field is required for cover type 'file'")
		}
		if c.File.File.URL == "" {
			return fmt.Errorf("file URL cannot be empty")
		}
		if c.External != nil {
			return fmt.Errorf("external field should be nil for cover type 'file'")
		}
	case CoverTypeExternal:
		if c.External == nil {
			return fmt.Errorf("external field is required for cover type 'external'")
		}
		if c.External.External.URL == "" {
			return fmt.Errorf("external URL cannot be empty")
		}
		if c.File != nil {
			return fmt.Errorf("file field should be nil for cover type 'external'")
		}
	default:
		return fmt.Errorf("unknown cover type: %s", c.Type)
	}
	return nil
}

// PageCreateRequest represents a request to create a new page.
type PageCreateRequest struct {
	Parent     Parent              `json:"parent"`
	Properties map[string]Property `json:"properties"`
	Children   []Block             `json:"children,omitempty"`
	Icon       *Icon               `json:"icon,omitempty"`
	Cover      *Cover              `json:"cover,omitempty"`
}

// PageUpdateRequest represents a request to update an existing page.
type PageUpdateRequest struct {
	Properties map[string]Property `json:"properties,omitempty"`
	Archived   *bool               `json:"archived,omitempty"`
	Icon       *Icon               `json:"icon,omitempty"`
	Cover      *Cover              `json:"cover,omitempty"`
	InTrash    *bool               `json:"in_trash,omitempty"`
}

// NewPageCreateRequest creates a new page creation request.
//
// Arguments:
// - parent: The parent container for the new page.
// - properties: Initial properties for the page.
//
// Returns:
// - *PageCreateRequest: A new page creation request.
//
// Example:
//
//	parent := Parent{Type: ParentTypeDatabase, DatabaseID: &dbID}
//	properties := map[string]Property{
//	    "Name": *NewTitleProperty("New Task"),
//	}
//	request := NewPageCreateRequest(parent, properties)
func NewPageCreateRequest(parent Parent, properties map[string]Property) *PageCreateRequest {
	return &PageCreateRequest{
		Parent:     parent,
		Properties: properties,
	}
}

// AddChild adds a child block to the page creation request.
//
// Arguments:
// - block: The block to add as a child.
//
// Example:
//
//	request.AddChild(*NewParagraphBlock([]RichText{*NewTextRichText("Hello!", nil)}))
func (pcr *PageCreateRequest) AddChild(block Block) {
	pcr.Children = append(pcr.Children, block)
}

// NewPageUpdateRequest creates a new page update request.
//
// Returns:
// - *PageUpdateRequest: A new empty page update request.
//
// Example:
//
//	request := NewPageUpdateRequest()
//	request.SetProperty("Status", *NewSelectProperty(&"Done", &ColorGreen))
func NewPageUpdateRequest() *PageUpdateRequest {
	return &PageUpdateRequest{
		Properties: make(map[string]Property),
	}
}

// SetProperty sets a property in the update request.
//
// Arguments:
// - name: The name of the property to update.
// - property: The new property value.
//
// Example:
//
//	request.SetProperty("Priority", *NewSelectProperty(&"High", &ColorRed))
func (pur *PageUpdateRequest) SetProperty(name string, property Property) {
	if pur.Properties == nil {
		pur.Properties = make(map[string]Property)
	}
	pur.Properties[name] = property
}
