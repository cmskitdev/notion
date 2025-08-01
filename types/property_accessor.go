package types

// PropertyAccessor provides typed property access methods for Property values
// through the `PropertyContainer` struct.
//
// This is embedded in `Page` and `Database` objects to provide convenient access
// to property values.
type PropertyAccessor[T any] struct {
	*PropertyContainer[Property]
}

// GetPropertyValue returns the raw value of a property, handling type conversion
// automatically.
//
// Arguments:
// - name: The name of the property to retrieve.
//
// Returns:
// - *T: The property value, or nil if not found.
func (pa *PropertyAccessor[T]) GetPropertyValue(name string) *T {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil {
		return nil
	}
	value := prop.GetValue()
	if value == nil {
		return nil
	}
	if typedValue, ok := value.(T); ok {
		return &typedValue
	}
	return nil
}

// GetValueOr returns the property value or a default value if the property is not found.
//
// Arguments:
// - name: The name of the property to retrieve.
// - defaultValue: The default value to return if the property is not found.
//
// Returns:
// - T: The property value, or the default value if not found.
func (pa *PropertyAccessor[T]) GetValueOr(name string, defaultValue T) T {
	value := pa.GetPropertyValue(name)
	if value == nil {
		return defaultValue
	}
	return *value
}

// GetTitleProperty returns the title property value as a RichText slice.
//
// Arguments:
// - name: The name of the title property.
//
// Returns:
// - []RichText: The title as a RichText slice, or nil if not found or not a title property.
func (pa *PropertyAccessor[T]) GetTitleProperty(name string) []RichText {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeTitle {
		return nil
	}
	return prop.Title
}

// GetTitleText returns the title property value as plain text.
//
// Arguments:
// - name: The name of the title property.
//
// Returns:
// - string: The title as plain text, or empty string if not found.
func (pa *PropertyAccessor[T]) GetTitleText(name string) string {
	title := pa.GetTitleProperty(name)
	if title == nil {
		return ""
	}
	return ToPlainText(title)
}

// GetRichTextProperty returns the rich text property value.
//
// Arguments:
// - name: The name of the rich text property.
//
// Returns:
// - []RichText: The rich text content, or nil if not found or not a rich text property.
func (pa *PropertyAccessor[T]) GetRichTextProperty(name string) []RichText {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeRichText {
		return nil
	}
	return prop.RichText
}

// GetRichTextText returns the rich text property value as plain text.
//
// Arguments:
// - name: The name of the rich text property.
//
// Returns:
// - string: The rich text as plain text, or empty string if not found.
func (pa *PropertyAccessor[T]) GetRichTextText(name string) string {
	richText := pa.GetRichTextProperty(name)
	if richText == nil {
		return ""
	}
	return ToPlainText(richText)
}

// GetSelectProperty returns the select property value.
//
// Arguments:
// - name: The name of the select property.
//
// Returns:
// - *SelectProperty: The selected option, or nil if not found or not a select property.
func (pa *PropertyAccessor[T]) GetSelectProperty(name string) *SelectProperty {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeSelect {
		return nil
	}
	return prop.Select
}

// GetSelectName returns the select property option name.
//
// Arguments:
// - name: The name of the select property.
//
// Returns:
// - string: The selected option name, or empty string if not found.
func (pa *PropertyAccessor[T]) GetSelectName(name string) string {
	option := pa.GetSelectProperty(name)
	if option == nil || option.Name == nil {
		return ""
	}
	return *option.Name
}

// GetMultiSelectProperty returns the multi-select property values.
//
// Arguments:
// - name: The name of the multi-select property.
//
// Returns:
// - []SelectOption: The selected options, or nil if not found or not a multi-select property.
func (pa *PropertyAccessor[T]) GetMultiSelectProperty(name string) []SelectOption {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeMultiSelect {
		return nil
	}
	return prop.MultiSelect
}

// GetMultiSelectNames returns the multi-select property option names.
//
// Arguments:
// - name: The name of the multi-select property.
//
// Returns:
// - []string: The selected option names, or empty slice if not found.
func (pa *PropertyAccessor[T]) GetMultiSelectNames(name string) []string {
	options := pa.GetMultiSelectProperty(name)
	if options == nil {
		return []string{}
	}

	names := make([]string, len(options))
	for i, option := range options {
		names[i] = option.Name
	}
	return names
}

// GetNumberProperty returns the number property value.
//
// Arguments:
// - name: The name of the number property.
//
// Returns:
// - *NumberProperty: The number property, or nil if not found or not a number property.
func (pa *PropertyAccessor[T]) GetNumberProperty(name string) *NumberProperty {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeNumber {
		return nil
	}
	return prop.Number
}

// GetNumberValue returns the number property value as a float64.
//
// Arguments:
// - name: The name of the number property.
//
// Returns:
// - *float64: The number value, or nil if not found or not a number property.
func (pa *PropertyAccessor[T]) GetNumberValue(name string) *float64 {
	numProp := pa.GetNumberProperty(name)
	if numProp == nil {
		return nil
	}
	return numProp.Number
}

// GetCheckboxProperty returns the checkbox property value.
//
// Arguments:
// - name: The name of the checkbox property.
//
// Returns:
// - *bool: The checkbox value, or nil if not found or not a checkbox property.
func (pa *PropertyAccessor[T]) GetCheckboxProperty(name string) *bool {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeCheckbox {
		return nil
	}
	return prop.Checkbox
}

// GetURLProperty returns the URL property value.
//
// Arguments:
// - name: The name of the URL property.
//
// Returns:
// - *string: The URL value, or nil if not found or not a URL property.
func (pa *PropertyAccessor[T]) GetURLProperty(name string) *string {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeURL {
		return nil
	}
	return prop.URL
}

// GetEmailProperty returns the email property value.
//
// Arguments:
// - name: The name of the email property.
//
// Returns:
// - *string: The email value, or nil if not found or not an email property.
func (pa *PropertyAccessor[T]) GetEmailProperty(name string) *string {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeEmail {
		return nil
	}
	return prop.Email
}

// GetPhoneNumberProperty returns the phone number property value.
//
// Arguments:
// - name: The name of the phone number property.
//
// Returns:
// - *string: The phone number value, or nil if not found or not a phone number property.
func (pa *PropertyAccessor[T]) GetPhoneNumberProperty(name string) *string {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypePhoneNumber {
		return nil
	}
	return prop.PhoneNumber
}

// GetDateProperty returns the date property value.
//
// Arguments:
// - name: The name of the date property.
//
// Returns:
// - *DateProperty: The date value, or nil if not found or not a date property.
func (pa *PropertyAccessor[T]) GetDateProperty(name string) *DateProperty {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeDate {
		return nil
	}
	return prop.Date
}

// GetPeopleProperty returns the people property value.
//
// Arguments:
// - name: The name of the people property.
//
// Returns:
// - []User: The people list, or nil if not found or not a people property.
func (pa *PropertyAccessor[T]) GetPeopleProperty(name string) []User {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypePeople {
		return nil
	}
	return prop.People
}

// GetFilesProperty returns the files property value.
//
// Arguments:
// - name: The name of the files property.
//
// Returns:
// - []FileProperty: The files list, or nil if not found or not a files property.
func (pa *PropertyAccessor[T]) GetFilesProperty(name string) []FileProperty {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeFiles {
		return nil
	}
	return prop.Files
}

// GetCreatedTimeProperty returns the created time property value.
//
// Arguments:
// - name: The name of the created time property.
//
// Returns:
// - *Timestamp: The created time timestamp, or nil if not found or not a created time property.
func (pa *PropertyAccessor[T]) GetCreatedTimeProperty(name string) *Timestamp {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeCreatedTime {
		return nil
	}
	return prop.CreatedTime
}

// GetCreatedByProperty returns the created by property value.
//
// Arguments:
// - name: The name of the created by property.
//
// Returns:
// - *User: The user who created the page, or nil if not found or not a created by property.
func (pa *PropertyAccessor[T]) GetCreatedByProperty(name string) *User {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeCreatedBy {
		return nil
	}
	return prop.CreatedBy
}

// GetLastEditedTimeProperty returns the last edited time property value.
//
// Arguments:
// - name: The name of the last edited time property.
//
// Returns:
// - *Timestamp: The last edited time timestamp, or nil if not found or not a last edited time property.
func (pa *PropertyAccessor[T]) GetLastEditedTimeProperty(name string) *Timestamp {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeLastEditedTime {
		return nil
	}
	return prop.LastEditedTime
}

// GetLastEditedByProperty returns the last edited by property value.
//
// Arguments:
// - name: The name of the last edited by property.
//
// Returns:
// - *User: The user who last edited the page, or nil if not found or not a last edited by property.
func (pa *PropertyAccessor[T]) GetLastEditedByProperty(name string) *User {
	prop, exists := pa.GetProperty(name)
	if !exists || prop == nil || prop.Type != PropertyTypeLastEditedBy {
		return nil
	}
	return prop.LastEditedBy
}
