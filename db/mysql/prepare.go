// package dbutils

// func prepareInsertDemo() {
// 	sqlStmt := `insert into user(name, age) values(?,?);`

// }

package dbutils

import (
	"fmt"
	"log"
)

func Prepare() {
	sqlStmt := `insert into user(name, age) values(?,?);`

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	userList := map[string]int{
		"wangwu":     20,
		"zhangqiang": 10,
		"luojin":     30,
		"liuming":    15,
	}

	for name, age := range userList {
		ret, err := stmt.Exec(name, age)
		if err != nil {
			log.Println(err)
		}
		id, _ := ret.LastInsertId()
		n, _ := ret.RowsAffected()
		fmt.Printf("受影响 %d 行 : 插入值 ID = %d \n", n, id)

	}
}
