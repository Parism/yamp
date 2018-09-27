package adminviews

import (
	"datastorage"
	"fmt"
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
	go utils.GetLd("lambdas", clambdas)
	go utils.GetLd("deltas", cdeltas)
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
	t, err := utils.LoadTemplates("proswpiko",
		"templates/adminviews/proswpiko.html",
		"templates/adminviews/navbar.html",
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	t.ExecuteTemplate(w, "proswpiko", data)
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
	res.Close()
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
	res.Close()
	c <- ranksArray
}
