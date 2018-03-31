package storage

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/CzarSimon/goatdb/pkg/models/schema"
)

const (
	TestDBName = "/tmp/goatdb-test"
)

func TestOpenDB(t *testing.T) {
	db := openTestDB(TestDBName, true, t)
	err := db.Close()
	if err != nil {
		t.Errorf("Failed to close database. Error = %s", err)
	}
	removeDB(TestDBName, true, t)
}

func TestCreateTable(t *testing.T) {
	db := openTestDB(TestDBName, true, t)
	table := getTestTable("tbl_1", true)
	err := db.CreateTable(table)
	if err != nil {
		t.Errorf("db.CreateTable failed. Error = %s", err)
	}
	db.Close()
	db = openTestDB(TestDBName, false, t)
	defer db.Close()
	if db.tableNames[0] != "tbl_1" {
		t.Errorf("db.tableNames wrong. Expected[0]=tbl_1 Got[0]=%s", db.tableNames[0])
	}
	actualTable, ok := db.tables[db.tableNames[0]]
	if !ok {
		t.Fatalf("Table 'tbl_1' missing in db.tables")
	}
	testTablesAreEqual(table, actualTable, t)
	err = db.CreateTable(table)
	if err != ErrTableExists {
		t.Errorf("Expected ErrTableExists on duplicate key creation got=%s", err)
	}
}

func testTablesAreEqual(expected, actual schema.Table, t *testing.T) {
	js, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to format expected table: %+v", expected)
	}
	expectedStr := string(js)
	js, err = json.Marshal(expected)
	if err != nil {
		t.Fatalf("Failed to format actual table: %+v", expected)
	}
	actualStr := string(js)
	if expectedStr != actualStr {
		t.Fatalf("Expected and actual table different.\nExpected=%s\nGot=%s",
			expectedStr, actualStr)
	}
}

func openTestDB(name string, new bool, t *testing.T) *DB {
	if new {
		removeDB(name, false, t)
	}
	db, err := Open(name)
	if err != nil {
		t.Fatalf("Failed to open database: Error=%s", err)
	}
	return db
}

func removeDB(name string, failOnError bool, t *testing.T) {
	err := os.RemoveAll(name)
	if failOnError && err != nil {
		t.Fatalf("Failed to remove database: %s", err)
	}
}

func getTestTable(name string, withPK bool) schema.Table {
	var constraints schema.Constraints
	if withPK {
		constraints = getTestConstraints()
	}
	return schema.NewTable(name, getTestColumns(), constraints)
}

func getTestColumns() map[string]schema.Column {
	columns := make(map[string]schema.Column)
	columns["col1"] = schema.Column{
		Name: "col1",
		Type: schema.ColumnType{
			Name:      schema.VARCHAR,
			Precision: schema.Precision{50, 0, true},
		},
	}
	return columns
}

func getTestConstraints() schema.Constraints {
	return schema.Constraints{
		PrimaryKey: &schema.PrimaryKey{
			Columns: []string{"col1"},
		},
	}
}
