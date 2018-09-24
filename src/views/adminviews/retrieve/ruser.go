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
	views.GetMux().HandleFunc("/retrieveuser", middleware.WithMiddleware(ruser,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func ruser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT username,rolestring from accounts join roles on accounts.role=roles.role where username =?", username)
	if err != nil {
		messages.SetMessage(r, "Invalid query")
		http.Redirect(w, r, "/users", http.StatusMovedPermanently)
	}
	var user models.User
	if res.Next() {
		_ = res.Scan(
			&user.Username,
			&user.Role,
		)
	}
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = user
	t := template.Must(template.ParseFiles("./templates/adminviews/ruser.html"))
	t.Execute(w, data)
}
