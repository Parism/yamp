package utils

import (
	"datastorage"
	"fmt"
	"log"
)

/*
CheckLabelAuthed function
checks if a label and its associated request
are valid
*/
func CheckLabelAuthed(labelform, labelredis int) bool {
	fmt.Printf("%d %d\n", labelform, labelredis)
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT id from ierarxia where id=? || parentid=?", labelredis, labelredis)
	var temp int
	for res.Next() {
		_ = res.Scan(&temp)
		log.Println("Checking", temp, "against", labelredis)
		if temp == labelform {
			log.Println("Equals")
			return true
		}
	}
	return false
}
