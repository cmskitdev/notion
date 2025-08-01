// package types provides types for Notion blocks.
package types

import (
	"encoding/json"
	"fmt"

	"github.com/cmskitdev/notion/id"
)

// BlockID represents a unique identifier for a Notion block.
type BlockID id.ID

// String returns the string representation of the BlockID.
//
// Returns:
// - string: The string representation of the BlockID.
func (p *BlockID) String() string {
	return string(*p)
}

// ParseBlockID parses a string as a Notion BlockID, accepting both dashed and undashed formats.
// It returns the BlockID in canonical dashed UUID format (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).
//
// Arguments:
// - id: The string to parse as a BlockID.
//
// Returns:
// - BlockID: The canonical, dashed BlockID.
// - error: If the input is not a valid 32-character hexadecimal string.
func ParseBlockID(id string) (BlockID, error) {
	blockID, err := IDParser.Parse(id)
	if err != nil {
		return "", err
	}
	return BlockID(blockID), nil
}

// UnmarshalJSON implements custom JSON unmarshaling for BlockID.
// It accepts both dashed and undashed UUID formats and converts them to canonical format.
//
// Arguments:
// - data: Raw JSON bytes containing the ID string.
//
// Returns:
// - error: Parsing error if the ID format is invalid, nil if successful.
func (p *BlockID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := ParseBlockID(s)
	if err != nil {
		return fmt.Errorf("invalid BlockID in JSON: %w", err)
	}

	*p = parsed
	return nil
}

// MarshalJSON implements custom JSON marshaling for BlockID.
//
// Returns:
// - []byte: JSON-encoded BlockID as a string.
// - error: Marshaling error, if any.
func (p BlockID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

// Validate ensures the BlockID is in the correct UUID format.
//
// Returns:
// - error: Validation error if the ID format is invalid, nil if valid.
//
// Example:
//
//	if err := blockID.Validate(); err != nil {
//	    log.Printf("Invalid BlockID: %v", err)
//	}
func (p BlockID) Validate() error {
	_, err := ParseBlockID(string(p))
	return err
}

// Block represents a block in Notion, which is the basic building unit of pages.
// Blocks can contain text, media, embeds, and other structured content.
type Block struct {
	BaseObject
	ID           BlockID   `json:"id"`
	Type         BlockType `json:"type"`
	HasChildren  bool      `json:"has_children"`
	Archived     bool      `json:"archived,omitempty"`
	Parent       *Parent   `json:"parent,omitempty"`
	CreatedBy    *User     `json:"created_by,omitempty"`
	LastEditedBy *User     `json:"last_edited_by,omitempty"`

	// Block type-specific content - only one should be set based on Type
	Paragraph        *RichTextBlock        `json:"paragraph,omitempty"`
	Heading1         *HeadingBlock         `json:"heading_1,omitempty"`
	Heading2         *HeadingBlock         `json:"heading_2,omitempty"`
	Heading3         *HeadingBlock         `json:"heading_3,omitempty"`
	BulletedListItem *ListItemBlock        `json:"bulleted_list_item,omitempty"`
	NumberedListItem *ListItemBlock        `json:"numbered_list_item,omitempty"`
	ToDo             *ToDoBlock            `json:"to_do,omitempty"`
	Toggle           *ToggleBlock          `json:"toggle,omitempty"`
	ChildPage        *ChildPageBlock       `json:"child_page,omitempty"`
	ChildDatabase    *ChildDatabaseBlock   `json:"child_database,omitempty"`
	Embed            *EmbedBlock           `json:"embed,omitempty"`
	Image            *FileBlock            `json:"image,omitempty"`
	Video            *FileBlock            `json:"video,omitempty"`
	File             *FileBlock            `json:"file,omitempty"`
	PDF              *FileBlock            `json:"pdf,omitempty"`
	Bookmark         *BookmarkBlock        `json:"bookmark,omitempty"`
	Callout          *CalloutBlock         `json:"callout,omitempty"`
	Quote            *RichTextBlock        `json:"quote,omitempty"`
	Equation         *EquationBlock        `json:"equation,omitempty"`
	Divider          *DividerBlock         `json:"divider,omitempty"`
	TableOfContents  *TableOfContentsBlock `json:"table_of_contents,omitempty"`
	Column           *ColumnBlock          `json:"column,omitempty"`
	ColumnList       *ColumnListBlock      `json:"column_list,omitempty"`
	LinkPreview      *LinkPreviewBlock     `json:"link_preview,omitempty"`
	SyncedBlock      *SyncedBlock          `json:"synced_block,omitempty"`
	Template         *TemplateBlock        `json:"template,omitempty"`
	LinkToPage       *LinkToPageBlock      `json:"link_to_page,omitempty"`
	Table            *TableBlock           `json:"table,omitempty"`
	TableRow         *TableRowBlock        `json:"table_row,omitempty"`
	Audio            *FileBlock            `json:"audio,omitempty"`
	Code             *CodeBlock            `json:"code,omitempty"`
	Breadcrumb       *BreadcrumbBlock      `json:"breadcrumb,omitempty"`
}

// BlockType represents the different types of blocks available in Notion.
type BlockType string

const (
	BlockTypeParagraph        BlockType = "paragraph"
	BlockTypeHeading1         BlockType = "heading_1"
	BlockTypeHeading2         BlockType = "heading_2"
	BlockTypeHeading3         BlockType = "heading_3"
	BlockTypeBulletedListItem BlockType = "bulleted_list_item"
	BlockTypeNumberedListItem BlockType = "numbered_list_item"
	BlockTypeToDo             BlockType = "to_do"
	BlockTypeToggle           BlockType = "toggle"
	BlockTypeChildPage        BlockType = "child_page"
	BlockTypeChildDatabase    BlockType = "child_database"
	BlockTypeEmbed            BlockType = "embed"
	BlockTypeImage            BlockType = "image"
	BlockTypeVideo            BlockType = "video"
	BlockTypeFile             BlockType = "file"
	BlockTypePDF              BlockType = "pdf"
	BlockTypeBookmark         BlockType = "bookmark"
	BlockTypeCallout          BlockType = "callout"
	BlockTypeQuote            BlockType = "quote"
	BlockTypeEquation         BlockType = "equation"
	BlockTypeDivider          BlockType = "divider"
	BlockTypeTableOfContents  BlockType = "table_of_contents"
	BlockTypeColumn           BlockType = "column"
	BlockTypeColumnList       BlockType = "column_list"
	BlockTypeLinkPreview      BlockType = "link_preview"
	BlockTypeSyncedBlock      BlockType = "synced_block"
	BlockTypeTemplate         BlockType = "template"
	BlockTypeLinkToPage       BlockType = "link_to_page"
	BlockTypeTable            BlockType = "table"
	BlockTypeTableRow         BlockType = "table_row"
	BlockTypeAudio            BlockType = "audio"
	BlockTypeCode             BlockType = "code"
	BlockTypeBreadcrumb       BlockType = "breadcrumb"
	BlockTypeUnsupported      BlockType = "unsupported"
)

// RichTextBlock represents blocks that contain rich text content (paragraph, quote).
type RichTextBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// HeadingBlock represents heading blocks (h1, h2, h3).
type HeadingBlock struct {
	RichText     []RichText `json:"rich_text"`
	Color        Color      `json:"color,omitempty"`
	IsToggleable bool       `json:"is_toggleable,omitempty"`
	Children     []Block    `json:"children,omitempty"`
}

// ListItemBlock represents list item blocks (bulleted and numbered).
type ListItemBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// ToDoBlock represents a to-do/checkbox block.
type ToDoBlock struct {
	RichText []RichText `json:"rich_text"`
	Checked  bool       `json:"checked"`
	Color    Color      `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// ToggleBlock represents a toggle block that can be expanded/collapsed.
type ToggleBlock struct {
	RichText []RichText `json:"rich_text"`
	Color    Color      `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// ChildPageBlock represents a link to a child page.
type ChildPageBlock struct {
	Title string `json:"title"`
}

// ChildDatabaseBlock represents a link to a child database.
type ChildDatabaseBlock struct {
	Title string `json:"title"`
}

// EmbedBlock represents an embedded content block.
type EmbedBlock struct {
	URL     string     `json:"url"`
	Caption []RichText `json:"caption,omitempty"`
}

// FileBlock represents file, image, video, audio, and PDF blocks.
type FileBlock struct {
	Type     FileBlockType `json:"type,omitempty"`
	File     *File         `json:"file,omitempty"`
	External *File         `json:"external,omitempty"`
	Caption  []RichText    `json:"caption,omitempty"`
	Name     string        `json:"name,omitempty"`
}

// FileBlockType represents the source type of a file block.
type FileBlockType string

const (
	FileBlockTypeFile     FileBlockType = "file"
	FileBlockTypeExternal FileBlockType = "external"
)

// BookmarkBlock represents a bookmark to an external URL.
type BookmarkBlock struct {
	URL     string     `json:"url"`
	Caption []RichText `json:"caption,omitempty"`
}

// CalloutBlock represents a callout block with an icon and text.
type CalloutBlock struct {
	RichText []RichText `json:"rich_text"`
	Icon     *Icon      `json:"icon,omitempty"`
	Color    Color      `json:"color,omitempty"`
	Children []Block    `json:"children,omitempty"`
}

// EquationBlock represents a block-level mathematical equation.
type EquationBlock struct {
	Expression string `json:"expression"`
}

// DividerBlock represents a horizontal divider.
type DividerBlock struct {
	// Divider blocks have no additional properties
}

// TableOfContentsBlock represents an automatically generated table of contents.
type TableOfContentsBlock struct {
	Color Color `json:"color,omitempty"`
}

// ColumnBlock represents a column within a column list.
type ColumnBlock struct {
	Children []Block `json:"children,omitempty"`
}

// ColumnListBlock represents a container for columns.
type ColumnListBlock struct {
	Children []Block `json:"children,omitempty"`
}

// LinkPreviewBlock represents a link preview.
type LinkPreviewBlock struct {
	URL string `json:"url"`
}

// SyncedBlock represents a synced block that can be referenced elsewhere.
type SyncedBlock struct {
	SyncedFrom *SyncedFromBlock `json:"synced_from,omitempty"`
	Children   []Block          `json:"children,omitempty"`
}

// SyncedFromBlock represents the source of a synced block.
type SyncedFromBlock struct {
	BlockID *BlockID `json:"block_id,omitempty"`
}

// TemplateBlock represents a template block.
type TemplateBlock struct {
	RichText []RichText `json:"rich_text"`
	Children []Block    `json:"children,omitempty"`
}

// LinkToPageBlock represents a link to another page.
type LinkToPageBlock struct {
	Type       LinkToPageType `json:"type"`
	PageID     *PageID        `json:"page_id,omitempty"`
	DatabaseID *DatabaseID    `json:"database_id,omitempty"`
}

// LinkToPageType represents the type of page being linked to.
type LinkToPageType string

const (
	LinkToPageTypePage     LinkToPageType = "page_id"
	LinkToPageTypeDatabase LinkToPageType = "database_id"
)

// TableBlock represents a table block.
type TableBlock struct {
	TableWidth      int     `json:"table_width"`
	HasColumnHeader bool    `json:"has_column_header"`
	HasRowHeader    bool    `json:"has_row_header"`
	Children        []Block `json:"children,omitempty"`
}

// TableRowBlock represents a row within a table.
type TableRowBlock struct {
	Cells [][]RichText `json:"cells"`
}

// CodeBlock represents a code block with syntax highlighting.
type CodeBlock struct {
	RichText []RichText `json:"rich_text"`
	Caption  []RichText `json:"caption,omitempty"`
	Language string     `json:"language,omitempty"`
}

// BreadcrumbBlock represents a breadcrumb navigation element.
type BreadcrumbBlock struct {
	// Breadcrumb blocks have no additional properties
}

// NewParagraphBlock creates a new paragraph block with the given text content.
//
// Arguments:
// - content: The paragraph text content as rich text elements.
//
// Returns:
// - *Block: A new block with paragraph type and specified content.
//
// Example:
//
//	block := NewParagraphBlock([]RichText{*NewTextRichText("Hello, world!", nil)})
func NewParagraphBlock(content []RichText) *Block {
	return &Block{
		BaseObject: BaseObject{
			Object: ObjectTypeBlock,
		},
		Type: BlockTypeParagraph,
		Paragraph: &RichTextBlock{
			RichText: content,
		},
	}
}

// NewHeading1Block creates a new heading 1 block with the given text content.
//
// Arguments:
// - content: The heading text content as rich text elements.
//
// Returns:
// - *Block: A new block with heading_1 type and specified content.
//
// Example:
//
//	block := NewHeading1Block([]RichText{*NewTextRichText("Chapter 1", nil)})
func NewHeading1Block(content []RichText) *Block {
	return &Block{
		BaseObject: BaseObject{
			Object: ObjectTypeBlock,
		},
		Type: BlockTypeHeading1,
		Heading1: &HeadingBlock{
			RichText: content,
		},
	}
}

// NewToDoBlock creates a new to-do block with the given text content and checked state.
//
// Arguments:
// - content: The to-do text content as rich text elements.
// - checked: Whether the to-do is checked/completed.
//
// Returns:
// - *Block: A new block with to_do type and specified content and state.
//
// Example:
//
//	block := NewToDoBlock([]RichText{*NewTextRichText("Buy groceries", nil)}, false)
func NewToDoBlock(content []RichText, checked bool) *Block {
	return &Block{
		BaseObject: BaseObject{
			Object: ObjectTypeBlock,
		},
		Type: BlockTypeToDo,
		ToDo: &ToDoBlock{
			RichText: content,
			Checked:  checked,
		},
	}
}

// NewCodeBlock creates a new code block with the given content and language.
//
// Arguments:
// - content: The code content as rich text elements.
// - language: The programming language for syntax highlighting (can be empty).
//
// Returns:
// - *Block: A new block with code type and specified content and language.
//
// Example:
//
//	code := []RichText{*NewTextRichText("fmt.Println(\"Hello, world!\")", nil)}
//	block := NewCodeBlock(code, "go")
func NewCodeBlock(content []RichText, language string) *Block {
	return &Block{
		BaseObject: BaseObject{
			Object: ObjectTypeBlock,
		},
		Type: BlockTypeCode,
		Code: &CodeBlock{
			RichText: content,
			Language: language,
		},
	}
}

// GetContent returns the appropriate content for the block based on its type.
//
// Returns:
// - interface{}: The block content, or nil if not applicable.
//
// Example:
//
//	content := block.GetContent()
//	if richText, ok := content.(*RichTextBlock); ok {
//	    fmt.Println("Text:", ToPlainText(richText.RichText))
//	}
func (b *Block) GetContent() interface{} {
	switch b.Type {
	case BlockTypeParagraph:
		return b.Paragraph
	case BlockTypeHeading1:
		return b.Heading1
	case BlockTypeHeading2:
		return b.Heading2
	case BlockTypeHeading3:
		return b.Heading3
	case BlockTypeBulletedListItem:
		return b.BulletedListItem
	case BlockTypeNumberedListItem:
		return b.NumberedListItem
	case BlockTypeToDo:
		return b.ToDo
	case BlockTypeToggle:
		return b.Toggle
	case BlockTypeChildPage:
		return b.ChildPage
	case BlockTypeChildDatabase:
		return b.ChildDatabase
	case BlockTypeEmbed:
		return b.Embed
	case BlockTypeImage:
		return b.Image
	case BlockTypeVideo:
		return b.Video
	case BlockTypeFile:
		return b.File
	case BlockTypePDF:
		return b.PDF
	case BlockTypeBookmark:
		return b.Bookmark
	case BlockTypeCallout:
		return b.Callout
	case BlockTypeQuote:
		return b.Quote
	case BlockTypeEquation:
		return b.Equation
	case BlockTypeDivider:
		return b.Divider
	case BlockTypeTableOfContents:
		return b.TableOfContents
	case BlockTypeColumn:
		return b.Column
	case BlockTypeColumnList:
		return b.ColumnList
	case BlockTypeLinkPreview:
		return b.LinkPreview
	case BlockTypeSyncedBlock:
		return b.SyncedBlock
	case BlockTypeTemplate:
		return b.Template
	case BlockTypeLinkToPage:
		return b.LinkToPage
	case BlockTypeTable:
		return b.Table
	case BlockTypeTableRow:
		return b.TableRow
	case BlockTypeAudio:
		return b.Audio
	case BlockTypeCode:
		return b.Code
	case BlockTypeBreadcrumb:
		return b.Breadcrumb
	}
	return nil
}

// GetRichText returns the rich text content for blocks that support it.
//
// Returns:
// - []RichText: The rich text content, or nil if the block type doesn't support rich text.
//
// Example:
//
//	richText := block.GetRichText()
//	if richText != nil {
//	    fmt.Println("Text:", ToPlainText(richText))
//	}
func (b *Block) GetRichText() []RichText {
	switch b.Type {
	case BlockTypeParagraph:
		if b.Paragraph != nil {
			return b.Paragraph.RichText
		}
	case BlockTypeHeading1:
		if b.Heading1 != nil {
			return b.Heading1.RichText
		}
	case BlockTypeHeading2:
		if b.Heading2 != nil {
			return b.Heading2.RichText
		}
	case BlockTypeHeading3:
		if b.Heading3 != nil {
			return b.Heading3.RichText
		}
	case BlockTypeBulletedListItem:
		if b.BulletedListItem != nil {
			return b.BulletedListItem.RichText
		}
	case BlockTypeNumberedListItem:
		if b.NumberedListItem != nil {
			return b.NumberedListItem.RichText
		}
	case BlockTypeToDo:
		if b.ToDo != nil {
			return b.ToDo.RichText
		}
	case BlockTypeToggle:
		if b.Toggle != nil {
			return b.Toggle.RichText
		}
	case BlockTypeCallout:
		if b.Callout != nil {
			return b.Callout.RichText
		}
	case BlockTypeQuote:
		if b.Quote != nil {
			return b.Quote.RichText
		}
	case BlockTypeTemplate:
		if b.Template != nil {
			return b.Template.RichText
		}
	case BlockTypeCode:
		if b.Code != nil {
			return b.Code.RichText
		}
	}
	return nil
}

// Validate ensures the Block has valid required fields based on its type.
//
// Returns:
// - error: Validation error if required fields are missing or invalid, nil if valid.
//
// Example:
//
//	if err := block.Validate(); err != nil {
//	    log.Printf("Invalid block: %v", err)
//	}
func (b *Block) Validate() error {
	if b.Object != ObjectTypeBlock {
		return fmt.Errorf("object type must be 'block', got: %s", b.Object)
	}

	switch b.Type {
	case BlockTypeParagraph:
		if b.Paragraph == nil {
			return fmt.Errorf("paragraph field is required for block type 'paragraph'")
		}
		return validateRichTextBlock(b.Paragraph)
	case BlockTypeHeading1, BlockTypeHeading2, BlockTypeHeading3:
		var heading *HeadingBlock
		switch b.Type {
		case BlockTypeHeading1:
			heading = b.Heading1
		case BlockTypeHeading2:
			heading = b.Heading2
		case BlockTypeHeading3:
			heading = b.Heading3
		}
		if heading == nil {
			return fmt.Errorf("%s field is required for block type '%s'", b.Type, b.Type)
		}
		return validateHeadingBlock(heading)
	case BlockTypeToDo:
		if b.ToDo == nil {
			return fmt.Errorf("to_do field is required for block type 'to_do'")
		}
		return validateToDoBlock(b.ToDo)
	case BlockTypeCode:
		if b.Code == nil {
			return fmt.Errorf("code field is required for block type 'code'")
		}
		return validateCodeBlock(b.Code)
	case BlockTypeEmbed:
		if b.Embed == nil {
			return fmt.Errorf("embed field is required for block type 'embed'")
		}
		return validateEmbedBlock(b.Embed)
	case BlockTypeBookmark:
		if b.Bookmark == nil {
			return fmt.Errorf("bookmark field is required for block type 'bookmark'")
		}
		return validateBookmarkBlock(b.Bookmark)
	case BlockTypeEquation:
		if b.Equation == nil {
			return fmt.Errorf("equation field is required for block type 'equation'")
		}
		return validateEquationBlock(b.Equation)
	case BlockTypeChildPage:
		if b.ChildPage == nil {
			return fmt.Errorf("child_page field is required for block type 'child_page'")
		}
		if b.ChildPage.Title == "" {
			return fmt.Errorf("child page title cannot be empty")
		}
	case BlockTypeChildDatabase:
		if b.ChildDatabase == nil {
			return fmt.Errorf("child_database field is required for block type 'child_database'")
		}
		if b.ChildDatabase.Title == "" {
			return fmt.Errorf("child database title cannot be empty")
		}
	}

	return nil
}

// Helper validation functions
func validateRichTextBlock(rtb *RichTextBlock) error {
	for i, rt := range rtb.RichText {
		if err := rt.Validate(); err != nil {
			return fmt.Errorf("invalid rich text at index %d: %w", i, err)
		}
	}
	return nil
}

func validateHeadingBlock(hb *HeadingBlock) error {
	if len(hb.RichText) == 0 {
		return fmt.Errorf("heading block must have at least one rich text element")
	}
	for i, rt := range hb.RichText {
		if err := rt.Validate(); err != nil {
			return fmt.Errorf("invalid rich text at index %d: %w", i, err)
		}
	}
	return nil
}

func validateToDoBlock(tdb *ToDoBlock) error {
	for i, rt := range tdb.RichText {
		if err := rt.Validate(); err != nil {
			return fmt.Errorf("invalid rich text at index %d: %w", i, err)
		}
	}
	return nil
}

func validateCodeBlock(cb *CodeBlock) error {
	for i, rt := range cb.RichText {
		if err := rt.Validate(); err != nil {
			return fmt.Errorf("invalid rich text at index %d: %w", i, err)
		}
	}
	return nil
}

func validateEmbedBlock(eb *EmbedBlock) error {
	if eb.URL == "" {
		return fmt.Errorf("embed URL cannot be empty")
	}
	return nil
}

func validateBookmarkBlock(bb *BookmarkBlock) error {
	if bb.URL == "" {
		return fmt.Errorf("bookmark URL cannot be empty")
	}
	return nil
}

func validateEquationBlock(eb *EquationBlock) error {
	if eb.Expression == "" {
		return fmt.Errorf("equation expression cannot be empty")
	}
	return nil
}
