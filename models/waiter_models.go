// models/customer.go

package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Orderitem struct {
	Foodid    string
	Foodname  string
	Foodprice int
	// Foodqty   int
}
type Item struct {
	Name string `json:"name"`
}

// Getorderitem fetches all users from the database
func Getorderitem(db *sql.DB) ([]Orderitem, error) {
	rows, err := db.Query("SELECT foodid,foodname,foodprice FROM orderitem")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var orderitems []Orderitem
	for rows.Next() {
		var orderitem Orderitem
		err := rows.Scan(&orderitem.Foodid, &orderitem.Foodname, &orderitem.Foodprice)
		if err != nil {
			log.Println(err)
			continue
		}
		orderitems = append(orderitems, orderitem)
	}
	return orderitems, nil
}
func insertvalue(db *sql.DB, itemName string) error {
	result, err := db.Exec("INSERT INTO item (name) VALUES (?)", itemName)

	if err != nil {
		fmt.Println("Err", err.Error())
	}
	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	return RowsAffected, nill
}
