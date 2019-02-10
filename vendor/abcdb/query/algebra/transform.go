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
	return BasicTransformer{}.Transform(nil)
}

// Transformer performs algebra equivalent transformations
// to a relational algebra operator tree (see query/algebra/operators.go for operator tree).
type Transformer interface {
	Transform(Operator) Operator
}

// BasicTransformer performs very basic (or even trivial) transformations
// to relational algebra operator trees.
//
// see [wiki](https://en.wikipedia.org/wiki/Relational_algebra#Use_of_algebraic_properties_for_query_optimization)
// for more information.
type BasicTransformer struct{}

func (t BasicTransformer) Transform(root Operator) Operator {
	// TODO:
	return root
}
