package algebra

import (
	"abcdb/sql"
)

// Operator : Relational algebra operator
type Operator interface {
	// RelatedTables returns all the tables involved in the operations
	// (select, projection, join etc.).
	// It's not usually useful in the implementation of other functions,
	// but it's convenient for expressing some contracts for `Operator`
	RelatedTables() []*sql.Table

	AcceptVisitor(OperatorVisitor) interface{}

	// TODO: there may be other helper functions who reflect something of
	//   the operator that may be useful for some equivalent transformations.
	//   Feel free to add them if you need one.
}

// Projection implements operator
//
// **contract**
//
// - `∀ proj: Projection, field ∈ proj.Fields.`
//   `field.Table ∈ proj.RelatedTables()`
type Projection struct {
	Fields []*sql.Field
	Child  Operator
}

// RelatedTables of a projection is the related tables of its child operator
func (p *Projection) RelatedTables() []*sql.Table {
	return p.Child.RelatedTables()
}

func (p *Projection) AcceptVisitor(v OperatorVisitor) interface{} {
	return v.VisitProjection(p)
}

// JustTable implements Operator, is the leaf of operator tree
type JustTable struct {
	Table *sql.Table
}

// RelatedTables of a table is just itself
func (j *JustTable) RelatedTables() []*sql.Table {
	return []*sql.Table{j.Table}
}

func (j *JustTable) AcceptVisitor(v OperatorVisitor) interface{} {
	return v.VisitTable(j)
}

/**
type Selection struct {
	// TODO: future requirement
}

type InnerJoin struct {
	// TODO: future requirement
}
**/
