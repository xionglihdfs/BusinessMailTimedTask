package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func GetMySQLResult(sqlString string) ([]string, [][]string, error) {

	log.WithFields(log.Fields{
		"MySQL": "执行SQL查询",
	}).Info("开始查询数据库")

	defer func() {
		log.WithFields(log.Fields{
			"MySQL": "结束SQL查询",
		}).Info("结束MySQL查询")
	}()

	//存所有行的内容totalValues
	totalValues := make([][]string, 0)

	mysqlPort, err := strconv.Atoi(GetConfigString("mysqlPort"))
	if err != nil {
		panic(err)
	}

	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		GetConfigString("mysqlUsername"),
		GetConfigString("mysqlPassword"),
		GetConfigString("mysqlIp"),
		mysqlPort,
		GetConfigString("database"),
		GetConfigString("charset")))
	defer db.Close()

	rows, err := db.Query(sqlString)

	if err != nil {
		panic(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	//values：一行的所有值,把每一行的各个字段放到values中，values长度==列数
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {

		//存每一行的内容
		var s []string

		//把每行的内容添加到scanArgs，也添加到了values
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for _, v := range values {
			s = append(s, string(v))
		}
		totalValues = append(totalValues, s)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return columns, totalValues, err
}
