package types

import (
	"encoding/json"
	"fmt"

	"github.com/cmskitdev/notion/id"
)

// DatabaseID represents a unique identifier for a Notion database.
type DatabaseID id.ID

// String returns the string representation of the DatabaseID.
//
// Returns:
// - string: The string representation of the DatabaseID.
func (p *DatabaseID) String() string {
	return string(*p)
}

// ParseDatabaseID parses a string as a Notion DatabaseID, accepting both dashed and undashed formats.
// It returns the DatabaseID in canonical dashed UUID format (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).
//
// Arguments:
// - id: The string to parse as a DatabaseID.
//
// Returns:
// - DatabaseID: The canonical, dashed DatabaseID.
// - error: If the input is not a valid 32-character hexadecimal string.
func ParseDatabaseID(id string) (DatabaseID, error) {
	databaseID, err := IDParser.Parse(id)
	if err != nil {
		return "", err
	}
	return DatabaseID(databaseID), nil
}

// UnmarshalJSON implements custom JSON unmarshaling for DatabaseID.
// It accepts both dashed and undashed UUID formats and converts them to canonical format.
//
// Arguments:
// - data: Raw JSON bytes containing the ID string.
//
// Returns:
// - error: Parsing error if the ID format is invalid, nil if successful.
func (p *DatabaseID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := ParseDatabaseID(s)
	if err != nil {
		return fmt.Errorf("invalid DatabaseID in JSON (%s): %w", data, err)
	}

	*p = parsed
	return nil
}

// MarshalJSON implements custom JSON marshaling for DatabaseID.
//
// Returns:
// - []byte: JSON-encoded DatabaseID as a string.
// - error: Marshaling error, if any.
func (p DatabaseID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(p))
}

// Validate ensures the DatabaseID is in the correct UUID format.
//
// Returns:
// - error: Validation error if the ID format is invalid, nil if valid.
//
// Example:
//
//	if err := databaseID.Validate(); err != nil {
//	    log.Printf("Invalid DatabaseID: %v", err)
//	}
func (p DatabaseID) Validate() error {
	_, err := ParseDatabaseID(string(p))
	return err
}

// Database represents a Notion database, which is a structured collection of pages with defined properties.
// Databases serve as the schema and container for related pages/records.
type Database struct {
	BaseObject
	PropertyContainer[DatabaseProperty]
	ID          DatabaseID `json:"id"`
	Title       []RichText `json:"title"`
	Description []RichText `json:"description"`
	Icon        *Icon      `json:"icon,omitempty"`
	Cover       *Cover     `json:"cover,omitempty"`
	Parent      *Parent    `json:"parent,omitempty"`
	URL         string     `json:"url"`
	PublicURL   *string    `json:"public_url,omitempty"`
	Archived    bool       `json:"archived,omitempty"`
	InTrash     bool       `json:"in_trash,omitempty"`
	IsInline    bool       `json:"is_inline,omitempty"`
}

// DatabaseProperty represents a property definition in a database schema.
// This defines the structure and constraints for data that can be stored in pages within the database.
type DatabaseProperty struct {
	ID             PropertyID            `json:"id"`
	Name           string                `json:"name"`
	Type           PropertyType          `json:"type"`
	Description    string                `json:"description,omitempty"`
	Title          *TitleConfig          `json:"title,omitempty"`
	RichText       *RichTextConfig       `json:"rich_text,omitempty"`
	Number         *NumberConfig         `json:"number,omitempty"`
	Select         *SelectConfig         `json:"select,omitempty"`
	MultiSelect    *MultiSelectConfig    `json:"multi_select,omitempty"`
	Date           *DateConfig           `json:"date,omitempty"`
	People         *PeopleConfig         `json:"people,omitempty"`
	Files          *FilesConfig          `json:"files,omitempty"`
	Checkbox       *CheckboxConfig       `json:"checkbox,omitempty"`
	URL            *URLConfig            `json:"url,omitempty"`
	Email          *EmailConfig          `json:"email,omitempty"`
	PhoneNumber    *PhoneNumberConfig    `json:"phone_number,omitempty"`
	Formula        *FormulaConfig        `json:"formula,omitempty"`
	Relation       *RelationConfig       `json:"relation,omitempty"`
	Rollup         *RollupConfig         `json:"rollup,omitempty"`
	CreatedTime    *CreatedTimeConfig    `json:"created_time,omitempty"`
	CreatedBy      *CreatedByConfig      `json:"created_by,omitempty"`
	LastEditedTime *LastEditedTimeConfig `json:"last_edited_time,omitempty"`
	LastEditedBy   *LastEditedByConfig   `json:"last_edited_by,omitempty"`
	Status         *StatusConfig         `json:"status,omitempty"`
	UniqueID       *UniqueIDConfig       `json:"unique_id,omitempty"`
	Verification   *VerificationConfig   `json:"verification,omitempty"`
}

// Property configuration types

// TitleConfig represents configuration for a title property.
type TitleConfig struct {
	// Title properties have no additional configuration
}

// UnmarshalJSON implements custom JSON unmarshaling for TitleConfig.
// It handles both empty objects {} and empty arrays [].
func (tc *TitleConfig) UnmarshalJSON(data []byte) error {
	// Handle empty array case
	if string(data) == "[]" {
		return nil
	}
	// Handle object case - since TitleConfig is empty, just ignore the content
	return nil
}

// RichTextConfig represents configuration for a rich text property.
type RichTextConfig struct {
	// Rich text properties have no additional configuration
}

// UnmarshalJSON implements custom JSON unmarshaling for RichTextConfig.
// It handles both empty objects {} and empty arrays [].
func (rtc *RichTextConfig) UnmarshalJSON(data []byte) error {
	// Handle empty array case
	if string(data) == "[]" {
		return nil
	}
	// Handle object case - since RichTextConfig is empty, just ignore the content
	return nil
}

// SelectConfig represents configuration for a select property.
type SelectConfig struct {
	Options []SelectOption `json:"options"`
}

// MultiSelectConfig represents configuration for a multi-select property.
type MultiSelectConfig struct {
	Options []SelectOption `json:"options"`
}

// UnmarshalJSON implements custom JSON unmarshaling for MultiSelectConfig.
// It handles both the standard format {options: [...]} and the direct array format [...].
func (msc *MultiSelectConfig) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as direct array first
	var options []SelectOption
	if err := json.Unmarshal(data, &options); err == nil {
		msc.Options = options
		return nil
	}

	// Fall back to standard format with "options" field
	type Alias MultiSelectConfig
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(msc),
	}
	return json.Unmarshal(data, &aux)
}

// DateConfig represents configuration for a date property.
type DateConfig struct {
	// Date properties have no additional configuration
}

// UnmarshalJSON implements custom JSON unmarshaling for DateConfig.
// It handles both empty objects {} and empty arrays [].
func (dc *DateConfig) UnmarshalJSON(data []byte) error {
	// Handle empty array case
	if string(data) == "[]" {
		return nil
	}
	// Handle object case - since DateConfig is empty, just ignore the content
	return nil
}

// PeopleConfig represents configuration for a people property.
type PeopleConfig struct {
	// People properties have no additional configuration
}

// UnmarshalJSON implements custom JSON unmarshaling for PeopleConfig.
// It handles both empty objects {} and empty arrays [].
func (pc *PeopleConfig) UnmarshalJSON(data []byte) error {
	// Handle empty array case
	if string(data) == "[]" {
		return nil
	}
	// Handle object case - since PeopleConfig is empty, just ignore the content
	return nil
}

// FilesConfig represents configuration for a files property.
type FilesConfig struct {
	// Files properties have no additional configuration
}

// UnmarshalJSON implements custom JSON unmarshaling for FilesConfig.
// It handles both empty objects {} and empty arrays [].
func (fc *FilesConfig) UnmarshalJSON(data []byte) error {
	// Handle empty array case
	if string(data) == "[]" {
		return nil
	}
	// Handle object case - since FilesConfig is empty, just ignore the content
	return nil
}

// CheckboxConfig represents configuration for a checkbox property.
type CheckboxConfig struct {
	// Checkbox properties have no additional configuration
}

// UnmarshalJSON implements custom JSON unmarshaling for CheckboxConfig.
// It handles both empty objects {} and empty arrays [].
func (cc *CheckboxConfig) UnmarshalJSON(data []byte) error {
	// Handle empty array case
	if string(data) == "[]" {
		return nil
	}
	// Handle object case - since CheckboxConfig is empty, just ignore the content
	return nil
}

// URLConfig represents configuration for a URL property.
type URLConfig struct {
	// URL properties have no additional configuration
}

// UnmarshalJSON implements custom JSON unmarshaling for URLConfig.
// It handles both empty objects {} and empty arrays [].
func (uc *URLConfig) UnmarshalJSON(data []byte) error {
	// Handle empty array case
	if string(data) == "[]" {
		return nil
	}
	// Handle object case - since URLConfig is empty, just ignore the content
	return nil
}

// EmailConfig represents configuration for an email property.
type EmailConfig struct {
	// Email properties have no additional configuration
}

// UnmarshalJSON implements custom JSON unmarshaling for EmailConfig.
// It handles both empty objects {} and empty arrays [].
func (ec *EmailConfig) UnmarshalJSON(data []byte) error {
	// Handle empty array case
	if string(data) == "[]" {
		return nil
	}
	// Handle object case - since EmailConfig is empty, just ignore the content
	return nil
}

// PhoneNumberConfig represents configuration for a phone number property.
type PhoneNumberConfig struct {
	// Phone number properties have no additional configuration
}

// UnmarshalJSON implements custom JSON unmarshaling for PhoneNumberConfig.
// It handles both empty objects {} and empty arrays [].
func (pnc *PhoneNumberConfig) UnmarshalJSON(data []byte) error {
	// Handle empty array case
	if string(data) == "[]" {
		return nil
	}
	// Handle object case - since PhoneNumberConfig is empty, just ignore the content
	return nil
}

// FormulaConfig represents configuration for a formula property.
type FormulaConfig struct {
	Expression string `json:"expression"`
}

// CreatedTimeConfig represents configuration for a created time property.
type CreatedTimeConfig struct {
	// Created time properties have no additional configuration
}

// CreatedByConfig represents configuration for a created by property.
type CreatedByConfig struct {
	// Created by properties have no additional configuration
}

// LastEditedTimeConfig represents configuration for a last edited time property.
type LastEditedTimeConfig struct {
	// Last edited time properties have no additional configuration
}

// LastEditedByConfig represents configuration for a last edited by property.
type LastEditedByConfig struct {
	// Last edited by properties have no additional configuration
}

// StatusConfig represents configuration for a status property.
type StatusConfig struct {
	Options []StatusOption `json:"options"`
	Groups  []StatusGroup  `json:"groups"`
}

// StatusOption represents an option in a status property.
type StatusOption struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Color       Color  `json:"color"`
	Description string `json:"description,omitempty"`
}

// StatusGroup represents a group of status options.
type StatusGroup struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Color     Color    `json:"color"`
	OptionIDs []string `json:"option_ids"`
}

// UniqueIDConfig represents configuration for a unique ID property.
type UniqueIDConfig struct {
	Prefix string `json:"prefix,omitempty"`
}

// VerificationConfig represents configuration for a verification property.
type VerificationConfig struct {
	// Verification properties have no additional configuration
}

// NewDatabase creates a new database with the specified title and properties.
//
// Arguments:
// - title: The database title as rich text elements.
// - properties: A map of property names to property configurations.
//
// Returns:
// - *Database: A new database with the specified title and properties.
//
// Example:
//
//	title := []RichText{*NewTextRichText("My Tasks", nil)}
//	properties := map[string]DatabaseProperty{
//	    "Name": {Type: PropertyTypeTitle, Title: &TitleConfig{}},
//	    "Status": {Type: PropertyTypeSelect, Select: &SelectConfig{Options: []SelectOption{
//	        {Name: "To Do", Color: ColorRed},
//	        {Name: "Done", Color: ColorGreen},
//	    }}},
//	}
//	db := NewDatabase(title, properties)
func NewDatabase(title []RichText, properties map[string]DatabaseProperty) *Database {
	return &Database{
		BaseObject: BaseObject{
			Object: ObjectTypeDatabase,
		},
		PropertyContainer: PropertyContainer[DatabaseProperty]{
			Properties: properties,
		},
		Title: title,
	}
}

// GetTitle returns the database title as plain text.
//
// Returns:
// - string: The database title as plain text.
//
// Example:
//
//	title := database.GetTitle()
//	fmt.Printf("Database: %s\n", title)
func (d *Database) GetTitle() string {
	return ToPlainText(d.Title)
}

// Note: GetProperty, SetProperty, and RemoveProperty methods are now inherited from PropertyContainer[DatabaseProperty]

// HasTitleProperty returns true if the database has a title property.
//
// Returns:
// - bool: True if a title property exists, false otherwise.
//
// Example:
//
//	if !database.HasTitleProperty() {
//	    fmt.Println("Warning: Database missing title property")
//	}
func (d *Database) HasTitleProperty() bool {
	for _, prop := range d.Properties {
		if prop.Type == PropertyTypeTitle {
			return true
		}
	}
	return false
}

// GetTitlePropertyName returns the name of the title property.
//
// Returns:
// - string: The name of the title property, or empty string if none exists.
//
// Example:
//
//	titleProp := database.GetTitlePropertyName()
//	fmt.Printf("Title property: %s\n", titleProp)
func (d *Database) GetTitlePropertyName() string {
	for name, prop := range d.Properties {
		if prop.Type == PropertyTypeTitle {
			return name
		}
	}
	return ""
}

// Archive marks the database as archived.
//
// Example:
//
//	database.Archive()
//	fmt.Println("Database archived:", database.Archived)
func (d *Database) Archive() {
	d.Archived = true
}

// Unarchive marks the database as not archived.
//
// Example:
//
//	database.Unarchive()
//	fmt.Println("Database unarchived:", database.Archived == false)
func (d *Database) Unarchive() {
	d.Archived = false
}

// Validate ensures the Database has valid required fields and structure.
//
// Returns:
// - error: Validation error if the database structure is invalid, nil if valid.
//
// Example:
//
//	if err := database.Validate(); err != nil {
//	    log.Printf("Invalid database: %v", err)
//	}
func (d *Database) Validate() error {
	if d.Object != ObjectTypeDatabase {
		return fmt.Errorf("object type must be 'database', got: %s", d.Object)
	}

	// Validate title
	if len(d.Title) == 0 {
		return fmt.Errorf("database must have a title")
	}
	for i, rt := range d.Title {
		if err := rt.Validate(); err != nil {
			return fmt.Errorf("invalid title rich text at index %d: %w", i, err)
		}
	}

	// Validate description if present
	for i, rt := range d.Description {
		if err := rt.Validate(); err != nil {
			return fmt.Errorf("invalid description rich text at index %d: %w", i, err)
		}
	}

	// Validate properties using inherited method
	if err := d.ValidateProperties(); err != nil {
		return err
	}

	// Validate all properties and check for title property
	hasTitleProperty := false
	for _, prop := range d.Properties {
		if prop.Type == PropertyTypeTitle {
			hasTitleProperty = true
			break
		}
	}

	// Ensure database has exactly one title property
	if !hasTitleProperty {
		return fmt.Errorf("database must have exactly one title property")
	}

	// Validate parent if present
	if d.Parent != nil {
		if err := d.Parent.Validate(); err != nil {
			return fmt.Errorf("invalid parent: %w", err)
		}
	}

	// Validate icon if present
	if d.Icon != nil {
		if err := d.Icon.Validate(); err != nil {
			return fmt.Errorf("invalid icon: %w", err)
		}
	}

	// Validate cover if present
	if d.Cover != nil {
		if err := d.Cover.Validate(); err != nil {
			return fmt.Errorf("invalid cover: %w", err)
		}
	}

	return nil
}

// Validate ensures the DatabaseProperty has valid configuration based on its type.
//
// Returns:
// - error: Validation error if the property configuration is invalid, nil if valid.
func (dp *DatabaseProperty) Validate() error {
	if dp.Name == "" {
		return fmt.Errorf("property name cannot be empty")
	}

	switch dp.Type {
	case PropertyTypeTitle:
		if dp.Title == nil {
			return fmt.Errorf("title config is required for property type 'title'")
		}
	case PropertyTypeNumber:
		if dp.Number == nil {
			return fmt.Errorf("number config is required for property type 'number'")
		}
	case PropertyTypeSelect:
		if dp.Select == nil {
			return fmt.Errorf("select config is required for property type 'select'")
		}
		return dp.Select.Validate()
	case PropertyTypeMultiSelect:
		if dp.MultiSelect == nil {
			return fmt.Errorf("multi_select config is required for property type 'multi_select'")
		}
		return dp.MultiSelect.Validate()
	case PropertyTypeFormula:
		if dp.Formula == nil {
			return fmt.Errorf("formula config is required for property type 'formula'")
		}
		return dp.Formula.Validate()
	case PropertyTypeRelation:
		if dp.Relation == nil {
			return fmt.Errorf("relation config is required for property type 'relation'")
		}
		return dp.Relation.Validate()
	case PropertyTypeRollup:
		if dp.Rollup == nil {
			return fmt.Errorf("rollup config is required for property type 'rollup'")
		}
		return dp.Rollup.Validate()
	case PropertyTypeStatus:
		if dp.Status == nil {
			return fmt.Errorf("status config is required for property type 'status'")
		}
		return dp.Status.Validate()
	case PropertyTypeUniqueID:
		if dp.UniqueID == nil {
			return fmt.Errorf("unique_id config is required for property type 'unique_id'")
		}
	}

	return nil
}

// Validation methods for configuration types

// Validate ensures the SelectConfig has valid options.
func (sc *SelectConfig) Validate() error {
	if len(sc.Options) == 0 {
		return fmt.Errorf("select config must have at least one option")
	}

	names := make(map[string]bool)
	for i, option := range sc.Options {
		if option.Name == "" {
			return fmt.Errorf("select option at index %d cannot have empty name", i)
		}
		if names[option.Name] {
			return fmt.Errorf("duplicate select option name: %s", option.Name)
		}
		names[option.Name] = true
	}
	return nil
}

// Validate ensures the MultiSelectConfig has valid options.
func (msc *MultiSelectConfig) Validate() error {
	if len(msc.Options) == 0 {
		return fmt.Errorf("multi-select config must have at least one option")
	}

	names := make(map[string]bool)
	for i, option := range msc.Options {
		if option.Name == "" {
			return fmt.Errorf("multi-select option at index %d cannot have empty name", i)
		}
		if names[option.Name] {
			return fmt.Errorf("duplicate multi-select option name: %s", option.Name)
		}
		names[option.Name] = true
	}
	return nil
}

// Validate ensures the FormulaConfig has a valid expression.
func (fc *FormulaConfig) Validate() error {
	if fc.Expression == "" {
		return fmt.Errorf("formula expression cannot be empty")
	}
	return nil
}

// Validate ensures the RelationConfig has valid configuration.
func (rc *RelationConfig) Validate() error {
	switch rc.Type {
	case RelationTypeDualProperty:
		if rc.DualProperty == nil {
			return fmt.Errorf("dual_property config is required for relation type 'dual_property'")
		}
		if rc.DualProperty.SyncedPropertyName == "" {
			return fmt.Errorf("synced property name cannot be empty")
		}
	}

	return nil
}

// Validate ensures the RollupConfig has valid configuration.
func (rc *RollupConfig) Validate() error {
	if rc.RelationPropertyName == "" {
		return fmt.Errorf("relation property name cannot be empty")
	}
	if rc.RollupPropertyName == "" {
		return fmt.Errorf("rollup property name cannot be empty")
	}

	return nil
}

// Validate ensures the StatusConfig has valid options and groups.
func (sc *StatusConfig) Validate() error {
	if len(sc.Options) == 0 {
		return fmt.Errorf("status config must have at least one option")
	}

	optionNames := make(map[string]bool)
	optionIDs := make(map[string]bool)

	for i, option := range sc.Options {
		if option.ID == "" {
			return fmt.Errorf("status option at index %d cannot have empty ID", i)
		}
		if option.Name == "" {
			return fmt.Errorf("status option at index %d cannot have empty name", i)
		}
		if optionIDs[option.ID] {
			return fmt.Errorf("duplicate status option ID: %s", option.ID)
		}
		if optionNames[option.Name] {
			return fmt.Errorf("duplicate status option name: %s", option.Name)
		}
		optionIDs[option.ID] = true
		optionNames[option.Name] = true
	}

	// Validate groups if present
	groupNames := make(map[string]bool)
	for i, group := range sc.Groups {
		if group.Name == "" {
			return fmt.Errorf("status group at index %d cannot have empty name", i)
		}
		if groupNames[group.Name] {
			return fmt.Errorf("duplicate status group name: %s", group.Name)
		}
		groupNames[group.Name] = true

		// Validate that all option IDs in the group exist
		for _, optionID := range group.OptionIDs {
			if !optionIDs[optionID] {
				return fmt.Errorf("status group '%s' references non-existent option ID: %s", group.Name, optionID)
			}
		}
	}

	return nil
}
