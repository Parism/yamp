package utils

import (
	"bytes"
	"datastorage"
	"log"
	"models"
)

/*
GetDynYpiresiesAll function
fetches all ypiresies on a date
*/
func GetDynYpiresiesAll(d string, c chan []models.Ypiresia) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT personid,name,surname,ranks.rank,typoiypiresiwn.perigrafi,date ")
	buffer.WriteString("from ypiresies ")
	buffer.WriteString("join proswpiko on proswpiko.id = ypiresies.personid ")
	buffer.WriteString("join ranks on proswpiko.rank = ranks.id ")
	buffer.WriteString("join typoiypiresiwn on typoiypiresiwn.id = ypiresies.typeid ")
	buffer.WriteString("where ypiresies.date = ? ")
	buffer.WriteString("order by ranks.id DESC, surname ASC")
	res, _ := dbc.Query(buffer.String(), d)
	var ypiresia models.Ypiresia
	var ypiresies []models.Ypiresia
	for res.Next() {
		_ = res.Scan(
			&ypiresia.PersonID,
			&ypiresia.Name,
			&ypiresia.Surname,
			&ypiresia.Rank,
			&ypiresia.Perigrafi,
			&ypiresia.Date,
		)
		ypiresia.Date = models.DateBuilder(ypiresia.Date)
		ypiresies = append(ypiresies, ypiresia)
	}
	res.Close()
	c <- ypiresies
}

/*
GetDynYpiresiesLabel function
returns all ypiresies for a certain label
*/
func GetDynYpiresiesLabel(d string, label int, c chan []models.Ypiresia) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT personid,name,surname,ranks.rank,typoiypiresiwn.perigrafi,date ")
	buffer.WriteString("from ypiresies ")
	buffer.WriteString("join proswpiko on proswpiko.id = ypiresies.personid ")
	buffer.WriteString("join ranks on proswpiko.rank = ranks.id ")
	buffer.WriteString("join typoiypiresiwn on typoiypiresiwn.id = ypiresies.typeid ")
	buffer.WriteString("join ierarxia on proswpiko.label = ierarxia.id ")
	buffer.WriteString("where ypiresies.date = ? ")
	buffer.WriteString("and (ierarxia.id=? || ierarxia.parentid=?) ")
	buffer.WriteString("order by ranks.id DESC, surname ASC")
	res, err := dbc.Query(buffer.String(), d, label, label)
	if err != nil {
		log.Println(err)
		c <- nil
		return
	}
	var ypiresia models.Ypiresia
	var ypiresies []models.Ypiresia
	for res.Next() {
		_ = res.Scan(
			&ypiresia.PersonID,
			&ypiresia.Name,
			&ypiresia.Surname,
			&ypiresia.Rank,
			&ypiresia.Perigrafi,
			&ypiresia.Date,
		)
		ypiresia.Date = models.DateBuilder(ypiresia.Date)
		ypiresies = append(ypiresies, ypiresia)
	}
	res.Close()
	c <- ypiresies
}
