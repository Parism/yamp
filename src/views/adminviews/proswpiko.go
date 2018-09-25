package adminviews

import (
	"datastorage"
	"html/template"
	"middleware"
	"models"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/proswpiko", middleware.WithMiddleware(proswpiko,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func proswpiko(w http.ResponseWriter, r *http.Request) {
	cproswpiko := make(chan []models.Proswpiko)
	cranks := make(chan []models.Rank)
	cdeltas := make(chan []models.Groupld)
	clambdas := make(chan []models.Groupld)
	go getProswpiko(cproswpiko)
	go getRanks(cranks)
	go getLd("lambdas", clambdas)
	go getLd("deltas", cdeltas)
	proswpikolist := <-cproswpiko
	ranks := <-cranks
	deltas := <-cdeltas
	lambdas := <-clambdas
	datamap := make(map[string]interface{})
	datamap["proswpiko"] = proswpikolist
	datamap["ranks"] = ranks
	datamap["deltas"] = deltas
	datamap["lambdas"] = lambdas
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = datamap
	t := template.Must(template.ParseFiles("templates/adminviews/proswpiko.html"))
	t.Execute(w, data)

}

func getProswpiko(c chan []models.Proswpiko) {
	proswpikoArray := []models.Proswpiko{}
	proswpiko := models.Proswpiko{}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT id,pname,surname,rank from proswpiko_sorted;")
	for res.Next() {
		_ = res.Scan(
			&proswpiko.ID,
			&proswpiko.Name,
			&proswpiko.Surname,
			&proswpiko.Rank,
		)
		proswpikoArray = append(proswpikoArray, proswpiko)
	}
	c <- proswpikoArray
}

func getRanks(c chan []models.Rank) {
	ranksArray := []models.Rank{}
	rank := models.Rank{}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT * from ranks;")
	for res.Next() {
		_ = res.Scan(
			&rank.ID,
			&rank.Rank,
		)
		ranksArray = append(ranksArray, rank)
	}
	c <- ranksArray
}

func getLd(q string, c chan []models.Groupld) {
	ldArray := []models.Groupld{}
	ld := models.Groupld{}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	query := "SELECT * FROM " + q + ";"
	res, _ := dbc.Query(query)
	for res.Next() {
		_ = res.Scan(
			&ld.ID,
			&ld.Name,
		)
		ldArray = append(ldArray, ld)
	}
	c <- ldArray
}
