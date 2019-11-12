package dbutils

import "fmt"

import "log"

type user struct {
	ID   int
	Name string
	Age  int
}

// QueryOne 返回单行数据
func QueryOne() {

	// 查询
	sqlStmt := "select id, name, age from user where id = ?;"
	row := db.QueryRow(sqlStmt, 1)

	// 取值
	var u user
	err := row.Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u.Name, u.ID, u.Age)

	//2
	db.QueryRow(sqlStmt, 2).Scan(&u.ID, &u.Name, &u.Age)
	fmt.Println(u.Name, u.ID, u.Age)

}

// QueryMore 返回多行数据
func QueryMore(n int) {
	sqlStmt := "select id, name, age from user where id > ?;"

	rows, err := db.Query(sqlStmt, n)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var u user
	for rows.Next() {
		rows.Scan(&u.ID, &u.Name, &u.Age)
		fmt.Println(u.ID, u.Name, u.Age)
	}
	fmt.Println("bye")
}
