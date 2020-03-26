package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"time"

	"github.com/tealeg/xlsx"
)

type Dtexcel struct {
	Row        bool
	Class      string
	Id         string
	Name       string
	Event      string
	ArrTime    string
	ReturnTime string
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

/*func readxlsx(filn string) []Dtexcel {
	xlFile, err := xlsx.OpenFile(filn)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	sheet := xlFile.Sheet["Sheet1"]
	var DT []Dtexcel
	for r := 1; r < sheet.MaxRow; r++ {
		var dt Dtexcel
		//cel, _ := sheet.Cell(r, 0)
		text0 := (r % 2) != 0
		text1, _ := sheet.Cell(r, 0)
		text2, _ := sheet.Cell(r, 1)
		text3, _ := sheet.Cell(r, 2)
		text4, _ := sheet.Cell(r, 3)
		text5, _ := sheet.Cell(r, 4)
		text6, _ := sheet.Cell(r, 5)
		dt = Dtexcel{text0, text1.Value, text2.Value, text3.Value, text4.Value, text5.Value, text6.Value}
		DT = append(DT, dt)
	}
	return DT
}*/

func readxlsx(filn string) []*Dtexcel {
	var dataSlice [][][]string
	dataSlice, _ = xlsx.FileToSlice(filn)
	DT := []*Dtexcel{}
	for i, v := range dataSlice[0] {
		if i != 0 && v[0] != "" {
			dt := &Dtexcel{(i % 2) != 0, v[0], v[1], v[2], v[3], v[4], v[5]}
			DT = append(DT, dt)
		}
	}
	return DT
}

func main() {
	DT := readxlsx("./Data/excel.xlsx")

	//fmt.Printf("%+V", DT)

	outfile := "./outfile_" + time.Now().Format("20060102150405") + ".html"
	copy("./Data/index.html", outfile)
	tpl, err := template.ParseFiles(outfile)
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.Create(outfile)
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	err = tpl.Execute(f, map[string][]*Dtexcel{"Data": DT})
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()
}
