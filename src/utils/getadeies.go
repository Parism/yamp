package utils

import (
	"datastorage"
	"models"
)

/*
GetAdeies function
returns on a channel the adeia objects of the requested person
*/
func GetAdeies(id int, c chan []models.Adeia) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT * from personadeies where idperson=?", id)
	var adeia models.Adeia
	var adeies []models.Adeia
	for res.Next() {
		_ = res.Scan(
			&adeia.ID,
			&adeia.Start,
			&adeia.End,
			&adeia.Typos,
			&adeia.Days,
			&adeia.IDPerson,
		)
		adeia.BuildRepr()
		adeies = append(adeies, adeia)
	}
	c <- adeies
}
