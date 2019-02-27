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
	// TODO:
	return AlgebraicTransform(nil)
}

// AlgebraicTransform performs basic transformation based on algebraic
//   properties of relational algebra trying to positively optimize the
//   structure of the operator.
func AlgebraicTransform(op Operator) Operator {
	// TODO: try to compose possible transformations implemented
	return op
}

// SquashRedundantProjection is one of the equivalent transformation that is
//   almost always good.
// See https://en.wikipedia.org/wiki/Relational_algebra#Basic_projection_properties
func SquashRedundantProjection(op Operator) Operator {
	// TODO: although usually this is not useful in real life, implement it as
	//   a practice.
	return nil
}
