package schema

// Datatype names
const (
	VARCHAR = "VARCHAR"
	NUMBER  = "NUMBER"
	INTEGER = "INTEGER"
	TEXT    = "TEXT"
	JSON    = "JSON"
)

// Column represents a column
type Column struct {
	Name string     `json:"name"`
	Type ColumnType `json"type"`
}

// ColumnType type information of a column datatype.
type ColumnType struct {
	Name      ColumnTypeName `json:"name"`
	Precision Precision      `json:"precision"`
}

// Precision precision of column type.
type Precision struct {
	Precision  int64 `json:"precision"`
	Scale      int64 `json:"scale"`
	Applicable bool  `json:"applicable"`
}

// ColumnTypeName name of column datatype.
type ColumnTypeName string
