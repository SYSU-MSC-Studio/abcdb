package query

import (
	"abcdb/query/algebra"
	"abcdb/sql"
)

// PlanAndOptimize : Make plans and optimize them,
// then try to pick the cheapest one to execute
func PlanAndOptimize(op algebra.Operator) Plan {
	return PickBestPlan(MakeQueryPlans(op))
}

// MakeQueryPlans returns several possible operation plans according to the
// relational algebra and table information (inside of operator structures)
func MakeQueryPlans(op algebra.Operator) []Plan {
	// TODO
	return nil
}

// PickBestPlan estimates the costs of all plans and returns the "cheapest"
func PickBestPlan(plans []Plan) Plan {
	// TODO
	return nil
}

// Plan : Evaluation plan used to guide execution of query
type Plan interface {
	Evaluate(Evaluator)

	// Estimated number of how many rows this plan would traverse,
	// a simple property of the query plan
	EstimatedRows() int

	// TODO: (future requirement) more relevant properties of query plans.
}

// LinearScan : traverse the table linearly and retrieve every row of it
//
// **contract**
//
// - `∀ scan. scan.fields ⊆ scan.Table.fields`
type LinearScan struct {
	Table  *sql.Table
	Fields []*sql.Field
}

// Evaluate a LinearScan
func (l *LinearScan) Evaluate(e Evaluator) {
	e.EvaluateLinearScan(l)
}

// EstimatedRows of LinearScan is exactly its table size
func (l *LinearScan) EstimatedRows() int {
	return l.Table.NRecords
}

/**
type Filter struct {
}

type IndexScan struct {
}

type Join struct {
}
**/
