package utils

import (
	"datastorage"
	"log"
)

/*
CountAitiseis function
returns asynchronously the total amount of aitiseis objects in the db.nil
query will get reworked
*/
func CountAitiseis(c chan int, role string) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT COUNT(*) from aitiseis where id not in(select idaitisi from ypografes_aitisewn where signedas = ?);", role)
	defer res.Close()
	if err != nil {
		c <- -1
		log.Println(err, "Count aitiseis function")
		return
	}
	var num int
	if res.Next() {
		_ = res.Scan(&num)
	}
	c <- num
}
