package schema

import "math"

// Datatype names
const (
	VARCHAR = "VARCHAR"
	NUMBER  = "NUMBER"
	INTEGER = "INTEGER"
	TEXT    = "TEXT"
	JSON    = "JSON"
)

var (
	IntegerPrecision = Precision{math.MaxInt64, math.MinInt64, true}
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

// Equals checks if column types are equal.
func (ct ColumnType) Equals(candidate ColumnType) bool {
	return ct.Name == candidate.Name
}

// Precision precision of column type.
type Precision struct {
	Precision  int64 `json:"precision"`
	Scale      int64 `json:"scale"`
	Applicable bool  `json:"applicable"`
}

// ColumnTypeName name of column datatype.
type ColumnTypeName string
