package sql

// This file defines `Record.Serialize` and `Deserializer that convert `[]byte`
// from and to `sql.Record`
//
// **contract**
//
// - Deserialize() and Serialize() are inverse functions (in a rough sense)
//   i.e. `∀ bytes: []byte, record: Record, table: Table`
//        `    table.MakeDeserializer().Deserialize(bytes) = (record, nil)`
//        ` -> record.Serialize() = bytes`

// Serializable object can Serialize itself to a byte array
type Serializable interface {
	Serialize() []byte
}

// Serialize Record to bytes
func (record Record) Serialize() []byte {
	// TODO:
	panic("NOT IMPLEMENTED.")
}

// Deserializer can Parse a slice of byte to sql.RecordValue
type Deserializer interface {
	// Deserialize returns error when format of the bytes is not correct w.r.t.
	// what `Record.Serialize()` does.
	//
	// **contract**
	//
	// - `(record == nil && error != nil) || (record != nil && error == nil)`
	Deserialize([]byte) (Record, error)
}

// MakeDeserializer return a specific TableParser according to `table`
func (table *Table) MakeDeserializer() Deserializer {
	// TODO:
	panic("NOT IMPLEMENTED.")
}
