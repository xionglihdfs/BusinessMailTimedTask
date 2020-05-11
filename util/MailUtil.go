package util

import (
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"mime"
	"strconv"
	"strings"
)

func SendFileMail(mailTo []string, subject string, body string, attach []string) error {

	log.WithFields(log.Fields{
		"subject": subject,
		"mailTo":  mailTo,
	}).Info("开始发送<" + subject + ">邮件")

	defer func() {
		log.WithFields(log.Fields{
			"subject": subject,
			"mailTo":  mailTo,
		}).Info("完成邮件<" + subject + ">发送")
	}()

	m := gomail.NewMessage()

	// 发件人信息
	m.SetHeader("From", m.FormatAddress(GetConfigString("mailFrom"), "报表系统"))
	// 收件人
	m.SetHeader("To", mailTo...)
	// 主题
	m.SetHeader("Subject", subject)
	// 内容
	m.SetBody("text/html", body)
	// 附件
	for i := 0; i < len(attach); i++ {
		m.Attach(attach[i],
			gomail.Rename(strings.Replace(attach[i], "result/", "", -1)),
			gomail.SetHeader(map[string][]string{
				"Content-Disposition": {
					fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", strings.Replace(attach[i], "result/", "", -1))),
				},
			}))
	}

	smtpPort, err := strconv.Atoi(GetConfigString("smtpPort"))
	if err != nil {
		panic(err)
	}

	d := gomail.NewDialer(GetConfigString("smtpKD"),
		smtpPort,
		GetConfigString("smtpUsername"),
		GetConfigString("smtpPassword"))

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = d.DialAndSend(m)
	if err != nil {
		panic(err)
	}

	return err
}
