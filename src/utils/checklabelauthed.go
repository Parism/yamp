package utils

import (
	"datastorage"
	"net/http"
	"strconv"
	"variables"
)

/*
CheckLabelAuthed function
checks if a label and its associated request
are valid
*/
func CheckLabelAuthed(r *http.Request, label int) bool {
	role := GetSessionValue(r, "role")
	roleint, _ := strconv.Atoi(role)
	if roleint >= variables.ADMIN {
		return true
	}
	labeltemp := GetSessionValue(r, "label")
	labelredis, _ := strconv.Atoi(labeltemp)
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("select id from ierarxia where id = ? || parentid = ?;", labelredis, labelredis)
	var temp int
	for res.Next() {
		_ = res.Scan(&temp)
		if label == temp {
			res.Close()
			return true
		}
	}
	res.Close()
	return false

}
