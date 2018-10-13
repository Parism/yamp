package utils

import (
	"bytes"
	"datastorage"
	"models"
)

/*
GetPersonAnafores function
fetches all the different objects in the database
of the table typoiadeiwn
*/
func GetPersonAnafores(id int, c chan []models.Anafora) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT anafores.id, proswpiko.id, perigrafi, ")
	buffer.WriteString("date, name, surname from anafores ")
	buffer.WriteString("JOIN proswpiko on proswpiko.id = anafores.idperson ")
	buffer.WriteString("WHERE idperson = ?")
	res, _ := dbc.Query(buffer.String(), id)
	var anafora models.Anafora
	var anafores []models.Anafora
	for res.Next() {
		_ = res.Scan(
			&anafora.ID,
			&anafora.IDPerson,
			&anafora.Perigrafi,
			&anafora.Date,
			&anafora.Name,
			&anafora.Surname,
		)
		anafora.Date = models.DateBuilder(anafora.Date)
		anafores = append(anafores, anafora)
	}
	res.Close()
	c <- anafores
}
