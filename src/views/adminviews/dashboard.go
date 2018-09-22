package adminviews

import (
	"html/template"
	"middleware"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/diaxeiristiko", middleware.WithMiddleware(Diaxeiristiko,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

/*
Diaxeiristiko function
serves as the function to the url /dashboardadmin
*/
func Diaxeiristiko(w http.ResponseWriter, r *http.Request) {
	context := utils.LoadContext(r)
	data := utils.Data{}
	data.Context = context
	t := template.Must(template.ParseFiles("./templates/adminviews/dashboard.html"))
	t.Execute(w, data)
}
