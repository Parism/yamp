package views

import (
	"html/template"
	"middleware"
	"net/http"
	"utils"
)

func init() {
	GetMux().HandleFunc("/", middleware.WithMiddleware(index,
		middleware.Time(),
		middleware.NeedsSession(),
	))
}

type data struct {
	Context utils.Context
}

/*
index function
main page of diagoras.
provides the login form
*/
func index(w http.ResponseWriter, r *http.Request) {
	context := utils.LoadContext(r)
	data := data{}
	data.Context = context
	t := template.Must(template.ParseFiles("./templates/index.html"))
	t.Execute(w, data)
}
