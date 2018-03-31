package schema

import "fmt"

// Row row in table.
type Row struct {
	table  *Table              `json:"table"`
	ID     string              `json:"id"`
	Fields map[string]RowField `json:"fields"`
}

// Key returns the key identifying the row.
func (r *Row) Key() string {
	return CreateRowKey(r.table, r.ID)
}

// CreateRowKey creates a key identifying a table row.
func CreateRowKey(table *Table, ID string) string {
	return fmt.Sprintf("%s%s%s", table.Key(), KeyDelimiter, ID)
}

// RowField field in a row with datatype and data,
type RowField struct {
	Type ColumnType `json:"type"`
	Data []byte     `json:"data"`
}

// Equals checks if a candidate is equeal to the compared field.
func (rf *RowField) Equals(candidate *RowField) bool {
	if !rf.Type.Equals(candidate.Type) {
		return false
	}
	if len(rf.Data) != len(candidate.Data) {
		return false
	}
	for i, b := range rf.Data {
		if b != candidate.Data[i] {
			return false
		}
	}
	return true
}
