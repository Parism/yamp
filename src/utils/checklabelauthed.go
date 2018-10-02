package utils

import (
	"datastorage"
	"log"
	"variables"
)

/*
CheckLabelAuthed function
checks if a label and its associated request
are valid
*/
func CheckLabelAuthed(labelform, labelredis int) bool {
	if labelform >= variables.ADMIN {
		return true
	}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT id from ierarxia where id=? || parentid=?", labelredis, labelredis)
	var temp int
	for res.Next() {
		_ = res.Scan(&temp)
		log.Println("Checking ", labelform, "against ", temp)
		if labelform == temp {
			return true
		}
	}
	return false
}
