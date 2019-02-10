package storage

import (
	"abcdb/pager"
	"abcdb/sql"
)

// Storage : Storage manager providing high level interfaces for
// manipulating db internal data.
//
// It only modifies data INDIRECTLY through Pager(see `pager.Pager`).
type Storage struct {
	// TODO
	Pager pager.Pager
}

// LinearScan returns a `Stream` (see `storage.Stream`) that retrieves
// sql.Record lazily. `fields` determine the wanted field(s)
//
// **contract**
//
// - `∀ field ∈ fields. field ∈ table.Fields`
func (s *Storage) LinearScan(
	table *sql.Table, fields []sql.Field) RecordStream {
	// TODO:
	panic("NOT IMPLEMENTED")
}

// CreateTable returns error when
//
// - `table.Name` is duplicated
func (s *Storage) CreateTable(table sql.Table) error {
	// TODO:
	if s.Pager == nil {
		panic("No Pager!")
	}
	s.Pager.Write("TableName", 0, []byte(table.Name))
	offset := 0
	for n, f := range table.Fields {
		s.Pager.Write("FieldsType", n, []byte{byte(f.Type)})
		s.Pager.Write("FieldsName", offset, []byte(f.Name))
		offset += len([]byte(f.Name))
		s.Pager.Write("FieldsName", offset, []byte("\t"))
		offset++
	}
	return nil
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
	Next() (sql.Record, error)
}
