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
	views.GetMux().HandleFunc("/deltas", middleware.WithMiddleware(listdeltas,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func listdeltas(w http.ResponseWriter, r *http.Request) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT id,name from deltas;")
	if err != nil {
		utils.RedirectWithError(w, r, "/deltas", "Σφάλμα ανάγνωσης deltas", err)
		return
	}
	var delta models.Groupld
	var deltas []models.Groupld
	for res.Next() {
		err = res.Scan(
			&delta.ID,
			&delta.Name)
		if err != nil {
			utils.RedirectWithError(w, r, "/diaxeiristiko", "Σφάλμα ανάγνωσης deltas", err)
			return
		}
		deltas = append(deltas, delta)
	}
	res.Close()
	t := template.Must(template.ParseFiles("templates/adminviews/deltas.html"))
	var data utils.Data
	data.Context = utils.LoadContext(r)
	data.Data = deltas
	t.Execute(w, data)
}