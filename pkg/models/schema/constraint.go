package schema

// Constraints constraint rules on data in a table.
type Constraints struct {
	PrimaryKey *PrimaryKey `json:"primaryKey,omitempty"`
}

// PrimaryKey list of columns comprising a primary key.
type PrimaryKey struct {
	Columns []string `json:"columns"`
}
