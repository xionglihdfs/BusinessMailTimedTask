package main

import (
	"businessmailtimingtask/util"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

const (
	// 删除文件时间, 单位小时
	diffTime = 24 * 7
	// 文件路径
	filePath = "result/"
)

var (
	// 2020-05-11 业务方报表数据
	// 业务方报表数据邮箱主题
	businessStatementTaskTitle = "业务方报表数据"
	// 导出黑名单废弃资源SQL文件
	businessStatementFileName = "业务方报表数据(每天上午9点执行).sql"
	// 业务方报表数据邮箱收件人, 支持一个和多个邮箱
	businessStatementMailTo = []string{
		"12345@qq.com",
		"23456@163.com",
		"34567@sina.com",
	}
)

func main() {

	log.WithFields(log.Fields{
		"Main": "main",
	}).Info("程序启动")

	ch := make(chan os.Signal)

	// 定时任务
	c := cron.New()

	log.WithFields(log.Fields{}).Info("业务方报表数据")
	// 每天9点执行一次
	_ = c.AddFunc("0 0 9 * * * ", businessStatement)

	// 删除csv和Excel文件
	_ = c.AddFunc("0 0 0 * * * ", cleanResultFiles)

	// 开始定时任务
	c.Start()

	//阻塞直至有信号传入
	<-ch

}

func businessStatement() {
	util.Execute(businessStatementTaskTitle, filePath, businessStatementFileName, businessStatementMailTo)
}

// 定时删除7天以外的报表文件
func cleanResultFiles() {
	rd, _ := ioutil.ReadDir(filePath)
	for _, fi := range rd {
		if !fi.IsDir() {
			if util.CheckFileIsExist(filePath + fi.Name()) {
				if time.Now().Sub(fi.ModTime()).Hours() >= diffTime {
					log.WithFields(log.Fields{
						"Main":     "main",
						"filename": filePath + fi.Name(),
					}).Info("删除文件")
					err := os.Remove(filePath + fi.Name())
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
}
