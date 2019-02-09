package sql

// DataType : Types of data stored in database
type DataType int

const (
	// BOOL : boolean value (true / false)
	BOOL DataType = iota
	// INT32 : 32 bits (4 bytes) signed integer
	INT32
	// DOUBLE : double precision (64 bit / 8 bytes) floating point number
	DOUBLE
	// CHAR64 : string with a maximum of 64 bytes (64 ascii characters)
	// (512 bits / 64 bytes)
	CHAR64
)

// Value : dynamicly typed SQL value
type Value interface {
	Type() DataType
	AcceptVisitor(ValueVisitor) interface{}
}

// Int32Value implements Value
type Int32Value struct {
	Value int32
}

func (i Int32Value) Type() DataType {
	return INT32
}

func (i Int32Value) AcceptVisitor(v ValueVisitor) interface{} {
	return v.VisitInt32(i)
}

// BoolValue implements Value
type BoolValue struct {
	Value bool
}

func (b BoolValue) Type() DataType {
	return BOOL
}

func (b BoolValue) AcceptVisitor(v ValueVisitor) interface{} {
	return v.VisitBool(b)
}

// DoubleValue implements value
type DoubleValue struct {
	Value float64
}

func (d DoubleValue) Type() DataType {
	return DOUBLE
}

func (d DoubleValue) AcceptVisitor(v ValueVisitor) interface{} {
	return v.VisitDouble(d)
}

// Char64Value implements Value
type Char64Value struct {
	Value [64]byte
}

func (s Char64Value) Type() DataType {
	return CHAR64
}

func (s Char64Value) AcceptVisitor(v ValueVisitor) interface{} {
	return v.VisitChar64(s)
}

type ValueVisitor interface {
	VisitInt32(Int32Value) interface{}
	VisitBool(BoolValue) interface{}
	VisitDouble(DoubleValue) interface{}
	VisitChar64(Char64Value) interface{}
}

// Record is a sequence of sql Value
type Record []Value
