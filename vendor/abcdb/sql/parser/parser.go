package parser

import (
	"abcdb/sql"
	"abcdb/sql/parser/ast"
)

// Parse SQL to a **valid** AST
//
// returns error when:
//
// - the SQL string is lexically or syntastically incorrect
// (e.g. `select !@#$!@#$ from table`, `select column1, column2 ,, from table`)
// - the semantic of sql is incorrect
// (e.g. table referenced in SQL doesn't exist in database, attempting to
// create an duplicate table, inserted value doesn't match its type in the
// table, etc.)
//
// **contract**
//
// - `database != nil`
// - `(ast == nil XOR error == nil) == true`
// - `database` is readonly
func Parse(database *sql.Database, sql string) (ast.SQL, error) {
	// TODO
	return nil, nil
}
