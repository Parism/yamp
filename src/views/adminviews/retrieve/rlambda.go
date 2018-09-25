package retrieve

import (
	"datastorage"
	"html/template"
	"messages"
	"middleware"
	"models"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/retrievelambda", middleware.WithMiddleware(rlambda,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func rlambda(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT id,name from lambdas where id =?", id)
	if err != nil {
		messages.SetMessage(r, "Invalid query")
		http.Redirect(w, r, "/lambdas", http.StatusMovedPermanently)
	}
	var lambda models.Groupld
	if res.Next() {
		err = res.Scan(
			&lambda.ID,
			&lambda.Name)
		if err != nil {
			utils.RedirectWithError(w, r, "/lambdas", "Ανεπιτυχης προσπέλαση lambda", err)
			return
		}
	}
	res.Close()
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = lambda
	t := template.Must(template.ParseFiles("templates/adminviews/rlambda.html"))
	t.Execute(w, data)
}
