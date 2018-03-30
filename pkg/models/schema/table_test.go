package schema

import (
	"testing"
)

func TestNewTable(t *testing.T) {
	tbl1 := getTestTable("test", true)
	if tbl1.Sequence != nil {
		t.Errorf("Expected table to have empty sequence. Found=%+v", tbl1.Sequence)
	}
	tbl2 := getTestTable("test", false)
	if tbl2.Sequence == nil {
		t.Errorf("Expected table to have sequence. Found=nil")
	}
}

func TestTableNameKey(t *testing.T) {
	table := getTestTable("test", true)
	actualKey := table.NameKey()
	if actualKey != "table:test" {
		t.Errorf("Table.NameKey() wrong. Expected=table:test Got=%s", actualKey)
	}
}

func getTestTable(name string, withPK bool) Table {
	var constraints Constraints
	if withPK {
		constraints = getTestConstraints()
	}
	return NewTable(name, getTestColumns(), constraints)
}

func getTestColumns() map[string]Column {
	columns := make(map[string]Column)
	columns["col1"] = Column{
		Name: "col1",
		Type: ColumnType{
			Name:      VARCHAR,
			Precision: Precision{50, 0, true},
		},
	}
	return columns
}

func getTestConstraints() Constraints {
	return Constraints{
		PrimaryKey: &PrimaryKey{
			Columns: []string{"col1"},
		},
	}
}
