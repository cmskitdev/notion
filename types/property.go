package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cmskitdev/notion/id"
)

// PropertyID represents a unique identifier for a Notion property.
type PropertyID id.ID

// String returns the string representation of the PropertyID.
//
// Returns:
// - string: The string representation of the PropertyID.
func (p *PropertyID) String() string {
	return string(*p)
}

// ParsePropertyID parses a string as a Notion PropertyID.
// PropertyIDs can be either UUID format (32-char hex with/without dashes) or simple string identifiers.
//
// Arguments:
// - id: The string to parse as a PropertyID.
//
// Returns:
// - PropertyID: The PropertyID, formatted as UUID if applicable, otherwise as-is.
// - error: If the input is empty.
func ParsePropertyID(id string) (PropertyID, error) {
	if id == "" {
		return "", fmt.Errorf("PropertyID cannot be empty")
	}

	// Try to parse as UUID first
	if len(strings.ReplaceAll(id, "-", "")) == 32 {
		propertyID, err := IDParser.Parse(id)
		if err == nil {
			return PropertyID(propertyID), nil
		}
	}

	// Return as-is for simple string identifiers
	return PropertyID(id), nil
}

// UnmarshalJSON implements custom JSON unmarshaling for PropertyID.
// It accepts UUID formats, simple string identifiers, or empty strings.
//
// Arguments:
// - data: Raw JSON bytes containing the ID string.
//
// Returns:
// - error: Parsing error if the JSON is malformed, nil if successful.
func (p *PropertyID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	// Allow empty PropertyID
	if s == "" {
		*p = PropertyID("")
		return nil
	}

	parsed, err := ParsePropertyID(s)
	if err != nil {
		return fmt.Errorf("invalid PropertyID in JSON: %w", err)
	}

	*p = parsed
	return nil
}

// MarshalJSON implements custom JSON marshaling for PropertyID.
//
// Returns:
// - []byte: JSON-encoded PropertyID as a string.
// - error: Marshaling error, if any.
func (p PropertyID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

// Validate ensures the PropertyID is valid.
// PropertyIDs can be UUID format or simple string identifiers, but cannot be empty.
//
// Returns:
// - error: Validation error if the ID is empty, nil if valid.
//
// Example:
//
//	if err := propertyID.Validate(); err != nil {
//	    log.Printf("Invalid PropertyID: %v", err)
//	}
func (p PropertyID) Validate() error {
	if string(p) == "" {
		return fmt.Errorf("PropertyID cannot be empty")
	}
	return nil
}

// Property represents a property within a database or page.
// Properties define the schema and store values for database columns.
type Property struct {
	ID   PropertyID   `json:"id"`
	Type PropertyType `json:"type"`
	Name string       `json:"name,omitempty"`

	// Value fields - only one should be set based on Type
	Title          []RichText            `json:"title,omitempty"`
	RichText       []RichText            `json:"rich_text,omitempty"`
	Number         *NumberProperty       `json:"number,omitempty"`
	Select         *SelectProperty       `json:"select,omitempty"`
	MultiSelect    []SelectOption        `json:"multi_select,omitempty"`
	Date           *DateProperty         `json:"date,omitempty"`
	People         []User                `json:"people,omitempty"`
	Files          []FileProperty        `json:"files,omitempty"`
	Checkbox       *bool                 `json:"checkbox,omitempty"`
	URL            *string               `json:"url,omitempty"`
	Email          *string               `json:"email,omitempty"`
	PhoneNumber    *string               `json:"phone_number,omitempty"`
	Formula        *FormulaProperty      `json:"formula,omitempty"`
	Relation       []RelationProperty    `json:"relation,omitempty"`
	Rollup         *RollupProperty       `json:"rollup,omitempty"`
	CreatedTime    *Timestamp            `json:"created_time,omitempty"`
	CreatedBy      *User                 `json:"created_by,omitempty"`
	LastEditedTime *Timestamp            `json:"last_edited_time,omitempty"`
	LastEditedBy   *User                 `json:"last_edited_by,omitempty"`
	Status         *StatusProperty       `json:"status,omitempty"`
	UniqueID       *UniqueIDProperty     `json:"unique_id,omitempty"`
	Verification   *VerificationProperty `json:"verification,omitempty"`
}

// PropertyType represents the different types of properties available in Notion.
type PropertyType string

// PropertyType represents the different types of properties available in Notion.
//
// See: https://developers.notion.com/reference/block
const (
	PropertyTypeCheckbox       PropertyType = "checkbox"
	PropertyTypeCreatedBy      PropertyType = "created_by"
	PropertyTypeCreatedTime    PropertyType = "created_time"
	PropertyTypeDate           PropertyType = "date"
	PropertyTypeEmail          PropertyType = "email"
	PropertyTypeFiles          PropertyType = "files"
	PropertyTypeFormula        PropertyType = "formula"
	PropertyTypeLastEditedBy   PropertyType = "last_edited_by"
	PropertyTypeLastEditedTime PropertyType = "last_edited_time"
	PropertyTypeMultiSelect    PropertyType = "multi_select"
	PropertyTypeNumber         PropertyType = "number"
	PropertyTypePeople         PropertyType = "people"
	PropertyTypePhoneNumber    PropertyType = "phone_number"
	PropertyTypeRelation       PropertyType = "relation"
	PropertyTypeRichText       PropertyType = "rich_text"
	PropertyTypeRollup         PropertyType = "rollup"
	PropertyTypeSelect         PropertyType = "select"
	PropertyTypeStatus         PropertyType = "status"
	PropertyTypeTitle          PropertyType = "title"
	PropertyTypeUniqueID       PropertyType = "unique_id"
	PropertyTypeURL            PropertyType = "url"
	PropertyTypeVerification   PropertyType = "verification"
)

// NumberProperty represents a number property value.
type NumberProperty struct {
	Number *float64 `json:"number"`
}

// UnmarshalJSON implements custom JSON unmarshaling for NumberProperty.
// It handles both direct number values and the structured format.
//
// Arguments:
// - data: Raw JSON bytes containing the number data.
//
// Returns:
// - error: Unmarshaling error if the JSON is malformed, nil if successful.
func (np *NumberProperty) UnmarshalJSON(data []byte) error {
	// Check for null first
	if string(data) == "null" {
		np.Number = nil
		return nil
	}

	// Try to unmarshal as a direct number first
	var directNumber float64
	if err := json.Unmarshal(data, &directNumber); err == nil {
		np.Number = &directNumber
		return nil
	}

	// Try to unmarshal as structured format
	type Alias NumberProperty
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(np),
	}
	return json.Unmarshal(data, &aux)
}

// SelectProperty represents a select property value.
type SelectProperty struct {
	ID    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Color *Color  `json:"color,omitempty"`
}

// SelectOption represents an option in a select or multi-select property.
type SelectOption struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color Color  `json:"color"`
}

// DateProperty represents a date property value with optional end date.
type DateProperty struct {
	Start    string  `json:"start"`
	End      *string `json:"end,omitempty"`
	TimeZone *string `json:"time_zone,omitempty"`
}

// FileProperty represents a file within a files property.
type FileProperty struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	File     *File  `json:"file,omitempty"`
	External *File  `json:"external,omitempty"`
}

// FormulaProperty represents the result of a formula property.
type FormulaProperty struct {
	Type    FormulaResultType `json:"type"`
	String  *string           `json:"string,omitempty"`
	Number  *float64          `json:"number,omitempty"`
	Boolean *bool             `json:"boolean,omitempty"`
	Date    *DateProperty     `json:"date,omitempty"`
}

// FormulaResultType represents the type of result from a formula.
type FormulaResultType string

const (
	FormulaResultTypeString  FormulaResultType = "string"
	FormulaResultTypeNumber  FormulaResultType = "number"
	FormulaResultTypeBoolean FormulaResultType = "boolean"
	FormulaResultTypeDate    FormulaResultType = "date"
)

// RelationProperty represents a relation to another database.
type RelationProperty struct {
	ID string `json:"id"`
}

// RollupProperty represents a rollup property that aggregates values from relations.
type RollupProperty struct {
	Type       RollupType    `json:"type"`
	Number     *float64      `json:"number,omitempty"`
	Date       *DateProperty `json:"date,omitempty"`
	Array      []Property    `json:"array,omitempty"`
	Incomplete *bool         `json:"incomplete,omitempty"`
}

// RollupType represents the type of rollup aggregation.
type RollupType string

const (
	RollupTypeNumber      RollupType = "number"
	RollupTypeDate        RollupType = "date"
	RollupTypeArray       RollupType = "array"
	RollupTypeUnsupported RollupType = "unsupported"
	RollupTypeIncomplete  RollupType = "incomplete"
)

// StatusProperty represents a status property value.
type StatusProperty struct {
	ID    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Color *Color  `json:"color,omitempty"`
}

// UniqueIDProperty represents a unique ID property value.
type UniqueIDProperty struct {
	Number *int    `json:"number,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
}

// VerificationProperty represents a verification property state.
type VerificationProperty struct {
	State      VerificationState `json:"state"`
	Date       *DateProperty     `json:"date,omitempty"`
	VerifiedBy *User             `json:"verified_by,omitempty"`
}

// VerificationState represents the state of a verification property.
type VerificationState string

const (
	VerificationStateUnverified VerificationState = "unverified"
	VerificationStateVerified   VerificationState = "verified"
	VerificationStateExpired    VerificationState = "expired"
)

// NewTitleProperty creates a new title property with the given text content.
//
// Arguments:
// - content: The title text content.
//
// Returns:
// - *Property: A new property with title type and specified content.
//
// Example:
//
//	prop := NewTitleProperty("My Page Title")
func NewTitleProperty(content string) *Property {
	return &Property{
		ID:    PropertyID("title"),
		Type:  PropertyTypeTitle,
		Title: []RichText{*NewTextRichText(content, nil)},
	}
}

// NewRichTextProperty creates a new rich text property with the given content.
//
// Arguments:
// - content: The rich text content.
//
// Returns:
// - *Property: A new property with rich text type and specified content.
//
// Example:
//
//	prop := NewRichTextProperty([]RichText{*NewTextRichText("Some text", nil)})
func NewRichTextProperty(content []RichText) *Property {
	return &Property{
		ID:       PropertyID("rich_text"),
		Type:     PropertyTypeRichText,
		RichText: content,
	}
}

// NewNumberProperty creates a new number property with the given value.
//
// Arguments:
// - value: The numeric value (can be nil for empty).
//
// Returns:
// - *Property: A new property with number type and specified value.
//
// Example:
//
//	value := 42.5
//	prop := NewNumberProperty(&value)
func NewNumberProperty(value *float64) *Property {
	var numberProp *NumberProperty
	if value != nil {
		numberProp = &NumberProperty{Number: value}
	}
	return &Property{
		ID:     PropertyID("number"),
		Type:   PropertyTypeNumber,
		Number: numberProp,
	}
}

// NewSelectProperty creates a new select property with the given option.
//
// Arguments:
// - optionName: The name of the selected option (can be nil for empty).
// - optionColor: The color of the selected option (can be nil).
//
// Returns:
// - *Property: A new property with select type and specified option.
//
// Example:
//
//	prop := NewSelectProperty(&"Option A", &ColorBlue)
func NewSelectProperty(optionName *string, optionColor *Color) *Property {
	var selectProp *SelectProperty
	if optionName != nil {
		selectProp = &SelectProperty{
			Name:  optionName,
			Color: optionColor,
		}
	}
	return &Property{
		ID:     PropertyID("select"),
		Type:   PropertyTypeSelect,
		Select: selectProp,
	}
}

// NewCheckboxProperty creates a new checkbox property with the given state.
//
// Arguments:
// - checked: Whether the checkbox is checked.
//
// Returns:
// - *Property: A new property with checkbox type and specified state.
//
// Example:
//
//	prop := NewCheckboxProperty(true)
func NewCheckboxProperty(checked bool) *Property {
	return &Property{
		ID:       PropertyID("checkbox"),
		Type:     PropertyTypeCheckbox,
		Checkbox: &checked,
	}
}

// GetValue returns the appropriate value for the property based on its type.
//
// Returns:
// - interface{}: The property value, or nil if not set.
//
// Example:
//
//	value := property.GetValue()
//	if title, ok := value.([]RichText); ok {
//	    fmt.Println("Title:", ToPlainText(title))
//	}

func (p *Property) GetValue() interface{} {
	switch p.Type {
	case PropertyTypeTitle:
		return p.Title
	case PropertyTypeRichText:
		return p.RichText
	case PropertyTypeNumber:
		if p.Number != nil {
			return p.Number.Number
		}
	case PropertyTypeSelect:
		return p.Select
	case PropertyTypeMultiSelect:
		return p.MultiSelect
	case PropertyTypeDate:
		return p.Date
	case PropertyTypePeople:
		return p.People
	case PropertyTypeFiles:
		return p.Files
	case PropertyTypeCheckbox:
		return p.Checkbox
	case PropertyTypeURL:
		return p.URL
	case PropertyTypeEmail:
		return p.Email
	case PropertyTypePhoneNumber:
		return p.PhoneNumber
	case PropertyTypeFormula:
		return p.Formula
	case PropertyTypeRelation:
		return p.Relation
	case PropertyTypeRollup:
		return p.Rollup
	case PropertyTypeCreatedTime:
		return p.CreatedTime
	case PropertyTypeCreatedBy:
		return p.CreatedBy
	case PropertyTypeLastEditedTime:
		return p.LastEditedTime
	case PropertyTypeLastEditedBy:
		return p.LastEditedBy
	case PropertyTypeStatus:
		return p.Status
	case PropertyTypeUniqueID:
		return p.UniqueID
	case PropertyTypeVerification:
		return p.Verification
	}
	return nil
}

func GetProperty[T any](key string, properties map[string]Property) *T {
	var property T
	for _, p := range properties {
		if p.ID.String() == key {
			property = p.GetValue().(T)
		}
	}
	return &property
}

// Validate ensures the Property has valid required fields based on its type.
//
// Returns:
// - error: Validation error if required fields are missing or invalid, nil if valid.
//
// Example:
//
//	if err := property.Validate(); err != nil {
//	    log.Printf("Invalid property: %v", err)
//	}
func (p *Property) Validate() error {
	switch p.Type {
	case PropertyTypeTitle:
		if len(p.Title) == 0 {
			return fmt.Errorf("title property must have at least one rich text element")
		}
		for i, rt := range p.Title {
			if err := rt.Validate(); err != nil {
				return fmt.Errorf("invalid rich text at index %d: %w", i, err)
			}
		}
	case PropertyTypeRichText:
		for i, rt := range p.RichText {
			if err := rt.Validate(); err != nil {
				return fmt.Errorf("invalid rich text at index %d: %w", i, err)
			}
		}
	case PropertyTypeDate:
		if p.Date != nil && p.Date.Start == "" {
			return fmt.Errorf("date property start cannot be empty")
		}
	case PropertyTypeSelect:
		if p.Select != nil && p.Select.Name != nil && *p.Select.Name == "" {
			return fmt.Errorf("select property name cannot be empty")
		}
	case PropertyTypeMultiSelect:
		for i, option := range p.MultiSelect {
			if option.Name == "" {
				return fmt.Errorf("multi-select option at index %d cannot have empty name", i)
			}
		}
	case PropertyTypeFormula:
		if p.Formula != nil {
			return p.Formula.Validate()
		}
	case PropertyTypeRollup:
		if p.Rollup != nil {
			return p.Rollup.Validate()
		}
	case PropertyTypeVerification:
		if p.Verification != nil {
			return p.Verification.Validate()
		}
	}

	return nil
}

// Validate ensures the FormulaProperty has a valid result type and corresponding value.
//
// Returns:
// - error: Validation error if the formula result is invalid, nil if valid.
func (fp *FormulaProperty) Validate() error {
	switch fp.Type {
	case FormulaResultTypeString:
		if fp.String == nil {
			return fmt.Errorf("string value is required for formula result type 'string'")
		}
	case FormulaResultTypeNumber:
		if fp.Number == nil {
			return fmt.Errorf("number value is required for formula result type 'number'")
		}
	case FormulaResultTypeBoolean:
		if fp.Boolean == nil {
			return fmt.Errorf("boolean value is required for formula result type 'boolean'")
		}
	case FormulaResultTypeDate:
		if fp.Date == nil {
			return fmt.Errorf("date value is required for formula result type 'date'")
		}
		if fp.Date.Start == "" {
			return fmt.Errorf("date start cannot be empty")
		}
	default:
		return fmt.Errorf("unknown formula result type: %s", fp.Type)
	}
	return nil
}

// Validate ensures the RollupProperty has a valid type and corresponding value.
//
// Returns:
// - error: Validation error if the rollup is invalid, nil if valid.
func (rp *RollupProperty) Validate() error {
	switch rp.Type {
	case RollupTypeNumber:
		if rp.Number == nil {
			return fmt.Errorf("number value is required for rollup type 'number'")
		}
	case RollupTypeDate:
		if rp.Date == nil {
			return fmt.Errorf("date value is required for rollup type 'date'")
		}
		if rp.Date.Start == "" {
			return fmt.Errorf("date start cannot be empty")
		}
	case RollupTypeArray:
		if rp.Array == nil {
			return fmt.Errorf("array value is required for rollup type 'array'")
		}
	case RollupTypeIncomplete:
		if rp.Incomplete == nil || !*rp.Incomplete {
			return fmt.Errorf("incomplete must be true for rollup type 'incomplete'")
		}
	case RollupTypeUnsupported:
		// No additional validation required for unsupported type
	default:
		return fmt.Errorf("unknown rollup type: %s", rp.Type)
	}
	return nil
}

// Validate ensures the VerificationProperty has valid configuration.
//
// Returns:
// - error: Validation error if the verification state is invalid, nil if valid.
func (vp *VerificationProperty) Validate() error {
	switch vp.State {
	case VerificationStateVerified:
		if vp.Date == nil {
			return fmt.Errorf("date is required for verification state 'verified'")
		}
		if vp.VerifiedBy == nil {
			return fmt.Errorf("verified_by is required for verification state 'verified'")
		}
	case VerificationStateExpired:
		if vp.Date == nil {
			return fmt.Errorf("date is required for verification state 'expired'")
		}
	case VerificationStateUnverified:
		// No additional validation required for unverified state
	default:
		return fmt.Errorf("unknown verification state: %s", vp.State)
	}
	return nil
}
