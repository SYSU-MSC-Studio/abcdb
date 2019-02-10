package pager

import (
	"math"
)

// ExamplePager : All data are stored in main memory and no paging and buffering
// AT ALL. Only for illustration and testing purpose.
type ExamplePager struct {
	Files map[string][]byte
}

func (p ExamplePager) Read(logicalFile string, offset int, nbytes int) []byte {
	data, exist := p.Files[logicalFile]
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
	store, exist := p.Files[logicalFile]
	dataBoundary := offset + len(data)

	// if no such file, create one
	if !exist {
		p.Files[logicalFile] = make([]byte, dataBoundary+1)
		store = p.Files[logicalFile]
	}

	// if original file is too small, extend it
	if len(store) <= dataBoundary {
		newStore := make([]byte, dataBoundary+1)
		copy(newStore, store)
		store = newStore
		p.Files[logicalFile] = newStore
	}

	// write data to specific location
	copy(store[offset:dataBoundary], data)

	// our simple pager never fail
	return true
}
