package storage

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/CzarSimon/goatdb/pkg/models/schema"
	"github.com/dgraph-io/badger"
)

var (
	ErrTableExists = errors.New("Table exists")
	TablesKey      = []byte(schema.TablesKey)
)

// DB top level storage reference.
type DB struct {
	storage    *badger.DB
	tables     map[string]schema.Table
	tableNames []string
}

// Open opens a storage connection and brings gets all current table metadata.
func Open(name string) (*DB, error) {
	storage, err := openStorageConn(name)
	if err != nil {
		return nil, err
	}
	db := &DB{
		storage: storage,
	}
	err = db.getExistingTables()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Close closes the connection to the underlying storage.
func (db *DB) Close() error {
	return db.storage.Close()
}

// CreateTable stores a new table.
func (db *DB) CreateTable(table schema.Table) error {
	_, tableExits := db.tables[table.Name]
	if tableExits {
		return ErrTableExists
	}
	prevTableNames, prevTables := db.addTable(table)
	err := db.storeTables()
	if err != nil {
		db.updateTablesInfo(prevTableNames, prevTables)
		return err
	}
	return nil
}

// storeTables stores the current table info.
func (db *DB) storeTables() error {
	tx := db.storage.NewTransaction(true)
	err := set(TablesKey, db.tableNames, tx)
	if err != nil {
		return err
	}
	for _, table := range db.tables {
		err := set([]byte(table.Key()), table, tx)
		if err != nil {
			return err
		}
	}
	return tx.Commit(commitCallback)
}

// getExistingTables gets existing tables from database.
func (db *DB) getExistingTables() error {
	tx := db.storage.NewTransaction(false)
	tableNames, err := getTableNames(tx)
	if err != nil {
		db.updateTablesInfo(tableNames, nil)
		tx.Discard()
		if err == badger.ErrKeyNotFound {
			return nil
		}
		return err
	}
	db.tableNames = tableNames
	db.tables, err = getTables(tableNames, tx)
	if err != nil {
		tx.Discard()
		return err
	}
	return tx.Commit(commitCallback)
}

// getTables gets all table definitions matching one of the provided names.
func getTables(names []string, tx *badger.Txn) (map[string]schema.Table, error) {
	tables := make(map[string]schema.Table)
	var table schema.Table
	for _, name := range names {
		err := get([]byte(schema.CreateTableKey(name)), &table, tx)
		if err != nil {
			return nil, err
		}
		tables[name] = table
	}
	return tables, nil
}

// updateTablesInfo changes the table info to new values.
func (db *DB) updateTablesInfo(tableNames []string, tables map[string]schema.Table) {
	db.tableNames = tableNames
	if tables == nil {
		tables = make(map[string]schema.Table)
	}
	db.tables = tables
}

// addTable adds a table to the tracked metadata,
// returns the old tableNames and table map.
func (db *DB) addTable(table schema.Table) ([]string, map[string]schema.Table) {
	prevTableNames := db.tableNames
	prevTables := db.tables
	db.tableNames = append(db.tableNames, table.Name)
	db.tables[table.Name] = table
	return prevTableNames, prevTables
}

// openStorageConn opens a connection to a badger key value store.
func openStorageConn(name string) (*badger.DB, error) {
	opts := badger.DefaultOptions
	opts.Dir = name
	opts.ValueDir = name
	return badger.Open(opts)
}

// getTableNames gets the stored tablenames.
func getTableNames(tx *badger.Txn) ([]string, error) {
	tableNames := make([]string, 0)
	err := get(TablesKey, &tableNames, tx)
	if err != nil {
		tx.Discard()
		return nil, err
	}
	return tableNames, nil
}

// set sets the value of a key in the underlying storeage.
func set(key []byte, data interface{}, tx *badger.Txn) error {
	serilizedData, err := json.Marshal(data)
	if err != nil {
		tx.Discard()
		return err
	}
	err = tx.Set(key, serilizedData)
	if err != nil {
		tx.Discard()
		return err
	}
	return nil
}

// get gets a serialized value from the underlying storage.
func get(key []byte, dest interface{}, tx *badger.Txn) error {
	item, err := tx.Get(key)
	if err != nil {
		return err
	}
	value, err := item.Value()
	if err != nil {
		return err
	}
	return json.Unmarshal(value, dest)
}

// commitCallback default function to be run after a callback.
func commitCallback(err error) {
	if err != nil {
		log.Println(err)
	}
}
