// db/database.go

// package db

// import (
// 	"database/sql"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var DB *sql.DB

// func InitDB() {
// 	dsn := "root:Mysql@123@tcp(127.0.0.1:3306)/aashirwad"

// 	var err error
// 	DB, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}

//		err = DB.Ping()
//		if err != nil {
//			log.Fatalf("Failed to ping database: %v", err)
//		}
//	}
package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Mysql@123@tcp(127.0.0.1:3306)/aashirwad")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
