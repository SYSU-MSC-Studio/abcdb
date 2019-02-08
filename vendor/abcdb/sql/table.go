package sql

// Table : Meta-information for SQL table
type Table struct {
	Name     string
	Fields   []Field
	NRecords int
}

// Field : Meta-information for each field in table.
//
// **contract**
//
// - `∀ field. field ∈ field.Table.Fields`
type Field struct {
	Name  string
	Type  DataType
	Table *Table
}

// Database is a collection of tables
type Database struct {
	Name   string
	Tables map[string]Table
}
