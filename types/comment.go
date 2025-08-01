package types

import (
	"encoding/json"
	"fmt"

	"github.com/cmskitdev/notion/id"
)

// CommentID represents a unique identifier for a Notion comment.
type CommentID id.ID

// String returns the string representation of the CommentID.
//
// Returns:
// - string: The string representation of the CommentID.
func (c *CommentID) String() string {
	return string(*c)
}

// ParseCommentID parses a string as a Notion CommentID, accepting both dashed and undashed formats.
// It returns the CommentID in canonical dashed UUID format (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).
//
// Arguments:
// - id: The string to parse as a CommentID.
//
// Returns:
// - CommentID: The canonical, dashed CommentID.
// - error: If the input is not a valid 32-character hexadecimal string.
func ParseCommentID(id string) (CommentID, error) {
	commentID, err := IDParser.Parse(id)
	if err != nil {
		return "", err
	}
	return CommentID(commentID), nil
}

// UnmarshalJSON implements custom JSON unmarshaling for CommentID.
// It accepts both dashed and undashed UUID formats and converts them to canonical format.
//
// Arguments:
// - data: Raw JSON bytes containing the ID string.
//
// Returns:
// - error: Parsing error if the ID format is invalid, nil if successful.
func (c *CommentID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := ParseCommentID(s)
	if err != nil {
		return fmt.Errorf("invalid CommentID in JSON: %w", err)
	}

	*c = parsed
	return nil
}

// MarshalJSON implements custom JSON marshaling for CommentID.
//
// Returns:
// - []byte: JSON-encoded CommentID as a string.
// - error: Marshaling error, if any.
func (c CommentID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

// Validate ensures the CommentID is in the correct UUID format.
//
// Returns:
// - error: Validation error if the ID format is invalid, nil if valid.
//
// Example:
//
//	if err := commentID.Validate(); err != nil {
//	    log.Printf("Invalid CommentID: %v", err)
//	}
func (c CommentID) Validate() error {
	_, err := ParseCommentID(string(c))
	return err
}

// DiscussionID represents a unique identifier for a Notion comment discussion thread.
type DiscussionID id.ID

// String returns the string representation of the DiscussionID.
//
// Returns:
// - string: The string representation of the DiscussionID.
func (d *DiscussionID) String() string {
	return string(*d)
}

// ParseDiscussionID parses a string as a Notion DiscussionID, accepting both dashed and undashed formats.
// It returns the DiscussionID in canonical dashed UUID format (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).
//
// Arguments:
// - id: The string to parse as a DiscussionID.
//
// Returns:
// - DiscussionID: The canonical, dashed DiscussionID.
// - error: If the input is not a valid 32-character hexadecimal string.
func ParseDiscussionID(id string) (DiscussionID, error) {
	discussionID, err := IDParser.Parse(id)
	if err != nil {
		return "", err
	}
	return DiscussionID(discussionID), nil
}

// UnmarshalJSON implements custom JSON unmarshaling for DiscussionID.
// It accepts both dashed and undashed UUID formats and converts them to canonical format.
//
// Arguments:
// - data: Raw JSON bytes containing the ID string.
//
// Returns:
// - error: Parsing error if the ID format is invalid, nil if successful.
func (d *DiscussionID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := ParseDiscussionID(s)
	if err != nil {
		return fmt.Errorf("invalid DiscussionID in JSON: %w", err)
	}

	*d = parsed
	return nil
}

// MarshalJSON implements custom JSON marshaling for DiscussionID.
//
// Returns:
// - []byte: JSON-encoded DiscussionID as a string.
// - error: Marshaling error, if any.
func (d DiscussionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(d))
}

// Validate ensures the DiscussionID is in the correct UUID format.
//
// Returns:
// - error: Validation error if the ID format is invalid, nil if valid.
//
// Example:
//
//	if err := discussionID.Validate(); err != nil {
//	    log.Printf("Invalid DiscussionID: %v", err)
//	}
func (d DiscussionID) Validate() error {
	_, err := ParseDiscussionID(string(d))
	return err
}

// Comment represents a comment in Notion, which can be attached to pages or blocks.
// Comments are part of discussion threads and contain rich text content.
type Comment struct {
	BaseObject
	ID           CommentID      `json:"id"`
	Parent       *CommentParent `json:"parent,omitempty"`
	DiscussionID DiscussionID   `json:"discussion_id" validate:"uuid"`
	RichText     []RichText     `json:"rich_text" validate:"required"`
}

// CommentParent represents the parent object of a comment.
// Comments can be attached to either pages or blocks.
type CommentParent struct {
	Type    CommentParentType `json:"type"`
	PageID  *PageID           `json:"page_id,omitempty" validate:"uuid"`
	BlockID *BlockID          `json:"block_id,omitempty" validate:"uuid"`
}

// CommentParentType represents the type of parent object for a comment.
type CommentParentType string

const (
	// CommentParentTypePage indicates the comment is attached to a page.
	CommentParentTypePage CommentParentType = "page_id"
	// CommentParentTypeBlock indicates the comment is attached to a block.
	CommentParentTypeBlock CommentParentType = "block_id"
)

// CommentAttachment represents an attachment in a comment.
// See: https://developers.notion.com/reference/comment-attachment
type CommentAttachment struct {
	// For input: UUID of uploaded file
	FileUploadID *string `json:"file_upload_id,omitempty" validate:"uuid"`
	// Type of attachment (currently only "file_upload")
	Type *string `json:"type,omitempty" validate:"required,equal=file_upload"`

	// For response: file category and details
	Category *string            `json:"category,omitempty"`
	File     *CommentFileObject `json:"file,omitempty"`
}

// CommentFileObject represents file details in a comment attachment.
type CommentFileObject struct {
	URL        string `json:"url"`
	ExpiryTime string `json:"expiry_time"`
}

// CommentDisplayName represents the display name configuration for a comment.
// See: https://developers.notion.com/reference/comment-attachment-copy
type CommentDisplayName struct {
	Type         string                    `json:"type" validate:"required,oneOf=integration,user,custom"`
	Custom       *CommentDisplayNameCustom `json:"custom,omitempty" validate:"object"`
	ResolvedName *string                   `json:"resolved_name,omitempty"`
}

// CommentDisplayNameCustom represents custom display name settings.
type CommentDisplayNameCustom struct {
	Name string `json:"name"`
}

// GetID returns the ID of the comment.
//
// Returns:
// - string: The ID of the comment.
func (c Comment) GetID() string {
	return c.ID.String()
}

// ObjectType returns the type of the comment.
//
// Returns:
// - ObjectType: The type of the comment.
func (c Comment) ObjectType() ObjectType {
	return ObjectTypeComment
}

// GetParentID returns the ID of the comment's parent.
//
// Returns:
// - string: The parent ID, or empty string if no parent.
//
// Example:
//
//	parentID := comment.GetParentID()
//	if parentID != "" {
//	    fmt.Printf("Parent ID: %s\n", parentID)
//	}
func (c *Comment) GetParentID() string {
	if c.Parent == nil {
		return ""
	}

	switch c.Parent.Type {
	case CommentParentTypePage:
		if c.Parent.PageID != nil {
			return c.Parent.PageID.String()
		}
	case CommentParentTypeBlock:
		if c.Parent.BlockID != nil {
			return c.Parent.BlockID.String()
		}
	}

	return ""
}

// IsPageComment returns true if the comment is attached to a page.
//
// Returns:
// - bool: True if the comment is attached to a page, false otherwise.
//
// Example:
//
//	if comment.IsPageComment() {
//	    fmt.Println("This comment is on a page")
//	}
func (c *Comment) IsPageComment() bool {
	return c.Parent != nil && c.Parent.Type == CommentParentTypePage
}

// IsBlockComment returns true if the comment is attached to a block.
//
// Returns:
// - bool: True if the comment is attached to a block, false otherwise.
//
// Example:
//
//	if comment.IsBlockComment() {
//	    fmt.Println("This comment is on a block")
//	}
func (c *Comment) IsBlockComment() bool {
	return c.Parent != nil && c.Parent.Type == CommentParentTypeBlock
}

// GetPlainText returns the comment content as plain text.
//
// Returns:
// - string: The comment content as plain text.
//
// Example:
//
//	text := comment.GetPlainText()
//	fmt.Printf("Comment: %s\n", text)
func (c *Comment) GetPlainText() string {
	return ToPlainText(c.RichText)
}

// Validate ensures the Comment has valid required fields and structure.
//
// Returns:
// - error: Validation error if the comment structure is invalid, nil if valid.
//
// Example:
//
//	if err := comment.Validate(); err != nil {
//	    log.Printf("Invalid comment: %v", err)
//	}
func (c *Comment) Validate() error {
	if c.Object != ObjectTypeComment {
		return fmt.Errorf("object type must be 'comment', got: %s", c.Object)
	}

	// Validate parent if present
	if c.Parent != nil {
		if err := c.Parent.Validate(); err != nil {
			return fmt.Errorf("invalid parent: %w", err)
		}
	}

	// Validate rich text content
	for i, rt := range c.RichText {
		if err := rt.Validate(); err != nil {
			return fmt.Errorf("invalid rich text at index %d: %w", i, err)
		}
	}

	return nil
}

// Validate ensures the CommentParent has valid configuration based on its type.
//
// Returns:
// - error: Validation error if the parent configuration is invalid, nil if valid.
func (cp *CommentParent) Validate() error {
	switch cp.Type {
	case CommentParentTypePage:
		if cp.PageID == nil {
			return fmt.Errorf("page_id is required for parent type 'page_id'")
		}
		if cp.BlockID != nil {
			return fmt.Errorf("block_id should be nil for parent type 'page_id'")
		}
		return cp.PageID.Validate()
	case CommentParentTypeBlock:
		if cp.BlockID == nil {
			return fmt.Errorf("block_id is required for parent type 'block_id'")
		}
		if cp.PageID != nil {
			return fmt.Errorf("page_id should be nil for parent type 'block_id'")
		}
		return cp.BlockID.Validate()
	default:
		return fmt.Errorf("unknown parent type: %s", cp.Type)
	}
}

// NewComment creates a new comment with the specified parent and rich text content.
//
// Arguments:
// - parent: The parent object (page or block) for the comment.
// - richText: The rich text content of the comment.
//
// Returns:
// - *Comment: A new comment with the specified parent and content.
//
// Example:
//
//	parent := &CommentParent{
//	    Type: CommentParentTypePage,
//	    PageID: &pageID,
//	}
//	richText := []RichText{*NewTextRichText("Great work!", nil)}
//	comment := NewComment(parent, richText)
func NewComment(parent *CommentParent, richText []RichText) *Comment {
	return &Comment{
		BaseObject: BaseObject{
			Object: ObjectTypeComment,
		},
		Parent:   parent,
		RichText: richText,
	}
}

// CommentCreateRequest represents a request to create a new comment.
type CommentCreateRequest struct {
	Parent   CommentParent `json:"parent"`
	RichText []RichText    `json:"rich_text"`
}

// NewCommentCreateRequest creates a new comment creation request.
//
// Arguments:
// - parent: The parent object for the new comment.
// - richText: The rich text content for the comment.
//
// Returns:
// - *CommentCreateRequest: A new comment creation request.
//
// Example:
//
//	parent := CommentParent{Type: CommentParentTypePage, PageID: &pageID}
//	richText := []RichText{*NewTextRichText("Nice job!", nil)}
//	request := NewCommentCreateRequest(parent, richText)
func NewCommentCreateRequest(parent CommentParent, richText []RichText) *CommentCreateRequest {
	return &CommentCreateRequest{
		Parent:   parent,
		RichText: richText,
	}
}

// CommentListResponse represents a response containing a list of comments.
type CommentListResponse struct {
	Object     ObjectType `json:"object"`
	Results    []Comment  `json:"results"`
	NextCursor *string    `json:"next_cursor"`
	HasMore    bool       `json:"has_more"`
	Type       string     `json:"type"`
	Comment    struct{}   `json:"comment"`
}

// Validate ensures the CommentListResponse has valid structure.
//
// Returns:
// - error: Validation error if the response structure is invalid, nil if valid.
func (clr *CommentListResponse) Validate() error {
	if clr.Object != ObjectTypeList {
		return fmt.Errorf("object type must be 'list', got: %s", clr.Object)
	}

	if clr.Type != "comment" {
		return fmt.Errorf("type must be 'comment', got: %s", clr.Type)
	}

	// Validate all comments in results
	for i, comment := range clr.Results {
		if err := comment.Validate(); err != nil {
			return fmt.Errorf("invalid comment at index %d: %w", i, err)
		}
	}

	return nil
}
