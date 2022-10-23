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

// TO-CHECK:DEEB ORM

// import "crypto/sha256"
h := sha256.New()
io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
fmt.Printf("% x", h.Sum(nil))
// import "crypto/sha1"
h := sha1.New()
io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
fmt.Printf("% x", h.Sum(nil))
// import "crypto/md5"
h := md5.New()
io.WriteString(h, "xasd")
fmt.Printf("%x", h.Sum(nil))