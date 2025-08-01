package types

import (
	"encoding/json"
	"fmt"
)

// RichText represents rich text content in Notion, supporting text, mentions, and equations.
// Rich text is used throughout Notion for formatted content in pages, blocks, and properties.
type RichText struct {
	Type        RichTextType `json:"type"`
	Annotations *Annotations `json:"annotations,omitempty"`
	PlainText   string       `json:"plain_text"`
	Href        *string      `json:"href,omitempty"`
	Text        *TextContent `json:"text,omitempty"`
	Mention     *Mention     `json:"mention,omitempty"`
	Equation    *Equation    `json:"equation,omitempty"`
}

func (rt *RichText) GetText() string {
	if rt.Text != nil {
		return rt.Text.Content
	}
	return ""
}

// RichTextType represents the type of rich text element.
type RichTextType string

const (
	RichTextTypeText     RichTextType = "text"
	RichTextTypeMention  RichTextType = "mention"
	RichTextTypeEquation RichTextType = "equation"
)

// Annotations represent formatting options for rich text elements.
type Annotations struct {
	Bold          bool  `json:"bold"`
	Italic        bool  `json:"italic"`
	Strikethrough bool  `json:"strikethrough"`
	Underline     bool  `json:"underline"`
	Code          bool  `json:"code"`
	Color         Color `json:"color"`
}

// TextContent represents plain text content with optional link information.
type TextContent struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

// Link represents a hyperlink within rich text content.
type Link struct {
	URL string `json:"url"`
}

// Mention represents a mention of a user, page, database, date, or other entity.
// Mentions are created when users type @ followed by an entity name.
type Mention struct {
	Type            MentionType         `json:"type"`
	User            *PartialUser        `json:"user,omitempty"`
	Page            *PageReference      `json:"page,omitempty"`
	Database        *DatabaseReference  `json:"database,omitempty"`
	Date            *DateMention        `json:"date,omitempty"`
	LinkPreview     *LinkPreviewMention `json:"link_preview,omitempty"`
	LinkMention     *LinkMention        `json:"link_mention,omitempty"`
	TemplateMention *TemplateMention    `json:"template_mention,omitempty"`
	CustomEmoji     *CustomEmoji        `json:"custom_emoji,omitempty"`
}

// MentionType represents the type of entity being mentioned.
type MentionType string

const (
	MentionTypeUser            MentionType = "user"
	MentionTypePage            MentionType = "page"
	MentionTypeDatabase        MentionType = "database"
	MentionTypeDate            MentionType = "date"
	MentionTypeLinkPreview     MentionType = "link_preview"
	MentionTypeLinkMention     MentionType = "link_mention"
	MentionTypeTemplateMention MentionType = "template_mention"
	MentionTypeCustomEmoji     MentionType = "custom_emoji"
)

// PageReference represents a reference to a Notion page.
type PageReference struct {
	ID string `json:"id"`
}

// DatabaseReference represents a reference to a Notion database.
type DatabaseReference struct {
	ID string `json:"id"`
}

// DateMention represents a date mention in rich text.
type DateMention struct {
	Start    string  `json:"start"`
	End      *string `json:"end,omitempty"`
	TimeZone *string `json:"time_zone,omitempty"`
}

// LinkPreviewMention represents a link preview mention.
type LinkPreviewMention struct {
	URL string `json:"url"`
}

// LinkMention represents a rich link mention with metadata.
type LinkMention struct {
	Href         string  `json:"href"`
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	LinkAuthor   *string `json:"link_author,omitempty"`
	LinkProvider *string `json:"link_provider,omitempty"`
	ThumbnailURL *string `json:"thumbnail_url,omitempty"`
	IconURL      *string `json:"icon_url,omitempty"`
	IframeURL    *string `json:"iframe_url,omitempty"`
	Height       *int    `json:"height,omitempty"`
	Padding      *int    `json:"padding,omitempty"`
	PaddingTop   *int    `json:"padding_top,omitempty"`
}

// TemplateMention represents template variables like "today" or "me".
type TemplateMention struct {
	Type                TemplateMentionType `json:"type"`
	TemplateMentionDate *string             `json:"template_mention_date,omitempty"`
	TemplateMentionUser *string             `json:"template_mention_user,omitempty"`
}

// TemplateMentionType represents the type of template mention.
type TemplateMentionType string

const (
	TemplateMentionTypeDate TemplateMentionType = "template_mention_date"
	TemplateMentionTypeUser TemplateMentionType = "template_mention_user"
)

// CustomEmoji represents a custom emoji in a mention.
type CustomEmoji struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Equation represents a mathematical equation in LaTeX format.
type Equation struct {
	Expression string `json:"expression"`
}

// NewTextRichText creates a new rich text element with plain text content.
//
// Arguments:
// - content: The text content to display.
// - annotations: Optional formatting annotations (can be nil).
//
// Returns:
// - *RichText: A new rich text element with text content and specified formatting.
//
// Example:
//
//	richText := NewTextRichText("Hello, world!", &Annotations{Bold: true})
func NewTextRichText(content string, annotations *Annotations) *RichText {
	return &RichText{
		Type:        RichTextTypeText,
		Annotations: annotations,
		PlainText:   content,
		Text: &TextContent{
			Content: content,
		},
	}
}

// NewLinkRichText creates a new rich text element with a hyperlink.
//
// Arguments:
// - content: The displayed text content.
// - url: The URL to link to.
// - annotations: Optional formatting annotations (can be nil).
//
// Returns:
// - *RichText: A new rich text element with text content and a hyperlink.
//
// Example:
//
//	richText := NewLinkRichText("Visit Google", "https://google.com", nil)
func NewLinkRichText(content, url string, annotations *Annotations) *RichText {
	return &RichText{
		Type:        RichTextTypeText,
		Annotations: annotations,
		PlainText:   content,
		Href:        &url,
		Text: &TextContent{
			Content: content,
			Link: &Link{
				URL: url,
			},
		},
	}
}

// NewUserMention creates a new rich text element mentioning a user.
//
// Arguments:
// - userID: The ID of the user to mention.
//
// Returns:
// - *RichText: A new rich text element with a user mention.
//
// Example:
//
//	mention := NewUserMention("550e8400-e29b-41d4-a716-446655440000")
func NewUserMention(userID string) *RichText {
	return &RichText{
		Type:      RichTextTypeMention,
		PlainText: "@user",
		Mention: &Mention{
			Type: MentionTypeUser,
			User: &PartialUser{
				ID:     UserID(userID),
				Object: ObjectTypeUser,
			},
		},
	}
}

// ToPlainText extracts plain text content from a slice of RichText elements.
//
// Arguments:
// - richTexts: Slice of RichText elements to convert.
//
// Returns:
// - string: Concatenated plain text content from all elements.
//
// Example:
//
//	plainText := ToPlainText([]RichText{
//	    {PlainText: "Hello, "},
//	    {PlainText: "world!"},
//	})
//	// Result: "Hello, world!"
func ToPlainText(richTexts []RichText) string {
	if len(richTexts) == 0 {
		return ""
	}

	var result string
	for _, rt := range richTexts {
		result += rt.PlainText
	}
	return result
}

// Validate ensures the RichText has valid required fields based on its type.
//
// Returns:
// - error: Validation error if required fields are missing or invalid, nil if valid.
//
// Example:
//
//	richText := &RichText{Type: RichTextTypeText}
//	if err := richText.Validate(); err != nil {
//	    log.Printf("Invalid rich text: %v", err)
//	}
func (rt *RichText) Validate() error {
	switch rt.Type {
	case RichTextTypeText:
		if rt.Text == nil {
			return fmt.Errorf("text field is required for rich text type 'text'")
		}
		if rt.Text.Content == "" {
			return fmt.Errorf("text content cannot be empty")
		}
	case RichTextTypeMention:
		if rt.Mention == nil {
			return fmt.Errorf("mention field is required for rich text type 'mention'")
		}
		return rt.Mention.Validate()
	case RichTextTypeEquation:
		if rt.Equation == nil {
			return fmt.Errorf("equation field is required for rich text type 'equation'")
		}
		if rt.Equation.Expression == "" {
			return fmt.Errorf("equation expression cannot be empty")
		}
	default:
		return fmt.Errorf("unknown rich text type: %s", rt.Type)
	}
	return nil
}

// Validate ensures the Mention has valid required fields based on its type.
//
// Returns:
// - error: Validation error if required fields are missing, nil if valid.
func (m *Mention) Validate() error {
	switch m.Type {
	case MentionTypeUser:
		if m.User == nil {
			return fmt.Errorf("user field is required for mention type 'user'")
		}
	case MentionTypePage:
		if m.Page == nil {
			return fmt.Errorf("page field is required for mention type 'page'")
		}
	case MentionTypeDatabase:
		if m.Database == nil {
			return fmt.Errorf("database field is required for mention type 'database'")
		}
	case MentionTypeDate:
		if m.Date == nil {
			return fmt.Errorf("date field is required for mention type 'date'")
		}
		if m.Date.Start == "" {
			return fmt.Errorf("date start cannot be empty")
		}
	case MentionTypeLinkPreview:
		if m.LinkPreview == nil {
			return fmt.Errorf("link_preview field is required for mention type 'link_preview'")
		}
		if m.LinkPreview.URL == "" {
			return fmt.Errorf("link preview URL cannot be empty")
		}
	case MentionTypeTemplateMention:
		if m.TemplateMention == nil {
			return fmt.Errorf("template_mention field is required for mention type 'template_mention'")
		}
		return m.TemplateMention.Validate()
	default:
		return fmt.Errorf("unknown mention type: %s", m.Type)
	}
	return nil
}

// Validate ensures the TemplateMention has valid configuration.
//
// Returns:
// - error: Validation error if configuration is invalid, nil if valid.
func (tm *TemplateMention) Validate() error {
	switch tm.Type {
	case TemplateMentionTypeDate:
		if tm.TemplateMentionDate == nil {
			return fmt.Errorf("template_mention_date field is required for type 'template_mention_date'")
		}
		if *tm.TemplateMentionDate != "today" && *tm.TemplateMentionDate != "now" {
			return fmt.Errorf("template_mention_date must be 'today' or 'now', got: %s", *tm.TemplateMentionDate)
		}
	case TemplateMentionTypeUser:
		if tm.TemplateMentionUser == nil {
			return fmt.Errorf("template_mention_user field is required for type 'template_mention_user'")
		}
		if *tm.TemplateMentionUser != "me" {
			return fmt.Errorf("template_mention_user must be 'me', got: %s", *tm.TemplateMentionUser)
		}
	default:
		return fmt.Errorf("unknown template mention type: %s", tm.Type)
	}
	return nil
}

// UnmarshalJSON implements custom JSON unmarshaling for RichText to handle
// the polymorphic nature of different rich text types.
//
// Arguments:
// - data: Raw JSON bytes containing the rich text data.
//
// Returns:
// - error: Unmarshaling error if the JSON is malformed, nil if successful.
func (rt *RichText) UnmarshalJSON(data []byte) error {
	// First unmarshal into a temporary struct to get the type
	var temp struct {
		Type        RichTextType `json:"type"`
		Annotations *Annotations `json:"annotations,omitempty"`
		PlainText   string       `json:"plain_text"`
		Href        *string      `json:"href,omitempty"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	rt.Type = temp.Type
	rt.Annotations = temp.Annotations
	rt.PlainText = temp.PlainText
	rt.Href = temp.Href

	// Now unmarshal the type-specific fields
	var typeSpecific map[string]json.RawMessage
	if err := json.Unmarshal(data, &typeSpecific); err != nil {
		return err
	}

	switch rt.Type {
	case RichTextTypeText:
		if textData, ok := typeSpecific["text"]; ok {
			var text TextContent
			if err := json.Unmarshal(textData, &text); err != nil {
				return fmt.Errorf("failed to unmarshal text data: %w", err)
			}
			rt.Text = &text
		}
	case RichTextTypeMention:
		if mentionData, ok := typeSpecific["mention"]; ok {
			var mention Mention
			if err := json.Unmarshal(mentionData, &mention); err != nil {
				return fmt.Errorf("failed to unmarshal mention data: %w", err)
			}
			rt.Mention = &mention
		}
	case RichTextTypeEquation:
		if equationData, ok := typeSpecific["equation"]; ok {
			var equation Equation
			if err := json.Unmarshal(equationData, &equation); err != nil {
				return fmt.Errorf("failed to unmarshal equation data: %w", err)
			}
			rt.Equation = &equation
		}
	}

	return nil
}
