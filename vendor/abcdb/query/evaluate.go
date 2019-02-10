package query

import (
	"abcdb/sql/parser/ast"
	"abcdb/storage"
	"io"
)

// Evaluator is
type Evaluator struct {
	Storage *storage.DataStore
	// TODO
}

// Execute the SQL statement
//
// A `Select` query (or subquery) is first transformed to relational algebra
//   (see 'abcdb/query/algebra/'), and then corresponding query plan is made,
//   which is finally evaluated with the `Evaluator`
//
// There're other queries other than `Select`, which are mostly relatively
//   straightforward to implement with the help of the interfaces of Storages.
//
// During the execution,
//   appropriate text output is generated and write to the `io.Writer`.
//   And errors yielded by low level interfaces are appropriately handled.
func Execute(
	ds *storage.DataStore, ts *storage.TableStore,
	sql ast.SQL, output io.Writer) error {
	// TODO
	return nil
}
