package types

import (
	"encoding/json"
	"testing"
)

// TestNumberProperty_UnmarshalJSON tests JSON unmarshaling for NumberProperty.
// This test ensures that NumberProperty can handle both direct number values and structured format.
func TestNumberProperty_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *float64
		wantErr  bool
	}{
		{
			name:     "Direct number value",
			input:    `42.5`,
			expected: func() *float64 { f := 42.5; return &f }(),
			wantErr:  false,
		},
		{
			name:     "Direct integer value",
			input:    `100`,
			expected: func() *float64 { f := 100.0; return &f }(),
			wantErr:  false,
		},
		{
			name:     "Null value",
			input:    `null`,
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Structured format",
			input:    `{"number": 75.25}`,
			expected: func() *float64 { f := 75.25; return &f }(),
			wantErr:  false,
		},
		{
			name:     "Structured format with null",
			input:    `{"number": null}`,
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "Invalid JSON format",
			input:    `{"invalid": "format"}`,
			expected: nil,
			wantErr:  false, // This will set Number to nil since it can't parse the structure
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var np NumberProperty
			err := json.Unmarshal([]byte(tt.input), &np)

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return // Skip value checking for error cases
			}

			if tt.expected == nil {
				if np.Number != nil {
					t.Errorf("Expected nil, got %v", *np.Number)
				}
			} else {
				if np.Number == nil {
					t.Errorf("Expected %v, got nil", *tt.expected)
				} else if *np.Number != *tt.expected {
					t.Errorf("Expected %v, got %v", *tt.expected, *np.Number)
				}
			}
		})
	}
}

// TestNumberProperty_MarshalJSON tests JSON marshaling for NumberProperty.
// This test ensures that NumberProperty marshals correctly to JSON.
func TestNumberProperty_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		property NumberProperty
		expected string
	}{
		{
			name:     "With number value",
			property: NumberProperty{Number: func() *float64 { f := 42.5; return &f }()},
			expected: `{"number":42.5}`,
		},
		{
			name:     "With nil value",
			property: NumberProperty{Number: nil},
			expected: `{"number":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshaled, err := json.Marshal(tt.property)
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
				return
			}

			if string(marshaled) != tt.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(marshaled), tt.expected)
			}
		})
	}
}

// TestProperty_NumberPropertyIntegration tests NumberProperty within Property struct.
// This test ensures that the Property struct correctly handles NumberProperty values.
func TestProperty_NumberPropertyIntegration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *float64
		wantErr  bool
	}{
		{
			name:     "Property with direct number",
			input:    `{"id":"test","type":"number","number":123.45}`,
			expected: func() *float64 { f := 123.45; return &f }(),
			wantErr:  false,
		},
		{
			name:     "Property with structured number",
			input:    `{"id":"test","type":"number","number":{"number":67.89}}`,
			expected: func() *float64 { f := 67.89; return &f }(),
			wantErr:  false,
		},
		{
			name:     "Property with null number",
			input:    `{"id":"test","type":"number","number":null}`,
			expected: nil,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var prop Property
			err := json.Unmarshal([]byte(tt.input), &prop)

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			// For null number case, NumberProperty might not be created at all
			if tt.expected == nil && prop.Number == nil {
				return // This is expected for null values
			}

			if prop.Number == nil {
				t.Errorf("Expected NumberProperty to be set, got nil")
				return
			}

			if tt.expected == nil {
				if prop.Number.Number != nil {
					t.Errorf("Expected nil number, got %v", *prop.Number.Number)
				}
			} else {
				if prop.Number.Number == nil {
					t.Errorf("Expected %v, got nil", *tt.expected)
				} else if *prop.Number.Number != *tt.expected {
					t.Errorf("Expected %v, got %v", *tt.expected, *prop.Number.Number)
				}
			}
		})
	}
}
