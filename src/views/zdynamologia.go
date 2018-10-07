package views

import (
	"fmt"
	"middleware"
	"net/http"
	"utils"
)

func init() {
	GetMux().HandleFunc("/dynamologia", middleware.WithMiddleware(dynamologia,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

func dynamologia(w http.ResponseWriter, r *http.Request) {
	t, err := utils.LoadTemplates("dynamologia",
		"templates/adminviews/dynamologia.html",
		"templates/adminviews/navbar.html",
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	datamap := make(map[string]interface{})
	data := utils.Data{}
	data.Data = datamap
	data.Context = utils.LoadContext(r)
	t.ExecuteTemplate(w, "dynamologia", data)
}
