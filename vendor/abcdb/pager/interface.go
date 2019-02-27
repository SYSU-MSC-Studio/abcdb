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

	// PageSize returns the size (in bytes) for each page managed by this pager
	PageSize() int

	// Flush forces the pager to write all (dirty) data in main memory related
	//   to the specified logicalFile to the disk whereever they are stored.
	Flush(logicalFile string) error

	// Like `Flush()`, except all dirty data are flushed to the disk.
	FlushAll() error

	//only I would use this function
	OpenFile()
}
