##### mac下编译linux和windows：
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

##### Windows下编译mac和linux：

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go

##### log目录：
/var/log/rdc.logs/mysql-cluster-monitor-log/

##### 配置log地址的文件：
logger/cmdline.go	
logD = flag.String("logdir", "C:/Temp", "log directory name")

##### 部署说明：
config.conf和main需要在用一个目录下

##### 配置：
config.conf

```js
 //数据库各个节点的配置地址
 --
"hosts" :{ 
        "n1":"user:password@(192.168.2.11:3306)/",
        "n2":"user:password@(192.168.2.194:3306)/",
        "n3":"user:password@(192.168.2.43:3306)/"
    },
 //邮件服务器地址
 --
"dialer": {
        "host": "xx",
        "port": "xx",
        "user": "xx",
        "password": "xx"
    },
 //email的发件人和接收人
 --
"email" : {
        "from": "xx",
        "to": "xx",
        "subject": "mysql 集群预警"
    }

```# mysql-cluster-monitor
