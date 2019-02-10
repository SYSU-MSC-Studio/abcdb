package storage

import (
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
}
