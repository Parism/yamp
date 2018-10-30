package utils

import (
	"datastorage"
	"log"
	"models"
)

/*
GetPersonAitiseis function
retrieves aitiseis objects
of certain person
*/
func GetPersonAitiseis(idint int, c chan []models.Aitisi) {
	var aitisi models.Aitisi
	var aitiseis []models.Aitisi
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT * FROM aitiseis where idperson = ? ORDER BY id DESC LIMIT 5", idint)
	if err != nil {
		log.Println(err)
		c <- nil
	}
	for res.Next() {
		_ = res.Scan(
			&aitisi.ID,
			&aitisi.Perigrafi,
			&aitisi.IDPerson,
			&aitisi.Date,
		)
		aitisi.Date = models.DateBuilder(aitisi.Date)
		aitiseis = append(aitiseis, aitisi)
	}
	res.Close()
	c <- aitiseis
}
