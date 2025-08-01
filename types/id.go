// package types provides type-safe ID types for Notion entities.
//
// This package contains ID types for different Notion entity types,
// including Page, Database, Block, User, and Property.
//
// The ID types are designed to be immutable and provide type safety
// for different entity identifiers. They prevent mixing up different
// types of IDs in function calls.
package types

import (
	"github.com/cmskitdev/notion/id"
)

// IDParser is a global parser for Notion IDs.
//
// It is thread safe and can be reused for multiple calls in goroutines.
// It also has a minimal cache to improve performance.
var IDParser = id.NewIDParser(id.NewNoOpCache())

// ExtractID extracts the ID from a Notion object.
//
// Arguments:
// - obj: The Notion object to extract the ID from.
//
// Returns:
// - The ID of the Notion object.
func ExtractID(obj interface{}) string {
	switch v := obj.(type) {
	case *Page:
		return v.ID.String()
	case *Database:
		return v.ID.String()
	case *Block:
		return v.ID.String()
	case *User:
		return v.ID.String()
	default:
		return "unknown"
	}
}
