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
	LinearScan(table *sql.Table, targetfields []sql.Field) (RecordStream, error)

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

// RecordLength : return the byte length of a record
func RecordLength(table *sql.Table) int {
	length := 0
	for _, field := range table.Fields {
		switch datatype := field.Type; datatype {
		case sql.BOOL:
			length++
		case sql.INT32:
			length += 4
		case sql.DOUBLE:
			length += 8
		case sql.CHAR64:
			length += 64
		}
	}
	return length
}

// InitDataStore : initialize a `DataStore` instance with a `Pager` instance
//   (see 'abcdb/pager/interface.go')
func InitDataStore(pager pager.Pager) DataStore {
	return &SimpleManager{CurrentPager: pager}
}

type SimpleManager struct {
	CurrentPager pager.Pager
	// TODO:
}

func (SimpleManager *SimpleManager) LinearScan(
	table *sql.Table, targetfields []sql.Field) (RecordStream, error) {
	panic("")
}

func (SimpleManager *SimpleManager) Insert(
	into *sql.Table, values []sql.FieldData) error {
	orderedrecord := make([]sql.Value, 0)
	for _, value := range values {
		for _, field := range into.Fields {
			if *value.Field == field {
				orderedrecord = append(orderedrecord, value.Value)
			}
		}
	}
	SimpleManager.CurrentPager.Write(
		into.Name, into.NRecords*RecordLength(into), Serialize(orderedrecord))
	return nil
}

func (SimpleManager *SimpleManager) Flush(tables []*sql.Table) error {
	for _, table := range tables {
		SimpleManager.CurrentPager.Flush(table.Name)
	}
	return nil
}
