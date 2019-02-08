package ast

import (
	"abcdb/sql"
)

// SQL is the top level of ast
type SQL interface {
	AcceptSQLVisitor(visitor SQLVisitor)
}

// Select implements SQL
//
// **contract**
//
// - `Table` is a valid table in database
// - `∀ field ∈ Fields. field.Table == Table`
type Select struct {
	Table  *sql.Table
	Fields []*sql.Field
}

// AcceptSQLVisitor of `SQL`
func (s *Select) AcceptSQLVisitor(visitor SQLVisitor) {
	visitor.VisitSelect(s)
}

// CreateTable implements SQL
type CreateTable struct {
	Table  string
	Fields []FieldInfo
}

// FieldInfo : a pair of name and type specified when creating table
type FieldInfo struct {
	Name string
	Type sql.DataType
}

// AcceptSQLVisitor of `SQL`
func (c *CreateTable) AcceptSQLVisitor(visitor SQLVisitor) {
	visitor.VisitCreateTable(c)
}

// Insert implements SQL
type Insert struct {
	Table  string
	Values []InsertValue
}

// InsertValue : which value insert to which field
//
// **contract**
//
// - `Field.Type == Value.Type()`
type InsertValue struct {
	Field *sql.Field
	Value sql.Value
}

// AcceptSQLVisitor of `SQL`
func (i *Insert) AcceptSQLVisitor(visitor SQLVisitor) {
	visitor.VisitInsert(i)
}
