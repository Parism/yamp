package adminviews

import (
	"datastorage"
	"fmt"
	"log"
	"middleware"
	"models"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/listes", middleware.WithMiddleware(listes,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func listes(w http.ResponseWriter, r *http.Request) {
	list := []models.Proswpiko{}
	list = nil
	rankmap := models.RankMap{}
	rankmap.Init()
	cdeltas := make(chan []models.Groupld)
	clambdas := make(chan []models.Groupld)
	go utils.GetLd("deltas", cdeltas)
	go utils.GetLd("lambdas", clambdas)
	deltas := <-cdeltas
	lambdas := <-clambdas
	if r.Method == "POST" {
		lambda := r.PostFormValue("lambda")
		delta := r.PostFormValue("delta")
		list = getProswpikoList(lambda, delta)
		for _, value := range getRankList(lambda, delta) {
			rankmap.SetKey(value)
		}
		for index := range list {
			rankmap.Set(list[index].Rank, list[index])
		}
	}
	datamap := make(map[string]interface{})
	datamap["deltas"] = deltas
	datamap["lambdas"] = lambdas
	datamap["rankmap"] = rankmap
	data := utils.Data{}
	data.Data = datamap
	data.Context = utils.LoadContext(r)
	t, err := utils.LoadTemplates("listes",
		"templates/adminviews/listes.html",
		"templates/adminviews/navbar.html",
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	t.ExecuteTemplate(w, "listes", data)
}

func getProswpikoList(lambda, delta string) []models.Proswpiko {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, _ := dbc.Query("SELECT * FROM proswpiko_sorted where lname=? || dname=?", lambda, delta)
	var proswpikoArray []models.Proswpiko
	var proswpiko models.Proswpiko
	for res.Next() {
		_ = res.Scan(
			&proswpiko.ID,
			&proswpiko.Surname,
			&proswpiko.Name,
			&proswpiko.Rank,
			&proswpiko.Lambda,
			&proswpiko.Delta,
		)
		proswpikoArray = append(proswpikoArray, proswpiko)
	}
	return proswpikoArray
}

func getRankList(lambda, delta string) []string {
	var rankArray []string
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("select distinct proswpiko_sorted.rank,ranks.id from proswpiko_sorted join ranks on proswpiko_sorted.rank = ranks.rank where lname=? || dname=? ORDER BY ranks.id DESC", lambda, delta)
	if err != nil {
		log.Fatalln(err)
	}
	var rank string
	var rankid int
	for res.Next() {
		_ = res.Scan(
			&rank,
			&rankid,
		)
		rankArray = append(rankArray, rank)
	}
	return rankArray
}
