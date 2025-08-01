// package types contains the types needed for talking to the Notion API.
// This package is imported by modules like client, plugins, etc as part of
// additional abstraction layer(s) to increase functionality and reduce coupling.
//
// See: https://github.com/notioncodes/types
package types

// DatabaseComposite is a composite type that contains a database and a list of pages.
// It is used to represent a database and its pages in a single type.
type DatabaseComposite struct {
	Database *Database `json:"database"`
	Pages    []*Page   `json:"pages"`
}
