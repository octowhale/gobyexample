package dbutils

import (
	"fmt"
	"log"
)

func dbExec() {
	sqlStmt := `insert into user(name,age) values("zhangshan",10);`

	ret, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(ret.LastInsertId(), ret.RowsAffected())
	id, _ := ret.LastInsertId()
	n, _ := ret.RowsAffected()
	fmt.Println(id, n)
}
