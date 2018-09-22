package adminviews

import (
	"datastorage"
	"html/template"
	"log"
	"messages"
	"middleware"
	"models"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/typoiadeiwn", middleware.WithMiddleware(listtypoiadeiwn,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func listtypoiadeiwn(w http.ResponseWriter, r *http.Request) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT id,name FROM typoiadeiwn;")
	if err != nil {
		messages.SetMessage(r, "Σφάλμα κατά την φόρτωση των αδειών")
		log.Println(err)
	}
	var adeies []models.TyposAdeias
	var adeia models.TyposAdeias
	for res.Next() {
		err = res.Scan(
			&adeia.ID,
			&adeia.TyposAdeias,
		)
		if err != nil {
			log.Println(err, "error parsing typoi adeiwn")
			http.Redirect(w, r, "/typoiadeiwn", http.StatusMovedPermanently)
			return
		}
		adeies = append(adeies, adeia)
	}
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = adeies
	t := template.Must(template.ParseFiles("templates/adminviews/typoiadeiwn.html"))
	t.Execute(w, data)
}
