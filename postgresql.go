package database

import (
	"database/sql"
	"fmt"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	dsn := ""
	db, err := sql.Open("postgres", dsn)
	checkErr(err)
	defer db.Close()
	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "jsampie", "biochem", "2019-11-09").Scan(&lastInsertId)
	checkErr(err)

	//Update
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)
	res, err := stmt.Exec("astaxieupdate", lastInsertId)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect, "rows changed")
	// Query
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid | username | department | created ")
		fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	}
	// Delete
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)
	res, err = stmt.Exec(lastInsertId)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect, "rows changed")
}
