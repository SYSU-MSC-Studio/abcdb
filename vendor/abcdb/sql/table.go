package sql

// Table : Meta-information for SQL table
type Table struct {
	Name   string
	Fields []Field
}

// Field : Meta-information for each field in table.
//
// **contract**
//
// - `this ∈ Table.Fields`
type Field struct {
	Name  string
	Type  DataType
	Table *Table
}

// Database is a collection of tables
type Database struct {
	Name   string
	Tables []Table
}
