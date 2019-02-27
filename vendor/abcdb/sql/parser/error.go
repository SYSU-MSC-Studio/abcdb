package parser

// ErrorCode categorizes the general type of error
//
// e.g. (delete them after finish)
// - invalid identifier in "seleeect field from table"
// - invalid field in "select NonExistentField from table"
// - invalid table in "select * from NonExistentTable"
type ErrorCode int

const (
	// OK : No error
	OK ErrorCode = iota
	// TODO: Add error codes by need
)

// Error : the error object containing specific information about error
type Error struct {
	Code ErrorCode
	// TODO: Design the structure of parser errors
}

// Error interface for go
func (e Error) Error() string {
	// TODO: stringize parser error
	return ""
}
