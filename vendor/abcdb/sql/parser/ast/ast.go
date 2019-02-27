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

//AcceptSQLVisitor is
func (s *Select) AcceptSQLVisitor(visitor SQLVisitor) interface{} {
	return visitor.VisitSelect(s)
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

//AcceptSQLVisitor is
func (c *CreateTable) AcceptSQLVisitor(visitor SQLVisitor) interface{} {
	return visitor.VisitCreateTable(c)
}

// Insert implements SQL
type Insert struct {
	Table  string
	Values []sql.FieldData
}

//AcceptSQLVisitor is
func (i *Insert) AcceptSQLVisitor(visitor SQLVisitor) interface{} {
	return visitor.VisitInsert(i)
}
