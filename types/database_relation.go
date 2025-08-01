package types

// RelationConfig represents configuration for a relation property.
type RelationConfig struct {
	DatabaseID     DatabaseID            `json:"database_id"`
	Type           RelationType          `json:"type,omitempty"`
	SingleProperty *SingleRelationConfig `json:"single_property,omitempty"`
	DualProperty   *DualRelationConfig   `json:"dual_property,omitempty"`
}

// RelationType represents the type of relation.
type RelationType string

// RelationType represents the type of relation.
//
// See: https://developers.notion.com/reference/property-object#relation
const (
	RelationTypeSingleProperty RelationType = "single_property"
	RelationTypeDualProperty   RelationType = "dual_property"
)

// SingleRelationConfig represents configuration for a single-property relation.
type SingleRelationConfig struct {
	// Single property relations have no additional configuration
}

// DualRelationConfig represents configuration for a dual-property relation.
type DualRelationConfig struct {
	DatabaseID         DatabaseID `json:"database_id" validate:"required,uuid"`
	SyncedPropertyName string     `json:"synced_property_name" validate:"required"`
	SyncedPropertyID   PropertyID `json:"synced_property_id" validate:"required"`
}
