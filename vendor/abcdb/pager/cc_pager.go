package pager

var p StructPager

//MAX means max number of pages ina pager
const MAX int = 100

//PAGEMAX means max number of bytes in a page
const PAGEMAX int = 10

//FILEMAX means max number of file in a pager
const FILEMAX int = 10

//StructPager is a Pager. This is a comment
type StructPager struct {
	files      [MAX]Page
	Name       map[int]string
	FileNumber int
	PageNumber int
}

//Page is a page from Pager
type Page struct {
	data   []byte
	header PageHeader
}

//PageHeader include Page's information
type PageHeader struct {
	begin     int //this beigin position of the file
	FileName  string
	UsedSpace int
}

//Read is a funciton
func (p StructPager) Read(logicalFile string, offset int, nbytes int) []byte {
	var i int
	var Page Page
	var start, end int //start & end is different from offset, it must % PAGEMAX
	var FirstPage, LastPage int

	result := make([]byte, 0, nbytes)

	boundary := offset + nbytes ///here offset start from 0
	FirstPage = offset/PAGEMAX + 1
	LastPage = boundary/PAGEMAX + 1
	//println(FirstPage, LastPage)
	//
	for i = FirstPage; i <= LastPage; i++ {
		start = (i - 1) * PAGEMAX
		end = i * PAGEMAX
		Page = CheckIfInPager(logicalFile, start)
		start = 0
		end = PAGEMAX
		if i == FirstPage {
			start = offset % PAGEMAX
		}
		if i == LastPage {
			end = boundary % PAGEMAX
		}
		result = append(result, Page.data[start:end]...)

	}
	println(p.files[1].header.FileName)
	return result
}

//Write is a function
func (p StructPager) Write(logicalFile string, offset int, data []byte) (success bool) {
	testPrintPager()
	println(p.files[0].header.FileName)
	var i int
	var Page Page
	var start, end int //start & end is different from offset, it must % PAGEMAX
	var FirstPage, LastPage int
	var nbytes = len(data)
	var temp []byte
	var tempdata []byte

	boundary := offset + nbytes ///here offset start from 0
	FirstPage = offset/PAGEMAX + 1
	LastPage = boundary/PAGEMAX + 1
	//
	for i = FirstPage; i <= LastPage; i++ {
		start = (i - 1) * PAGEMAX
		end = i * PAGEMAX
		Page = CheckIfInPager(logicalFile, start)
		if i == FirstPage {
			start = offset - offset
		}
		if i == LastPage {
			end = boundary - offset
		}
		tempdata = data[start:end]
		start = 0
		end = PAGEMAX
		if i == FirstPage {
			start = offset % PAGEMAX
		}
		if i == LastPage {
			end = boundary % PAGEMAX
		}
		temp = Page.data[end:]
		Page.data = append(Page.data[:start], tempdata...)
		Page.data = append(Page.data, temp...)
	}
	testPrintPager()
	return true
}

// PageSize is
func (p StructPager) PageSize() int {
	// var i int
	// var size int
	// for i = 0; i < p.PageNumber; i++ {
	// 	size += p.files[i].header.UsedSpace
	// }
	return PAGEMAX
}

// Flush for Pager: do nothing
func (p StructPager) Flush(FileName string) error {
	return flush(FileName)

}

// FlushAll for Pager: do nothing
func (p StructPager) FlushAll() error {
	var i int
	println(p.FileNumber)
	for i = 0; i < p.FileNumber; i++ {
		p.Flush(p.Name[i])
	}
	return nil
}
