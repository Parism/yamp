package userviews

import (
	"html/template"
	"middleware"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/dashboard", middleware.WithMiddleware(Dashboard,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

/*
Dashboard function of package userviews
provides the main page of the user when logged in
*/
func Dashboard(w http.ResponseWriter, r *http.Request) {
	context := utils.LoadContext(r)
	data := utils.Data{}
	data.Context = context
	t := template.Must(template.ParseFiles("./templates/userviews/dashboard.html"))
	t.Execute(w, data)
}
