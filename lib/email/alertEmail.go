package email

import (
	"fmt"
	"log"
	"mysql-cluster-monitor/logger"
	"sync/atomic"
)

var (
	count uint32
)

type AlertEmail struct {
	*Email
	Members map[int]map[string]string
}

func NewAlertEmail(config map[string]map[string]string, members map[int]map[string]string) (*AlertEmail, error) {
	m, err := NewEmail(config)
	if err != nil {
		return nil, err
	}
	return &AlertEmail{
		Email:   m,
		Members: members,
	}, nil

}

func (ae *AlertEmail) Send() (err error) {
	value := atomic.LoadUint32(&count)
	if value < 2 {
		ae.Email.Body = ae.ParseMysqlClusterToHtml()
		err = ae.SendEmail()
		if err != nil {
			log.Println("Email SendMail Err:", err)
			logger.LOG_ERROR("Email SendMail Err: %v", err)
		} else {
			atomic.AddUint32(&count, 1)
		}
	}
	return
}

func (ae *AlertEmail) ParseMysqlClusterToHtml() string {
	var str string
	var title string = "<p>MYSQL CLUSTER MONITOR EXCEPTION:</p>"
	for _, member := range ae.Members {
		str += fmt.Sprintf(
			"<p>MEMBER_PORT: %s,MEMBER_STATE: %s,CHANNEL_NAME: %s,MEMBER_ID: %s,MEMBER_HOST: %s</p>",
			member["MEMBER_PORT"],
			member["MEMBER_STATE"],
			member["CHANNEL_NAME"],
			member["MEMBER_ID"],
			member["MEMBER_HOST"],
		)
	}
	message := fmt.Sprintf("[异常]:现存mysql节点数: %d, 配置的节点数: %d", len(ae.Members), len(ae.Email.Hosts))
	return title + str + message
}
