package cluster

import (
	"database/sql"

	"mysql-cluster-monitor/repository"
)

type Cluster struct {
	dbHosts map[string]string
	dbList  map[string]*sql.DB
}

func NewCluster(hosts map[string]string) *Cluster {
	return &Cluster{dbHosts: hosts, dbList: make(map[string]*sql.DB)}
}

// connect
func (c *Cluster) Connect() error {
	for key, con := range c.dbHosts {
		db, err := sql.Open("mysql", con)
		if err != nil {
			return err
		}
		c.dbList[key] = db
	}
	return nil
}

// 关闭mysql
func (c *Cluster) Close() error {
	for _, db := range c.dbList {
		db.Close()
	}
	return nil
}

// 获取集群用户
func (c *Cluster) GetMembers() (result map[int]map[string]string, err error) {
	var q string = "SELECT * FROM performance_schema.replication_group_members"

	for _, db := range c.dbList {
		result, err = repository.Query(db, q)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return result, err
}

// 获取集群主主节点
func (c *Cluster) GetPrimaryMemberID() (result map[int]map[string]string, err error) {
	var q string = "SHOW GLOBAL STATUS LIKE 'group_replication_primary_member'"
	for _, db := range c.dbList {
		result, err = repository.Query(db, q)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return result, err
}
