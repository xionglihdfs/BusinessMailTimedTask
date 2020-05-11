package util

import (
	log "github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"os"
)

func WriteToExcel(filePath string, fileName string, columns []string, totalValues [][]string) error {

	log.WithFields(log.Fields{
		"file": filePath + fileName,
	}).Info("开始导入Excel文件")
	defer func() {
		log.WithFields(log.Fields{
			"file": filePath + fileName,
		}).Info("结束导入Excel文件")
	}()

	excelFile := filePath + fileName

	f, err := os.Create(excelFile)
	defer func() {
		if f != nil {
			_ = f.Close()
		}
	}()

	if err != nil {
		return err
	}

	// 写入 Excel 文件
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")

	row = sheet.AddRow()
	for column := range columns {
		cell = row.AddCell()
		cell.Value = columns[column]
	}

	for i := 0; i < len(totalValues); i++ {
		row = sheet.AddRow()
		for j := 0; j < len(columns); j++ {
			cell = row.AddCell()
			cell.Value = totalValues[i][j]
			_ = file.Save(excelFile)
		}
	}

	return err
}
