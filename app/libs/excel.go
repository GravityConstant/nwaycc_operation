package libs

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/utils"
	"github.com/tealeg/xlsx"
)

func ExportExcel(data [][]string, filename string) (filePath string, err error) {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("sheet1")
	for r := range data {
		row = sheet.AddRow()
		for c := range data[0] {
			cell = row.AddCell()
			cell.Value = data[r][c]
		}
	}

	if !utils.FileExists("logs") {
		os.MkdirAll("logs", os.ModePerm)
	}
	if len(filename) == 0 {
		filePath = "logs/last.xlsx"
	} else {
		filePath = "logs/" + filename
	}

	err = file.Save(filePath)

	return filePath, err
}

func ReadExcelLastId(filename string) string {
	excelFileName := filename
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf(err.Error())
		return ""
	}
	return xlFile.Sheets[0].Rows[1].Cells[0].Value
}
