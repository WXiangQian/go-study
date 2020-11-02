package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlObj struct {
	// sql.DB 表示一个连接池
	MysqlPool *sql.DB
}

func NewMysql() *MysqlObj {
	mysqlConf := make(map[string]string)
	mysqlConf["host"] = "127.0.0.1"
	mysqlConf["port"] = "3306"
	mysqlConf["pass"] = "123456"
	mysqlConf["user"] = "root"
	mysqlConf["db"] = "laravel-api"
	// sql.Open 的第一个参数是驱动名称，这里是 "mysql"
	// sql.Open 的第二个参数是数据源名称，这里通过 mysql.Config 结构来配置，然后调用 FormatDSN 方法得出数据源名称为："root:xxxxxx@tcp(127.0.0.1:3306)/mydb"

	pool, err := sql.Open("mysql", mysqlConf["user"]+":"+mysqlConf["pass"]+"@tcp("+mysqlConf["host"]+":"+mysqlConf["port"]+")/"+mysqlConf["db"]+"?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	pool.SetMaxIdleConns(300)
	pool.SetMaxOpenConns(300)
	pool.SetConnMaxLifetime(time.Second * 10)
	return &MysqlObj{pool}
}

func main() {

	mysqlObj := NewMysql()

	var num = 300
	var insertCount = 3000
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))
	for j := 0; j <= num; j++ {

		var insertSql = "insert into articles(user_id,type_id,title,content,created_at,updated_at) values"
		for i := 0; i <= insertCount; i++ {

			logData := make(map[string]string)
			logData["user_id"] = "1"
			logData["type_id"] = strconv.Itoa(rand.Intn(3) + 1)
			logData["title"] = "title"
			logData["content"] = "content"
			logData["created_at"] = time.Now().Format("2006-01-02 15:04:05")
			logData["updated_at"] = time.Now().Format("2006-01-02 15:04:05")

			//tmpstr :=  "('" + logData["user_id"] + "','" + logData["type_id"] + "','" + logData["title"] + "','" + logData["content"] + "','" + logData["created_at"] + "','" + logData["updated_at"] +"'),"
			tmpstr := "(" + logData["user_id"] + "," + logData["type_id"] + "," + logData["title"] + "," + logData["content"] + ",'" + logData["created_at"] + "','" + logData["updated_at"] + "'),"
			insertSql += tmpstr

			if i == insertCount {
				tmp := strings.TrimRight(insertSql, ",") + ";"

				InsertTestData(tmp, mysqlObj)
				insertSql = "insert into articles(user_id,type_id,title,content,created_at,updated_at) values"
			}
		}
	}

	fmt.Print(time.Now().Format("2006-01-02 15:04:05"))

}

func InsertTestData(sql string, mysqlObj *MysqlObj) {

	//fmt.Println(sql)
	_, err := mysqlObj.MysqlPool.Exec(sql)
	fmt.Println(err)
}
