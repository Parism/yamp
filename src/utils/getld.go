package utils

import (
	"datastorage"
	"models"
)

/*
GetLd function
is used to fetch either the lambda or the delta values from the db
Meant to be used asynchronously
*/
func GetLd(q string, c chan []models.Groupld) {
	ldArray := []models.Groupld{}
	ld := models.Groupld{}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	query := "SELECT * FROM " + q + ";"
	res, _ := dbc.Query(query)
	for res.Next() {
		_ = res.Scan(
			&ld.ID,
			&ld.Name,
		)
		ldArray = append(ldArray, ld)
	}
	res.Close()
	c <- ldArray
}
