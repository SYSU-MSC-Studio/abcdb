package storage

import (
	"abcdb/sql"
)

// DataStore provids high level interfaces for retrieving and manipulating
//   internal data of the database (which are stored as "tables")
//
// It only modifies data INDIRECTLY through Pager(see `pager.Pager`).
type DataStore interface {
	// LinearScan returns a `Stream` (see `storage.Stream`) that retrieves
	// sql.Record lazily. `fields` determine the wanted field(s)
	//
	// **contract**
	//
	// - `∀ field ∈ fields. field ∈ table.Fields`
	LinearScan(table *sql.Table, fields []sql.Field) (RecordStream, error)

	// Insert returns nil or an error when
	//
	// **contracts**
	//
	// - `∀ value ∈ values. value.Field ∈ into.Fields`
	Insert(into *sql.Table, values []InsertValue) error
}

type InsertValue struct {
	Field *sql.Field
	Value sql.Value
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
