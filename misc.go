// package database

// import "fmt"

// func trnsactions() {
// 	stmt, err := database.Prepare("update task set is_deleted='Y',last_modified_at=datetime() where id=?")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	tx, err := database.Begin()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	_, err = tx.Stmt(trashSQL).Exec(id)
// 	if err != nil {
// 		fmt.Println("doing rollback")
// 		tx.Rollback()
// 	} else {
// 		tx.Commit()
// 	}
// }
