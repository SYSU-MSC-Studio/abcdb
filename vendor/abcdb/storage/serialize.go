package storage

import (
	"abcdb/sql"
)

// This file defines `Record.Serialize` and `Deserializer that convert `[]byte`
// from and to `sql.Record`
//
// **contract**
//
// - Deserialize() and Serialize() are inverse functions (in a rough sense)
//   i.e. `∀ bytes: []byte, record: Record, table: Table`
//        `    table.MakeDeserializer().Deserialize(bytes) = (record, nil)`
//        ` -> Serialize(record) = bytes`

// Serializable object can Serialize itself to a byte array
type Serializable interface {
	Serialize() []byte
}

// Serialize Record to bytes
func Serialize(record sql.Record) []byte {
	// TODO:
	panic("NOT IMPLEMENTED.")
}

// Deserializer can Parse a slice of byte to sql.Record
type Deserializer interface {
	// Deserialize returns error when format of the bytes is not correct w.r.t.
	// what `Record.Serialize()` does.
	//
	// **contract**
	//
	// - `(record == nil && error != nil) || (record != nil && error == nil)`
	Deserialize([]byte) (sql.Record, error)
}

// MakeDeserializer returns a specific Deserializer for the records of `table`
func MakeDeserializer(table *sql.Table) Deserializer {
	// TODO:
	panic("NOT IMPLEMENTED.")
}
