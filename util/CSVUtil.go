package util

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"os"
)

func WriteToCSV(filePath string, fileName string, columns []string, totalValues [][]string) error {

	log.WithFields(log.Fields{
		"file": filePath + fileName,
	}).Info("开始导入CSV文件")
	defer func() {
		log.WithFields(log.Fields{
			"file": filePath + fileName,
		}).Info("结束导入CSV文件")
	}()

	csvFile := filePath + fileName

	f, err := os.Create(csvFile)
	defer func() {
		if f != nil {
			_ = f.Close()
		}
	}()

	if err != nil {
		//panic(err)
		return err
	}

	// 写入 csv 文件
	_, _ = f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	for i, row := range totalValues {
		//第一次写列名+第一行数据
		if i == 0 {
			_ = w.Write(columns)
			_ = w.Write(row)
		} else {
			_ = w.Write(row)
		}
	}
	w.Flush()

	return nil
}
