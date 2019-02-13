package storage

import (
	"abcdb/pager"
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
	Insert(into *sql.Table, values []sql.FieldData) error

	// Flush data specified in `tables` to the disk (with `Pager.Flush(..)`),
	//   if `tables == nil`, data in every table under management are flushed
	Flush(tables []*sql.Table) error
}

// RecordStream is a Record generator that returns `sql.Record` lazily by
// method `Next()`, which returns error after the last record is returned.
//
// **contract**
//
// - `(sql.Record == nil XOR error == nil) == true`
type RecordStream interface {
	Next() (sql.Record, error)
}

// InitDataStore : initialize a `DataStore` instance with a `Pager` instance
//   (see 'abcdb/pager/interface.go')
func InitDataStore(pager pager.Pager) DataStore {
	return SimpleManager{CurrentPager: pager}
}

type SimpleManager struct {
	CurrentPager pager.Pager
	// TODO:
}

func (SimpleManager *SimpleManager) LinearScan(
	table *sql.Table, fields []sql.Field) (RecordStream, error) {
	panic("")
}

func (SimpleManager *SimpleManager) Insert(
	into *sql.Table, values []sql.FieldData) error {
	panic("")
}

func (SimpleManager *SimpleManager) Flush(tables []*sql.Table) error {
	panic("")
}
