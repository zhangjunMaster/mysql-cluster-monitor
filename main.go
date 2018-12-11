package main

import (
	"fmt"
	"log"
	"mysql-cluster-monitor/cluster"
	"mysql-cluster-monitor/lib"
	"mysql-cluster-monitor/lib/email"
	"mysql-cluster-monitor/logger"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	count  uint32
	status string
)

func main() {
	config, err := lib.ParseConfig()
	if err != nil {
		logger.LOG_ERROR("%v", err)
		return
	}
	hosts := config["hosts"]
	for {
		cluster := cluster.NewCluster(hosts)
		err = cluster.Connect()
		if err != nil {
			logger.LOG_ERROR("connect err %v", err)
		}
		members, err := cluster.GetMembers()
		if err != nil {
			logger.LOG_ERROR("GetMembers err %v", err.Error())
		}
		logger.LOG_INFO("members %+v", members)
		log.Println("MYSQL CLUSTER MONIT IS STARTED")
		if len(members) != len(hosts) {
			alertEmail, _ := email.NewAlertEmail(config, members)
			alertEmail.Send()
			status = "异常"
			logger.LOG_ERROR("ERROR mysql is missing %+v", members)
		} else {
			status = "正常"
			atomic.StoreUint32(&count, 0)
		}
		str := fmt.Sprintf("[%s]:现存mysql节点数: %d, 配置的节点数: %d", status, len(members), len(hosts))
		log.Println(str)
		logger.LOG_INFO("%s", str)
		time.Sleep(60 * time.Second)
	}
}
