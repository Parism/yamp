package utils

import (
	"datastorage"
	"models"
)

/*
GetPersonYpiresies function
fetches all the different objects in the database
of the table typoiadeiwn
*/
func GetPersonYpiresies(id int, c chan []models.Ypiresia) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("select idypiresies,perigrafi,date from ypiresies join typoiypiresiwn on typeid = typoiypiresiwn.id where personid=?", id)
	var ypiresia models.Ypiresia
	var ypiresies []models.Ypiresia
	for res.Next() {
		_ = res.Scan(
			&ypiresia.ID,
			&ypiresia.Perigrafi,
			&ypiresia.Date,
		)
		ypiresia.Date = models.DateBuilder(ypiresia.Date)
		ypiresies = append(ypiresies, ypiresia)
	}
	res.Close()
	c <- ypiresies
}
