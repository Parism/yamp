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
	res, _ := dbc.Query("select typoiadeiwn.id,name,categories_adeiwn.category from typoiadeiwn join categories_adeiwn on typoiadeiwn.category = categories_adeiwn.id;")
	var typos models.TyposAdeias
	var typoiadeiwn []models.TyposAdeias
	for res.Next() {
		_ = res.Scan(
			&typos.ID,
			&typos.TyposAdeias,
			&typos.Category,
		)
		typoiadeiwn = append(typoiadeiwn, typos)
	}
	res.Close()
	c <- typoiadeiwn
}
