package utils

import (
	"bytes"
	"datastorage"
	"log"
	"models"
)

/*
GetDynProswpikoAll function
fetches all the different objects in the database
of the table typoiadeiwn
*/
func GetDynProswpikoAll(d string, c chan []models.Proswpiko) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT proswpiko.id, proswpiko.name,proswpiko.surname,ranks.rank ")
	buffer.WriteString("from proswpiko ")
	buffer.WriteString("JOIN ranks on proswpiko.rank = ranks.id ")
	buffer.WriteString("WHERE proswpiko.id not in ")
	buffer.WriteString("(SELECT idperson from adeies ")
	buffer.WriteString("WHERE ? between adeies.start and adeies.end) ORDER BY ranks.id DESC")
	res, _ := dbc.Query(buffer.String(), d)
	var proswpiko models.Proswpiko
	var proswpikoArray []models.Proswpiko
	for res.Next() {
		_ = res.Scan(
			&proswpiko.ID,
			&proswpiko.Name,
			&proswpiko.Surname,
			&proswpiko.Rank,
		)
		proswpikoArray = append(proswpikoArray, proswpiko)
	}
	res.Close()
	c <- proswpikoArray
}

/*
GetDynProswpikoLabel function
returns asynchronously all objects in proswpiko table
that are not connected with an adeia object on a certain date
*/
func GetDynProswpikoLabel(d string, label int, c chan []models.Proswpiko) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT proswpiko.id, proswpiko.name,proswpiko.surname,ranks.rank ")
	buffer.WriteString("from proswpiko ")
	buffer.WriteString("JOIN ranks on proswpiko.rank = ranks.id ")
	buffer.WriteString("JOIN ierarxia on proswpiko.label = ierarxia.id ")
	buffer.WriteString("WHERE proswpiko.id not in ")
	buffer.WriteString("(SELECT idperson from adeies ")
	buffer.WriteString("WHERE ? between adeies.start and adeies.end) and ")
	buffer.WriteString("(ierarxia.id = ? || ierarxia.parentid = ?) ")
	buffer.WriteString("ORDER BY ranks.id DESC")
	res, err := dbc.Query(buffer.String(), d, label, label)
	if err != nil {
		log.Println(err)
		c <- nil
		return
	}
	var proswpiko models.Proswpiko
	var proswpikoArray []models.Proswpiko
	for res.Next() {
		_ = res.Scan(
			&proswpiko.ID,
			&proswpiko.Name,
			&proswpiko.Surname,
			&proswpiko.Rank,
		)
		proswpikoArray = append(proswpikoArray, proswpiko)
	}
	res.Close()
	c <- proswpikoArray
}
