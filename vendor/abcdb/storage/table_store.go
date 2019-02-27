package storage

import (
	"abcdb/pager"
	"abcdb/sql"
)

// TableStore manages the "meta-data" of `Table`s and `Database`s
// (see 'abcdb/sql/table.go').
//
// **contract**
//
// - ∀ tableStore, table.
//     if tableStore.CreateTable(table) == nil
//        table ∈ tableStore.Database().Tables
type TableStore interface {
	// CreateTable returns error when
	//
	// - `table.Name` is duplicated
	CreateTable(table sql.Table) error

	// Database returns the meta-data describing the running database
	Database() *sql.Database

	// Flush all database information to the disk (with `Pager.Flush(..)`)
	Flush() error
}

// InitTableStore : initialize a `TableStore` instance with a `Pager` instance
//   (see 'abcdb/pager/interface.go')
func InitTableStore(pager pager.Pager) TableStore {
	// TODO
	return nil
}
