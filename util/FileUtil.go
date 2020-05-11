package util

import (
	"io/ioutil"
	"os"
)

/**
 * 判断文件是否存在: 存在返回 true, 不存在返回false
 */
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ReadSQLFromFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile("sql/" + fileName)
	if err != nil {
		return "select '读取SQL文件失败,请核实文件是否存在.'", err
	}
	return string(bytes), err
}
