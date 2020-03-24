package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"log"
	"os"
	"text/template"
	"time"
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

func main() {
	//fmt.Println("vim-go")
	excelFileName := "./Data/excel.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	sheet := xlFile.Sheet["Sheet1"]
	//fmt.Printf("open failed: %s\n", sheet)
	var DT []Dtexcel
	for r, row := range sheet.Rows {
		var dt Dtexcel
		if r != 0 {
			/*for j, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}*/
			text0 := (r % 2) != 0
			text1 := row.Cells[0].String()
			text2 := row.Cells[1].String()
			text3 := row.Cells[2].String()
			text4 := row.Cells[3].String()
			text5 := row.Cells[4].String()
			text6 := row.Cells[5].String()
			dt = Dtexcel{text0, text1, text2, text3, text4, text5, text6}
			DT = append(DT, dt)
		}
	}
	//fmt.Printf("%s\n", DT)
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

	//tpl.Funcs(template.FuncMap{"mod": func(i, j int) bool { return i%j != 0 }})
	err = tpl.Execute(f, map[string][]Dtexcel{"Data": DT})
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()

}
