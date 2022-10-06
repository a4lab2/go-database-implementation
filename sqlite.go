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
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)
	res, err := stmt.Exec("sam", "comp", "2023-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("samp", id)
	checkErr(err)
	affectedrows, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affectedrows)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time
	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	// close query ops
	rows.Close()

	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}
