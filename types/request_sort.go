package types

// PropertySort represents a sort order for the following endpoints:
// - /databases/{database_id}/query
type PropertySort struct {
	// Property name to sort by
	Property *string `json:"property,omitempty"`
	// Direction to sort in (ascending or descending)
	Direction string `json:"direction"`
}

// TimestampSort represents a sort order for the following endpoints:
// - /search
// - /databases/{database_id}/query
type TimestampSort struct {
	// Timestamp to sort by (created_time or last_edited_time)
	Timestamp *string `json:"timestamp,omitempty"`
	// Direction to sort in (ascending or descending)
	Direction string `json:"direction"`
}
