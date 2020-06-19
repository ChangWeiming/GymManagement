package mysql

import (
	"database/sql"
	"log"

	//mysql interfaces for query
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// RunMySQL connect mysql on the server.
func RunMySQL() {
	var err error
	db, err = sql.Open("mysql", "Gym:gym2333@tcp(localhost:3306)/?charset=utf8")
	res, err2 := db.Query("use Gym")

	if err2 != nil {
		log.Println(err2)
	} else {
		res.Close()
	}
	if err != nil {
		log.Println(err)
	}
}

//GetDB returns db poiter
func GetDB() *sql.DB {
	return db
}

//CloseMySQL close mysql connection
func CloseMySQL() {
	db.Close()
}

//GetResult accept SQL query result return map of it
func GetResult(query *sql.Rows) (map[int]map[string]string, error) {
	column, _ := query.Columns()              //读出查询出的列字段名
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到bye里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for query.Next() {
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			log.Println(err)
			return nil, err
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	return results, nil
}
