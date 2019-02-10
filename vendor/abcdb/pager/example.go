package pager

import (
	"math"
)

// ExamplePager : All data are stored in main memory and no paging and buffering
// AT ALL. Only for illustration and testing purpose.
type ExamplePager struct {
	files map[string][]byte
}

func (p ExamplePager) Read(logicalFile string, offset int, nbytes int) []byte {
	data, exist := p.files[logicalFile]
	// no such logical file
	if !exist {
		return nil
	}
	// file is not large enough
	if len(data) <= offset {
		return nil
	}

	dataBoundary := int(
		math.Min(float64(offset+nbytes), float64(len(data)-1)))

	resultLength := dataBoundary - offset + 1

	// returns a COPY of original data whereever they are,
	// which means external modification to result won't affect db internal data
	result := make([]byte, resultLength)
	copy(result, data[offset:dataBoundary])

	return result
}

func (p ExamplePager) Write(logicalFile string, offset int, data []byte) bool {
	store, exist := p.files[logicalFile]
	dataBoundary := offset + len(data)

	// if no such file, create one
	if !exist {
		p.files[logicalFile] = make([]byte, dataBoundary+1)
		store = p.files[logicalFile]
	}

	// if original file is too small, extend it
	if len(store) <= dataBoundary {
		newStore := make([]byte, dataBoundary+1)
		copy(newStore, store)
		store = newStore
		p.files[logicalFile] = newStore
	}

	// write data to specific location
	copy(store[offset:dataBoundary], data)

	// our simple pager never fail
	return true
}

// PageSize of ExamplePager makes no sense
func (p ExamplePager) PageSize() int {
	return math.MaxInt32
}

// Flush for ExamplePager: do nothing
func (p ExamplePager) Flush(file string) error {
	return nil
}

// FlushAll for ExamplePager: do nothing
func (p ExamplePager) FlushAll() error {
	return nil
}
