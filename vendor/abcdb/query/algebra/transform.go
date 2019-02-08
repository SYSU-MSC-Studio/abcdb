package algebra

import (
	"abcdb/sql/parser/ast"
)

// FromAST : Construct the relational algebra operator from SQL select statement
//
// **contract**
//
// - input sql is guaranteed (or assumed) to be perfectly valid
// - output operator should express identical semantic of input SQL
func FromAST(sql ast.Select) Operator {
	// TODO
	return transform(nil)
}

// basic relational algebra equivalent transformation.
// see [wiki](https://en.wikipedia.org/wiki/Relational_algebra#Use_of_algebraic_properties_for_query_optimization)
// for more information
func transform(op Operator) Operator {
	// TODO
	return op
}
