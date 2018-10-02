package utils

import (
	"datastorage"
	"models"
)

/*
GetTypoiAdeiwn function
fetches all the different objects in the database
of the table typoiadeiwn
*/
func GetTypoiAdeiwn(c chan []models.TyposAdeias) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("select id,name from typoiadeiwn")
	var typos models.TyposAdeias
	var typoiadeiwn []models.TyposAdeias
	for res.Next() {
		_ = res.Scan(
			&typos.ID,
			&typos.TyposAdeias,
		)
		typoiadeiwn = append(typoiadeiwn, typos)
	}
	c <- typoiadeiwn
}
