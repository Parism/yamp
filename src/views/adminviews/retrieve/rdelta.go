package retrieve

import (
	"datastorage"
	"fmt"
	"messages"
	"middleware"
	"models"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/retrievedelta", middleware.WithMiddleware(rdelta,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func rdelta(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT id,name from deltas where id =?", id)
	if err != nil {
		messages.SetMessage(r, "Invalid query")
		http.Redirect(w, r, "/deltas", http.StatusMovedPermanently)
	}
	var delta models.Groupld
	if res.Next() {
		err = res.Scan(
			&delta.ID,
			&delta.Name)
		if err != nil {
			utils.RedirectWithError(w, r, "/deltas", "Ανεπιτυχης προσπέλαση delta", err)
			return
		}
	}
	res.Close()
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = delta
	t, err := utils.LoadTemplates("rdelta",
		"templates/adminviews/rdelta.html",
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html",
		"templates/adminviews/navbar.html")
	if err != nil {
		fmt.Fprintf(w, "->%s", err)
		return
	}
	t.ExecuteTemplate(w, "rdelta", data)
}
