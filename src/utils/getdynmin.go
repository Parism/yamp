package utils

import (
	"bytes"
	"datastorage"
	"log"
	"models"
)

/*
GetDynMinAll function
returns asynchronously
the minimal version of the dynamologio
*/
func GetDynMinAll(d string, c chan []models.MinDynRecord) {
	var record models.MinDynRecord
	var records []models.MinDynRecord
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("select ranks.rank,typoiadeiwn.name,COUNT(*) from adeies join proswpiko on adeies.idperson = proswpiko.id join ranks on proswpiko.rank = ranks.id join typoiadeiwn on adeies.type = typoiadeiwn.id where (? between adeies.start and adeies.end) GROUP BY type,rank", d)
	if err != nil {
		log.Println(err)
		c <- nil
		return
	}
	for res.Next() {
		_ = res.Scan(
			&record.Rank,
			&record.Metaboli,
			&record.Count,
		)
		records = append(records, record)
	}
	c <- records
}

/*
GetDynMinAdeiesAll function
returns asynchronously the sum of the categories in a certain date
*/
func GetDynMinAdeiesAll(d string, c chan []models.MinDynAdeiaRecord) {
	var record models.MinDynAdeiaRecord
	var records []models.MinDynAdeiaRecord
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("select ranks.rank,categories_adeiwn.category,COUNT(*) from adeies ")
	buffer.WriteString("join typoiadeiwn on adeies.type = typoiadeiwn.id ")
	buffer.WriteString("join categories_adeiwn on typoiadeiwn.category = categories_adeiwn.id ")
	buffer.WriteString("join proswpiko on proswpiko.id = adeies.idperson ")
	buffer.WriteString("join ierarxia on ierarxia.id = proswpiko.label ")
	buffer.WriteString("join ranks on proswpiko.rank = ranks.id ")
	buffer.WriteString("where (? between adeies.start and adeies.end) GROUP BY proswpiko.rank,categories_adeiwn.category ORDER BY category ASC, ranks.id DESC")
	res, err := dbc.Query(buffer.String(), d)
	if err != nil {
		log.Println(err)
		c <- nil
		return
	}
	for res.Next() {
		_ = res.Scan(
			&record.Rank,
			&record.Category,
			&record.Count,
		)
		records = append(records, record)
	}
	c <- records
}

/*
GetDynMinAdeiesLabel function
fetches min adeies data based on a label
*/
func GetDynMinAdeiesLabel(d string, label int, c chan []models.MinDynAdeiaRecord) {
	var record models.MinDynAdeiaRecord
	var records []models.MinDynAdeiaRecord
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("select ranks.rank,categories_adeiwn.category,COUNT(*) from adeies ")
	buffer.WriteString("join typoiadeiwn on adeies.type = typoiadeiwn.id ")
	buffer.WriteString("join categories_adeiwn on typoiadeiwn.category = categories_adeiwn.id ")
	buffer.WriteString("join proswpiko on proswpiko.id = adeies.idperson ")
	buffer.WriteString("join ierarxia on ierarxia.id = proswpiko.label ")
	buffer.WriteString("join ranks on proswpiko.rank = ranks.id ")
	buffer.WriteString("where (? between adeies.start and adeies.end) ")
	buffer.WriteString("and ")
	buffer.WriteString("(ierarxia.id=? || ierarxia.parentid=?) ")
	buffer.WriteString("GROUP BY proswpiko.rank,categories_adeiwn.category")
	res, err := dbc.Query(buffer.String(), d, label, label)
	if err != nil {
		log.Println(err)
		c <- nil
		return
	}
	for res.Next() {
		_ = res.Scan(
			&record.Rank,
			&record.Category,
			&record.Count,
		)
		records = append(records, record)
	}
	c <- records
}

/*
GetDynMinLabel function
returns asynchronously
the minimal version of the dynamologio
according to the label
*/
func GetDynMinLabel(d string, label int, c chan []models.MinDynRecord) {
	var record models.MinDynRecord
	var records []models.MinDynRecord
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var buffer bytes.Buffer
	buffer.WriteString("SELECT rank,typoiadeiwn.name,COUNT(*) ")
	buffer.WriteString("from adeies ")
	buffer.WriteString("join proswpiko_sorted on adeies.idperson = proswpiko_sorted.id ")
	buffer.WriteString("join typoiadeiwn on adeies.type = typoiadeiwn.id where iid = ? || pid=? ")
	buffer.WriteString("and ")
	buffer.WriteString("(? between adeies.start and adeies.end) ")
	buffer.WriteString("GROUP BY proswpiko_sorted.rank,typoiadeiwn.name")
	res, err := dbc.Query(buffer.String(), label, label, d)
	if err != nil {
		log.Println(err)
		c <- nil
		return
	}
	for res.Next() {
		_ = res.Scan(
			&record.Rank,
			&record.Metaboli,
			&record.Count,
		)
		records = append(records, record)
	}
	res.Close()
	c <- records
}
