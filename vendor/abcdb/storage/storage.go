package storage

import "abcdb/sql"

// Storage : Storage manager providing high level interfaces for
// manipulating db internal data.
//
// It only modifies data INDIRECTLY through Pager(see `pager.Pager`).
type Storage struct {
	// TODO
}

// LinearScan returns a `Stream` (see `storage.Stream`) that retrieves
// sql.Record lazily. `fields` determine the wanted field(s)
//
// **contract**
//
// - `∀ field ∈ fields. field ∈ table.Fields`
func (s *Storage) LinearScan(table *sql.Table, fields []sql.Field) Stream {
	// TODO:
	panic("NOT IMPLEMENTED")
}

// CreateTable returns error when
//
// - `table.Name` is duplicated
func (s *Storage) CreateTable(table sql.Table) error {
	// TODO:
	panic("NOT IMPLEMENTED")
}

type InsertValue struct {
	Field *sql.Field
	Value sql.Value
}

// Insert returns nil or an error when
//
// **contracts**
//
// - `∀ value ∈ values. value.Field ∈ into.Fields`
func (s Storage) Insert(into *sql.Table, values []InsertValue) {
	// TODO:
	panic("NOT IMPLEMENTED")
}

// RecordStream is a Record generator that returns `sql.Record` lazily by
// method `Next()`, which returns error after the last record is returned.
//
// **contract**
//
// - `(sql.Record == nil && error != nil) || (sql.Record != nil && error == nil)`
type RecordStream interface {
	Next() (sql.RecordValue, error)
}
