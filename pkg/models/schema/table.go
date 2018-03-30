package schema

import (
	"fmt"
)

const (
	KeyDelimiter = ":"
)

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
func (t *Table) NameKey() string {
	return fmt.Sprintf("table%s%s", KeyDelimiter, t.Name)
}

type Row struct {
	Table *Table
	Data  map[string]interface{}
}
