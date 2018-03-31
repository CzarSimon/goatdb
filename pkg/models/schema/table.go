package schema

import (
	"encoding/json"
	"fmt"
)

const (
	TablesKey    = "tables"
	KeyDelimiter = ":"
)

// Entity storable entity that can be serialized and identified by a key.
type Entity interface {
	Key() string
	Bytes() ([]byte, error)
}

// Table datastructure represnting tables.
type Table struct {
	Name        string            `json:"name"`
	Columns     map[string]Column `json:"columns"`
	Constraints Constraints       `json:"constraints"`
	Sequence    *Sequence         `json:"sequence,omitempty"`
}

// NewTable creates a new table.
func NewTable(name string, columns map[string]Column, constraints Constraints) Table {
	table := Table{
		Name:        name,
		Columns:     columns,
		Constraints: constraints,
	}
	if table.Constraints.PrimaryKey == nil {
		table.Sequence = NewSequence(name)
	}
	return table
}

// NameKey returns the identifying key for the table.
func (t *Table) Key() string {
	return CreateTableKey(t.Name)
}

// Bytes returns a representation of a table info as an array slice.
func (t *Table) Bytes() ([]byte, error) {
	return json.Marshal(*t)
}

func (t *Table) KeyPrefix() string {
	return t.Key() + KeyDelimiter
}

// CreateTableKey creates the identifying key for a table.
func CreateTableKey(name string) string {
	return fmt.Sprintf("table%s%s", KeyDelimiter, name)
}
