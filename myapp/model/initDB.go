package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB


func init()  {

	var err error

	DB, err = sql.Open("mysql", "root:123456@(127.0.0.1:3306)/test2?charset=utf8&parseTime=true&loc=Local")

	if err != nil {
		log.Fatal("数据库打开出现了问题：", err)
		return
	}
	//	defer db.Close()
	err = DB.Ping()
	if err != nil {
		log.Fatal("数据库连接出现了问题：", err)
		return
	}
}
