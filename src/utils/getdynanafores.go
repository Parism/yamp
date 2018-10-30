package utils

import (
	"bytes"
	"datastorage"
	"models"
)

/*
GetDynAnaforesAll function
fetches all the different objects in the database
of the table anafores for a certain date
*/
func GetDynAnaforesAll(d string, c chan []models.Anafora) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT idperson,date,name,surname,perigrafi FROM anafores ")
	buffer.WriteString("JOIN proswpiko ON idperson = proswpiko.id ")
	buffer.WriteString("JOIN ranks ON proswpiko.rank = ranks.id ")
	buffer.WriteString("WHERE date = ? ORDER BY ranks.id DESC")
	res, _ := dbc.Query(buffer.String(), d)
	var anafora models.Anafora
	var anafores []models.Anafora
	for res.Next() {
		_ = res.Scan(
			&anafora.IDPerson,
			&anafora.Date,
			&anafora.Name,
			&anafora.Surname,
			&anafora.Perigrafi,
		)
		anafora.Date = models.DateBuilder(anafora.Date)
		anafores = append(anafores, anafora)
	}
	res.Close()
	c <- anafores
}

/*
GetDynAnaforesLabel function
returns all anafores objects for a certain label
*/
func GetDynAnaforesLabel(d string, label int, c chan []models.Anafora) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT idperson,date,name,surname,anafores.perigrafi FROM anafores ")
	buffer.WriteString("JOIN proswpiko ON idperson = proswpiko.id ")
	buffer.WriteString("JOIN ranks ON proswpiko.rank = ranks.id ")
	buffer.WriteString("JOIN ierarxia ON proswpiko.label = ierarxia.id ")
	buffer.WriteString("WHERE date = ? and (ierarxia.id =? || ierarxia.parentid = ?) ORDER BY ranks.id DESC")
	res, _ := dbc.Query(buffer.String(), d, label, label)
	var anafora models.Anafora
	var anafores []models.Anafora
	for res.Next() {
		_ = res.Scan(
			&anafora.IDPerson,
			&anafora.Date,
			&anafora.Name,
			&anafora.Surname,
			&anafora.Perigrafi,
		)
		anafora.Date = models.DateBuilder(anafora.Date)
		anafores = append(anafores, anafora)
	}
	res.Close()
	c <- anafores
}
