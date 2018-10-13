package utils

import (
	"bytes"
	"datastorage"
	"log"
	"models"
)

/*
GetDynAdeies function
Fetches ALL the info about dynadeies
*/
func GetDynAdeies(d string, c chan []models.AdeiaDyn) {
	var metaboles []models.AdeiaDyn
	var metaboli models.AdeiaDyn
	var buffer bytes.Buffer
	buffer.WriteString("select personid,pname,sname,start,end,typoiadeiwn.name,")
	buffer.WriteString("days,monada,rank,categories_adeiwn.category ")
	buffer.WriteString("FROM adeiesdynamologiou ")
	buffer.WriteString("join typoiadeiwn on adeiesdynamologiou.name = typoiadeiwn.name ")
	buffer.WriteString("join categories_adeiwn on typoiadeiwn.category = categories_adeiwn.id ")
	buffer.WriteString("WHERE ? BETWEEN start and end")
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query(buffer.String(), d)
	if err != nil {
		log.Println("Error fetching labeled adeies dyn", err)
		c <- nil
		return
	}
	for res.Next() {
		_ = res.Scan(
			&metaboli.PersonID,
			&metaboli.Name,
			&metaboli.Surname,
			&metaboli.Start,
			&metaboli.End,
			&metaboli.Typos,
			&metaboli.Days,
			&metaboli.Monada,
			&metaboli.Rank,
			&metaboli.Category,
		)
		metaboli.BuildRepr()
		metaboles = append(metaboles, metaboli)
	}
	c <- metaboles
}

/*
GetDynAdeiesLabeled function
is an asynchronous function that fetches
info about persons and adeia objects aggregated
*/
func GetDynAdeiesLabeled(d string, label int, c chan []models.AdeiaDyn) {
	var metaboles []models.AdeiaDyn
	var metaboli models.AdeiaDyn
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("select personid,pname,sname,start,end,typoiadeiwn.name,")
	buffer.WriteString("days,monada,rank,categories_adeiwn.category ")
	buffer.WriteString("FROM adeiesdynamologiou ")
	buffer.WriteString("join typoiadeiwn on adeiesdynamologiou.name = typoiadeiwn.name ")
	buffer.WriteString("join categories_adeiwn on typoiadeiwn.category = categories_adeiwn.id ")
	buffer.WriteString("WHERE (? BETWEEN start and end) ")
	buffer.WriteString("and ")
	buffer.WriteString("(iid = ? || pid=?)")
	res, err := dbc.Query(buffer.String(), d, label, label)
	if err != nil {
		log.Println("Error fetching labeled adeies dyn", err)
		c <- nil
		return
	}
	for res.Next() {
		_ = res.Scan(
			&metaboli.PersonID,
			&metaboli.Name,
			&metaboli.Surname,
			&metaboli.Start,
			&metaboli.End,
			&metaboli.Typos,
			&metaboli.Days,
			&metaboli.Monada,
			&metaboli.Rank,
			&metaboli.Category,
		)
		metaboli.BuildRepr()
		metaboles = append(metaboles, metaboli)
	}
	res.Close()
	c <- metaboles
}
