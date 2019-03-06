package algebra

// OperatorVisitor : visitor for operator tree
type OperatorVisitor interface {
	VisitTable(*JustTable) interface{}
	VisitProjection(*Projection) interface{}
	// VisitSelection(*Selection) interface{}
	// VisitInnerJoin(*InnerJoin) interface{}
}
