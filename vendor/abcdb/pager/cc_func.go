package pager

import (
	"fmt"
	"log"
	"os"
)

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
	testPrintPager()
	println(p.files[0].header.UsedSpace)
	file, err := os.OpenFile(
		FileName,
		os.O_WRONLY,
		0666,
	)
	println("I'm here!")
	println(FileName)
	println(p.files[1].header.FileName)
	println(p.files[1].header.FileName == FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	////////
	var byteSlice []byte
	var offset int64
	var whence int //0 means the original positiohn
	/////////
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

//CheckIfInPager do what it is called
func CheckIfInPager(logicalFile string, begin int) Page {
	var i int
	println(p.files[0].header.FileName)
	println(p.PageNumber)
	for i = 0; i < p.PageNumber; i++ {
		if logicalFile == p.files[i].header.FileName &&
			begin == p.files[i].header.begin {
			println("page is here!")
			println(p.files[i].header.FileName)
			//fmt.Println(i, "##", p.files[i].data)
			return p.files[i]
		}
	}
	//println("hhhere!")
	return OpenFile(logicalFile, begin)
}

//OpenFile add a page
func OpenFile(logicalFile string, begin int) Page {
	println(p.files[0].header.FileName)
	file, err := os.Open(logicalFile)
	defer file.Close() //defer, run after all code, like stack
	if err != nil {
		log.Fatal(err)
		temp := new(Page)
		return *temp
	}
	/////////////
	var page Page

	useless, err := file.Seek(int64(begin), 0)
	if err != nil {
		log.Fatal(err)
	}
	_ = useless
	///////
	page.data = make([]byte, 10)
	bytesWritten, err := file.Read(page.data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)
	////////
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
