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
	views.GetMux().HandleFunc("/ierarxia", middleware.WithMiddleware(ierarxia,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func ierarxia(w http.ResponseWriter, r *http.Request) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT id,perigrafi FROM ierarxia;")
	if err != nil {
		utils.RedirectWithError(w, r, "/diaxeiristiko", "Σφάλμα κατά την φόρτωση της ιεραρχίας", err)
		return
	}
	var ierarxia models.Ierarxia
	var ierarxiaArray []models.Ierarxia
	for res.Next() {
		err = res.Scan(
			&ierarxia.ID,
			&ierarxia.Perigrafi,
		)
		if err != nil {
			utils.RedirectWithError(w, r, "/diaxeiristiko", "Σφάλμα κατά την φόρτωση της ιεραρχίας", err)
			return
		}
		ierarxiaArray = append(ierarxiaArray, ierarxia)
	}
	res.Close()
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = ierarxiaArray
	t, err := utils.LoadTemplates("ierarxia",
		"templates/adminviews/ierarxia.html",
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html",
		"templates/adminviews/navbar.html")
	if err != nil {
		fmt.Fprintf(w, "Error->%s", err)
		return
	}
	t.ExecuteTemplate(w, "ierarxia", data)
}
