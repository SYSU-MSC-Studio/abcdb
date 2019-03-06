package ast

// SQLVisitor :
//
// [Visitor pattern](https://en.wikipedia.org/wiki/Visitor_pattern)
// is used to traverse the abstract syntax tree.
type SQLVisitor interface {
	VisitSelect(*Select) interface{}
	VisitCreateTable(*CreateTable) interface{}
	VisitInsert(*Insert) interface{}
}
