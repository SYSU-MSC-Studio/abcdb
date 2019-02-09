package sql

// NOTICE:
// This file defines TableParser that parses `[]byte` to `sql.RecordValue`
// For parser that parses user input to AST, see Package Parser.

// Serializable object can Serialize itself to a slice of byte
type Serializable interface {
	Serialize() []byte
}

func (record RecordValue) Serialize() []byte {
	// TODO:
	panic("NOT IMPLEMENTED.")
}

// TableParser can Parse a slice of byte to sql.RecordValue
// and return an error on failure.
//
// **contract**
//
// - `(RecordValue == nil && error != nil) || (RecordValue != nil && error == nil)`
//
// - forall `recordValue` and its relative `table` `(table.MakeParser().Parse(recordValue.Serialize()) == recordValue)`
// NOTICE. this means Parse() and Serialize() are reciprocal.
type TableParser interface {
	Parse([]byte) (RecordValue, error)
}

// MakeParser return a specific TableParser according to `table`
func (table Table) MakeParser() TableParser {
	// TODO:
	panic("NOT IMPLEMENTED.")
}
