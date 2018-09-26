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
	t, err := utils.LoadTemplates("lambdas",
		"templates/adminviews/lambdas.html",
		"templates/adminviews/navbar.html",
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	var data utils.Data
	data.Context = utils.LoadContext(r)
	data.Data = lambdas
	t.ExecuteTemplate(w, "lambdas", data)
}
