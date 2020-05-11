package util

import (
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func Execute(task string, filePath string, fileName string, mailTo []string) {

	subject := task
	body := "Hi,<br /><br />&nbsp;&nbsp;&nbsp;&nbsp;这个是 " + task + " 报表发送邮件，数据在附件，提供了csv和excel两种文件，请查收。<br /><br />&nbsp;&nbsp;&nbsp;&nbsp;系统邮件，请勿回复。<br /><br />"

	start := time.Now()

	log.WithFields(log.Fields{
		"task": task,
	}).Info(task + "任务启动")

	defer func() {
		log.WithFields(log.Fields{
			"task": task,
		}).Info(task + "任务结束")
	}()

	sqlStr, _ := ReadSQLFromFile(fileName)

	columns, totalValues, _ := GetMySQLResult(sqlStr)

	if len(totalValues) > 1 {
		log.WithFields(log.Fields{
			"task": task,
		}).Info("结果集不为空, 发送邮件.")
	} else {
		log.WithFields(log.Fields{
			"task": task,
		}).Info("结果集为空, 不发送邮件并退出.")
		return
	}

	// 去掉结尾的.sql字符
	fileStr := strings.Replace(fileName, ".sql", "", -1) + "_"
	csvFileName := fileStr + start.Format("2006-01-02-150405") + ".csv"
	xlsxFileName := fileStr + start.Format("2006-01-02-150405") + ".xlsx"

	// 写 csv 文件
	_ = WriteToCSV(filePath, csvFileName, columns, totalValues)
	// 写 Excel文件
	_ = WriteToExcel(filePath, xlsxFileName, columns, totalValues)

	// 发邮件
	attach := [...]string{
		filePath + csvFileName,
		filePath + xlsxFileName,
	}
	_ = SendFileMail(mailTo, subject, body, attach[:])

}
