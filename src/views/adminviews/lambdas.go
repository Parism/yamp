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
	views.GetMux().HandleFunc("/lambdas", middleware.WithMiddleware(listlambdas,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func listlambdas(w http.ResponseWriter, r *http.Request) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT id,name from lambdas;")
	if err != nil {
		utils.RedirectWithError(w, r, "/diaxeiristiko", "Σφάλμα ανάγνωσης lambdas", err)
		return
	}
	var lambda models.Groupld
	var lambdas []models.Groupld
	for res.Next() {
		err = res.Scan(
			&lambda.ID,
			&lambda.Name)
		if err != nil {
			utils.RedirectWithError(w, r, "/lambdas", "Σφάλμα ανάγνωσης lambdas", err)
			return
		}
		lambdas = append(lambdas, lambda)
	}
	res.Close()
	t := template.Must(template.ParseFiles("templates/adminviews/lambdas.html"))
	var data utils.Data
	data.Context = utils.LoadContext(r)
	data.Data = lambdas
	t.Execute(w, data)
}
