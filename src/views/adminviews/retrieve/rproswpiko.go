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
	views.GetMux().HandleFunc("/retrieveproswpiko", middleware.WithMiddleware(rproswpiko,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

/*
rproswpiko function
retrieves a single personel object
*/
func rproswpiko(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("select * from rproswpiko where id=?", id)
	if err != nil {
		messages.SetMessage(r, "Invalid query")
		http.Redirect(w, r, "/proswpiko", http.StatusMovedPermanently)
	}
	var proswpiko models.Proswpiko
	if res.Next() {
		err = res.Scan(
			&proswpiko.ID,
			&proswpiko.Name,
			&proswpiko.Surname,
			&proswpiko.Rank,
			&proswpiko.Lambda,
			&proswpiko.Delta,
		)
		if err != nil {
			utils.RedirectWithError(w, r, "/proswpiko", "Ανεπιτυχής ανάκτηση προσωπικού", err)
		}
	}
	res.Close()
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = proswpiko
	t, err := utils.LoadTemplates("rproswpiko",
		"./templates/adminviews/rproswpiko.html",
		"./templates/adminviews/header.html",
		"./templates/adminviews/footer.html",
		"./templates/adminviews/navbar.html",
	)
	if err != nil {
		fmt.Fprintf(w, "Error -> %s", err)
		return
	}
	t.ExecuteTemplate(w, "rproswpiko", data)
}
