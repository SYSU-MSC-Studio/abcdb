package pager

// Pager provide abstraction for low level file management.
//
// **constracts**
//
// - `∀ pager, file, offset, data != nil`
//   `pager.WriteData(file, offset, data);`
//   `DeepEqual(pager.ReadData(file, offset, len(data)), data)`
// - many more .. see `ExamplePager`
type Pager interface {
	Read(logicalFile string, offset int, nbytes int) []byte
	Write(logicalFile string, offset int, data []byte) (success bool)
}
