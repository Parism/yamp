package views

import (
	"html/template"
	"middleware"
	"net/http"
	"utils"
)

func init() {
	GetMux().HandleFunc("/welcome", middleware.WithMiddleware(splitter,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
	))
}

func splitter(w http.ResponseWriter, r *http.Request) {
	context := utils.LoadContext(r)
	data := data{}
	data.Context = context
	t := template.Must(template.ParseFiles("./templates/splitter.html"))
	w.Header().Set("Server", "diagoras v0.1")
	t.Execute(w, data)
}
