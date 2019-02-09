package storage

import "abcdb/sql"

// Storage is a storage manager providing high level management to db internal data.
//
// It only modifies data INDIRECTLY through Pager(see `pager.Pager`).
type Storage struct {
}

// FullTableScan returns a Stream(see `storage.Stream`) that retrieves sql.RecordValue lazily.
// `fields` determine the wanted field(s)
//
// **contract**
//
// - `forall field in fields(field.Table.Name == tableName)`
func (s Storage) FullTableScan(tableName string, fields []sql.Field) Stream {
	// TODO:
	panic("NOT IMPLEMENTED")
}

// CreateTable returns nil or an error when
//
// - length of `table.Name` exceeds `sql.MaxTableNameLength`.
//
// - `table.Name` is duplicated
//
// - not all characters in `table.Name` are alphabets or digits. (e.g. `people!`or`@db`)
func (s Storage) CreateTable(table sql.Table) error {
	// TODO:
	panic("NOT IMPLEMENTED")
}

type InsertValue struct {
	Field *sql.Field
	Value sql.Value
}

type Insert struct {
	Table  string
	Values []InsertValue
}

// Insert returns nil or an error when
//
// - `insert.Table` not exists
//
// - forall `i`, `insert.Values[i].Field` is not consistent with table meta information.
func (s Storage) Insert(insert Insert) error {
	// TODO:
	panic("NOT IMPLEMENTED")
}

// Stream is a Record generator that returns `sql.Record` lazily by method `next()`
// `next` returns an error after the last record returned.
//
// **contract**
//
// - `(sql.Record == nil && error != nil) || (sql.Record != nil && error == nil)`
type Stream interface {
	next() (sql.RecordValue, error)
}
