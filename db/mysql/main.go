package dbutils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn = "root:V69Gz4wsysY@tcp(127.0.0.1:3306)/gomysql"
)

// 全局变量
var db *sql.DB

func init() {
	NewConn(dsn)
}

// NewConn return a DB conn
func NewConn(dsn string) (err error) {
	// sql.Open 不会
	db, err = sql.Open("mysql", dsn)
	if err != nil { //检查 dsn 格式是否正确s
		log.Println(err)
		return err
	}

	err = db.Ping() // 检查数据库是否连接成功
	if err != nil {
		log.Println(err)
		return err
	}

	db.SetMaxOpenConns(10) // 设置最大连接数
	db.SetMaxIdleConns(4)  // 设置最大空闲连接数

	return nil
}
