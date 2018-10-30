package utils

import (
	"bytes"
	"datastorage"
	"models"
)

/*
GetPersonErgasies function
fetches all the different objects in the database
of the table typoiadeiwn
*/
func GetPersonErgasies(id int, c chan []models.Ergasia) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT ergasies.id, proswpiko.id, perigrafi, ")
	buffer.WriteString("date, name, surname from ergasies ")
	buffer.WriteString("JOIN proswpiko on proswpiko.id = ergasies.idperson ")
	buffer.WriteString("WHERE idperson = ?")
	res, _ := dbc.Query(buffer.String(), id)
	var ergasia models.Ergasia
	var ergasies []models.Ergasia
	for res.Next() {
		_ = res.Scan(
			&ergasia.ID,
			&ergasia.IDPerson,
			&ergasia.Perigrafi,
			&ergasia.Date,
			&ergasia.Name,
			&ergasia.Surname,
		)
		ergasia.Date = models.DateBuilder(ergasia.Date)
		ergasies = append(ergasies, ergasia)
	}
	res.Close()
	c <- ergasies
}
