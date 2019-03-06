package pager

import (
	"fmt"
	"log"
	"os"
)

///////////////////////////////////////////////////
//  in this program, pager is a global variable  //
//  so you can change the data of pages          //
///////////////////////////////////////////////////

//TestPrintPager
func testPrintPager() error {
	var i int
	for i = 0; i < p.PageNumber; i++ {
		fmt.Println("page", i, "name", p.files[i].header.FileName, p.files[i].data)
	}
	return nil
}

func flush(FileName string) error {
	var i int

	file, err := os.OpenFile(
		FileName,
		os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var byteSlice []byte
	var offset int64
	var whence int //0 means the original positiohn

	for i = 0; i < p.PageNumber; i++ {
		if p.files[i].header.FileName == FileName {
			offset = int64(p.files[i].header.begin)
			newPosition, err := file.Seek(offset, whence)
			if err != nil {
				log.Fatal(err)
			}
			byteSlice = p.files[i].data
			bytesWritten, err := file.Write(byteSlice) //???? completely still need improving
			if err != nil {
				log.Fatal(err)
			}
			_ = newPosition
			_ = bytesWritten
		}
	}
	return nil
}

//CheckIfInPager does what it is called
func CheckIfInPager(logicalFile string, begin int) Page {
	var i int

	for i = 0; i < p.PageNumber; i++ {
		if logicalFile == p.files[i].header.FileName &&
			begin == p.files[i].header.begin {
			return p.files[i]
		}
	}
	return OpenFile(logicalFile, begin)
}

//OpenFile add a page
func OpenFile(logicalFile string, begin int) Page {
	file, err := os.Open(logicalFile)
	defer file.Close() //defer, run after all code, like stack
	if err != nil {
		log.Fatal(err)
		temp := new(Page)
		return *temp
	}

	var page Page
	useless, err := file.Seek(int64(begin), 0)
	if err != nil {
		log.Fatal(err)
	}
	_ = useless
	page.data = make([]byte, 10)
	bytesWritten, err := file.Read(page.data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	page.header.begin = begin
	println("FileName =", logicalFile)
	page.header.FileName = logicalFile
	page.header.UsedSpace = int(bytesWritten)
	p.files[p.PageNumber] = page
	p.PageNumber++

	var flag bool
	for i := 0; i < p.FileNumber; i++ {
		if p.Name[i] == logicalFile {
			flag = false
			break
		}
	}
	if flag {
		p.Name[p.FileNumber] = logicalFile
		p.FileNumber++
	}
	return page
}
