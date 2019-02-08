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
}

// Int32Value implements Value
type Int32Value struct {
	Value int32
}

// Type for `Value`
func (i Int32Value) Type() DataType {
	return INT32
}

// BoolValue implements Value
type BoolValue struct {
	Value bool
}

// Type for `Value`
func (b BoolValue) Type() DataType {
	return BOOL
}

// DoubleValue implements value
type DoubleValue struct {
	Value float64
}

// Type for `Value`
func (d DoubleValue) Type() DataType {
	return DOUBLE
}

// Char64Value implements Value
type Char64Value struct {
	Value [64]byte
}

// Type for `Value`
func (c Char64Value) Type() DataType {
	return CHAR64
}
