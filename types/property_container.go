package types

import "fmt"

// PropertyContainer provides common property management functionality for both
// Database and Page types through generics to maintain type safety while eliminating
// code duplication.
type PropertyContainer[T any] struct {
	Properties map[string]T `json:"properties"`
}

// GetProperty returns a property by name with type safety.
//
// Arguments:
// - name: The name of the property to retrieve.
//
// Returns:
// - *T: The property if found, nil otherwise.
// - bool: True if the property exists, false otherwise.
func (pc *PropertyContainer[T]) GetProperty(name string) (*T, bool) {
	prop, exists := pc.Properties[name]
	if !exists {
		return nil, false
	}
	return &prop, true
}

// SetProperty sets a property value with type safety.
//
// Arguments:
// - name: The name of the property to set.
// - property: The property value to set.
func (pc *PropertyContainer[T]) SetProperty(name string, property T) {
	if pc.Properties == nil {
		pc.Properties = make(map[string]T)
	}
	pc.Properties[name] = property
}

// RemoveProperty removes a property by name.
//
// Arguments:
// - name: The name of the property to remove.
//
// Returns:
// - bool: True if the property was removed, false if it didn't exist.
func (pc *PropertyContainer[T]) RemoveProperty(name string) bool {
	if pc.Properties == nil {
		return false
	}

	if _, exists := pc.Properties[name]; exists {
		delete(pc.Properties, name)
		return true
	}

	return false
}

// HasProperty returns true if a property with the given name exists.
//
// Arguments:
// - name: The name of the property to check.
//
// Returns:
// - bool: True if the property exists, false otherwise.
func (pc *PropertyContainer[T]) HasProperty(name string) bool {
	_, exists := pc.Properties[name]
	return exists
}

// GetPropertyNames returns a slice of all property names.
//
// Returns:
// - []string: A slice containing all property names.
func (pc *PropertyContainer[T]) GetPropertyNames() []string {
	names := make([]string, 0, len(pc.Properties))
	for name := range pc.Properties {
		names = append(names, name)
	}
	return names
}

// PropertyCount returns the number of properties.
//
// Returns:
// - int: The number of properties.
func (pc *PropertyContainer[T]) PropertyCount() int {
	return len(pc.Properties)
}

// ValidateProperties validates all properties in the container.
//
// Returns:
// - error: Validation error if any property is invalid, nil if all are valid.
func (pc *PropertyContainer[T]) ValidateProperties() error {
	if pc.Properties == nil {
		return fmt.Errorf("properties map cannot be nil")
	}

	// Validate all properties using interface{} assertion to call Validate if available
	for name, prop := range pc.Properties {
		if name == "" {
			return fmt.Errorf("property name cannot be empty")
		}

		// Check if the property has a Validate method
		if validator, ok := any(prop).(interface{ Validate() error }); ok {
			if err := validator.Validate(); err != nil {
				return fmt.Errorf("invalid property '%s': %w", name, err)
			}
		}
	}

	return nil
}
