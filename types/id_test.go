package types

import (
	"encoding/json"
	"testing"
)

// TestPageID_ParsePageID tests the PageID-specific parsing function.
// This test ensures that PageID parsing works correctly and maintains type safety.
func TestPageID_ParsePageID(t *testing.T) {
	input := "550e8400e29b41d4a716446655440000"
	expected := PageID("550e8400-e29b-41d4-a716-446655440000")

	result, err := ParsePageID(input)
	if err != nil {
		t.Errorf("ParsePageID() error = %v, want nil", err)
		return
	}
	if result != expected {
		t.Errorf("ParsePageID() = %v, want %v", result, expected)
	}
}

// TestPageID_StringMethods tests the String and Compact methods for PageID.
// This test ensures that PageID can be converted to string representations correctly.
func TestPageID_StringMethods(t *testing.T) {
	pageID := PageID("550e8400-e29b-41d4-a716-446655440000")

	// Test String() method
	if pageID.String() != "550e8400-e29b-41d4-a716-446655440000" {
		t.Errorf("String() = %v, want %v", pageID.String(), "550e8400-e29b-41d4-a716-446655440000")
	}
}

// TestPageID_JSONMarshaling tests JSON marshaling and unmarshaling for PageID.
// This test ensures that PageID can be safely serialized to/from JSON in UUID format.
func TestPageID_JSONMarshaling(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected PageID
	}{
		{
			name:     "UUID with dashes",
			input:    `"550e8400-e29b-41d4-a716-446655440000"`,
			expected: PageID("550e8400-e29b-41d4-a716-446655440000"),
		},
		{
			name:     "32 char hex string without dashes",
			input:    `"550e8400e29b41d4a716446655440000"`,
			expected: PageID("550e8400-e29b-41d4-a716-446655440000"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pageID PageID
			err := json.Unmarshal([]byte(tt.input), &pageID)
			if err != nil {
				t.Errorf("json.Unmarshal() error = %v, want nil", err)
				return
			}
			if pageID != tt.expected {
				t.Errorf("json.Unmarshal() = %v, want %v", pageID, tt.expected)
			}

			// Test marshaling back
			marshaled, err := json.Marshal(pageID)
			if err != nil {
				t.Errorf("json.Marshal() error = %v, want nil", err)
				return
			}
			expectedJSON := `"` + string(tt.expected) + `"`
			if string(marshaled) != expectedJSON {
				t.Errorf("json.Marshal() = %v, want %v", string(marshaled), expectedJSON)
			}
		})
	}
}

// TestPageID_JSONUnmarshalingErrors tests JSON unmarshaling error cases for PageID.
// This test ensures that invalid JSON input is properly rejected during unmarshaling.
func TestPageID_JSONUnmarshalingErrors(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Invalid hex characters",
			input: `"550g8400e29b41d4a716446655440000"`,
		},
		{
			name:  "Too short",
			input: `"550e8400e29b41d4a71644665544000"`,
		},
		{
			name:  "Too long",
			input: `"550e8400e29b41d4a7164466554400000"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pageID PageID
			err := json.Unmarshal([]byte(tt.input), &pageID)
			if err == nil {
				t.Errorf("json.Unmarshal() error = nil, want error")
			}
		})
	}
}

// TestPageID_Validation tests the Validate method for PageID.
// This test ensures that PageID validation works correctly for valid and invalid IDs.
func TestPageID_Validation(t *testing.T) {
	tests := []struct {
		name      string
		pageID    PageID
		wantError bool
	}{
		{
			name:      "Valid UUID with dashes",
			pageID:    PageID("550e8400-e29b-41d4-a716-446655440000"),
			wantError: false,
		},
		{
			name:      "Invalid - too short",
			pageID:    PageID("550e8400-e29b-41d4-a716-44665544000"),
			wantError: true,
		},
		{
			name:      "Invalid - contains non-hex",
			pageID:    PageID("550g8400-e29b-41d4-a716-446655440000"),
			wantError: true,
		},
		{
			name:      "Invalid - wrong format",
			pageID:    PageID("not-a-uuid"),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.pageID.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("PageID.Validate() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

// TestAllIDTypes_TypeSafety tests that different ID types are properly distinct.
// This test ensures that the type system prevents mixing different ID types.
func TestAllIDTypes_TypeSafety(t *testing.T) {
	// Test that we can create different ID types from the same string
	idString := "550e8400e29b41d4a716446655440000"

	pageID, err := ParsePageID(idString)
	if err != nil {
		t.Errorf("ParsePageID() error = %v, want nil", err)
	}

	databaseID, err := ParseDatabaseID(idString)
	if err != nil {
		t.Errorf("ParseDatabaseID() error = %v, want nil", err)
	}

	blockID, err := ParseBlockID(idString)
	if err != nil {
		t.Errorf("ParseBlockID() error = %v, want nil", err)
	}

	userID, err := ParseUserID(idString)
	if err != nil {
		t.Errorf("ParseUserID() error = %v, want nil", err)
	}

	propertyID, err := ParsePropertyID(idString)
	if err != nil {
		t.Errorf("ParsePropertyID() error = %v, want nil", err)
	}

	// All should have the same string representation (UUID format)
	expectedUUID := "550e8400-e29b-41d4-a716-446655440000"
	if pageID.String() != expectedUUID {
		t.Errorf("PageID.String() = %v, want %v", pageID.String(), expectedUUID)
	}
	if databaseID.String() != expectedUUID {
		t.Errorf("DatabaseID.String() = %v, want %v", databaseID.String(), expectedUUID)
	}
	if blockID.String() != expectedUUID {
		t.Errorf("BlockID.String() = %v, want %v", blockID.String(), expectedUUID)
	}
	if userID.String() != expectedUUID {
		t.Errorf("UserID.String() = %v, want %v", userID.String(), expectedUUID)
	}
	if propertyID.String() != expectedUUID {
		t.Errorf("PropertyID.String() = %v, want %v", propertyID.String(), expectedUUID)
	}
}

// TestAllIDTypes_JSONSerialization tests JSON serialization for all ID types.
// This test ensures that all ID types serialize to UUID format for API calls.
func TestAllIDTypes_JSONSerialization(t *testing.T) {
	idString := "550e8400e29b41d4a716446655440000"
	expectedJSON := `"550e8400-e29b-41d4-a716-446655440000"`

	tests := []struct {
		name string
		id   interface{}
	}{
		{name: "PageID", id: must(ParsePageID(idString))},
		{name: "DatabaseID", id: must(ParseDatabaseID(idString))},
		{name: "BlockID", id: must(ParseBlockID(idString))},
		{name: "UserID", id: must(ParseUserID(idString))},
		{name: "PropertyID", id: must(ParsePropertyID(idString))},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshaled, err := json.Marshal(tt.id)
			if err != nil {
				t.Errorf("json.Marshal() error = %v, want nil", err)
				return
			}
			if string(marshaled) != expectedJSON {
				t.Errorf("json.Marshal() = %v, want %v", string(marshaled), expectedJSON)
			}
		})
	}
}

// Helper function for tests
func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
