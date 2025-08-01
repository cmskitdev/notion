package types

// CheckboxFilter represents a checkbox property filter.
type CheckboxFilter struct {
	Equals       *bool `json:"equals,omitempty"`
	DoesNotEqual *bool `json:"does_not_equal,omitempty"`
}

// DateFilter represents a date property filter.
type DateFilter struct {
	Equals     *string `json:"equals,omitempty"`
	Before     *string `json:"before,omitempty"`
	After      *string `json:"after,omitempty"`
	OnOrBefore *string `json:"on_or_before,omitempty"`
	OnOrAfter  *string `json:"on_or_after,omitempty"`
	IsEmpty    *bool   `json:"is_empty,omitempty"`
	IsNotEmpty *bool   `json:"is_not_empty,omitempty"`
}

// FilesFilter represents a files property filter.
type FilesFilter struct {
	IsEmpty    *bool `json:"is_empty,omitempty"`
	IsNotEmpty *bool `json:"is_not_empty,omitempty"`
}

// NumberFilter represents a number property filter.
type NumberFilter struct {
	Equals               *float64 `json:"equals,omitempty"`
	DoesNotEqual         *float64 `json:"does_not_equal,omitempty"`
	GreaterThan          *float64 `json:"greater_than,omitempty"`
	LessThan             *float64 `json:"less_than,omitempty"`
	GreaterThanOrEqualTo *float64 `json:"greater_than_or_equal_to,omitempty"`
	LessThanOrEqualTo    *float64 `json:"less_than_or_equal_to,omitempty"`
	IsEmpty              *bool    `json:"is_empty,omitempty"`
	IsNotEmpty           *bool    `json:"is_not_empty,omitempty"`
}

// RichTextFilter represents a rich text property filter.
type RichTextFilter struct {
	Equals         *string `json:"equals,omitempty"`
	DoesNotEqual   *string `json:"does_not_equal,omitempty"`
	Contains       *string `json:"contains,omitempty"`
	DoesNotContain *string `json:"does_not_contain,omitempty"`
	StartsWith     *string `json:"starts_with,omitempty"`
	EndsWith       *string `json:"ends_with,omitempty"`
	IsEmpty        *bool   `json:"is_empty,omitempty"`
	IsNotEmpty     *bool   `json:"is_not_empty,omitempty"`
}

// SelectFilter represents a select property filter.
type SelectFilter struct {
	Equals       *string `json:"equals,omitempty"`
	DoesNotEqual *string `json:"does_not_equal,omitempty"`
	IsEmpty      *bool   `json:"is_empty,omitempty"`
	IsNotEmpty   *bool   `json:"is_not_empty,omitempty"`
}

// StatusFilter represents a status property filter.
type StatusFilter struct {
	Equals       *string `json:"equals,omitempty"`
	DoesNotEqual *string `json:"does_not_equal,omitempty"`
	IsEmpty      *bool   `json:"is_empty,omitempty"`
	IsNotEmpty   *bool   `json:"is_not_empty,omitempty"`
}

// TimestampFilter represents a timestamp property filter.
type TimestampFilter struct {
	CreatedTime    *DateFilter `json:"created_time,omitempty"`
	LastEditedTime *DateFilter `json:"last_edited_time,omitempty"`
}
