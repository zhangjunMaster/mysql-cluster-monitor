package repository

import "database/sql"

func Query(db *sql.DB, q string) (map[int]map[string]string, error) {
	result := make(map[int]map[string]string)
	rows2, err := db.Query(q)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	cols, err := rows2.Columns()
	if err != nil {
		return nil, err
	}
	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	for rows2.Next() {
		//填充数据 Query的结果是Rows，方法func (rs *Rows) Scan(dest ...interface{}) error
		//5.将rows2遍历的结果填入到scans中的地址上()
		rows2.Scan(scans...)
		//6.定义每行数据的格式
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	return result, nil
}
