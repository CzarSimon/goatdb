package schema

import (
	"strings"
	"testing"
)

func TestRowKey(t *testing.T) {
	table := getTestTable("tbl", false)
	row := Row{table: &table, ID: "1"}
	checkRowKey(row, t)
	if row.Key() != "table:tbl:1" {
		t.Errorf("Row.Key() wrong. Expected=table:tbl:1 Got=%s", row.Key())
	}
}

func checkRowKey(row Row, t *testing.T) {
	keyPrefix := row.table.KeyPrefix()
	if !strings.HasPrefix(row.Key(), keyPrefix) {
		t.Fatalf("Row key doesn't start with table key prefix. Expected=%s Got=%s",
			keyPrefix, row.Key())
	}
	if keyPrefix == row.Key() {
		t.Fatalf("Row key should not be equeal to key prefix '%s'", keyPrefix)
	}
}

func TestRowFieldEquals(t *testing.T) {
	row := getTestRowField(VARCHAR, Precision{50, 0, true}, "1000")
	candidate := getTestRowField(VARCHAR, Precision{50, 0, true}, "1000")
	if !row.Equals(candidate) {
		t.Errorf("row.Equals failed. Expected values to be equal:\nE: %+v\nG: %+v",
			row, candidate)
	}
	wrongDataLengthCandidate := getTestRowField(VARCHAR, Precision{50, 0, true}, "100")
	if row.Equals(wrongDataLengthCandidate) {
		t.Errorf("row.Equals expected data length to be different. Expected: %s != %s",
			row.Data, wrongDataLengthCandidate.Data)
	}
	wrongDataCandidate := getTestRowField(VARCHAR, Precision{50, 0, true}, "1001")
	if row.Equals(wrongDataCandidate) {
		t.Errorf("row.Equals expected data to be different. Expected: %s != %s",
			row.Data, wrongDataCandidate.Data)
	}
	wrongTypeCandidate := getTestRowField(INTEGER, IntegerPrecision, "1000")
	if row.Equals(wrongTypeCandidate) {
		t.Errorf("row.Equals failed. Expected type to be different")
	}
}

func getTestRowField(typeName ColumnTypeName, prec Precision, val string) *RowField {
	return &RowField{
		Type: ColumnType{
			Name:      typeName,
			Precision: prec,
		},
		Data: []byte(val),
	}
}
