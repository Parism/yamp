package adminviews

import (
	"datastorage"
	"html/template"
	"middleware"
	"models"
	"net/http"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/listusers", middleware.WithMiddleware(listusers,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsAdmin(),
	))
}

func listusers(w http.ResponseWriter, r *http.Request) {
	context := utils.LoadContext(r)
	db, _ := datastorage.GetDataRouter().GetDb("common")
	mysqlclient := db.GetMysqlClient()
	res, err := mysqlclient.Query("SELECT id,username,rolestring from accounts join roles on accounts.role=roles.role order by roles.role desc, accounts.username asc")
	if err != nil {
		context.Message = err.Error()
		http.Redirect(w, r, "/listusers", http.StatusMovedPermanently)
	}
	var users []models.User
	var user models.User
	for res.Next() {
		err = res.Scan(
			&user.ID,
			&user.Username,
			&user.Role)
		users = append(users, user)
	}
	res.Close()
	data := utils.Data{}
	data.Context = context
	data.Data = users
	t := template.Must(template.ParseFiles("./templates/adminviews/users.html"))
	t.Execute(w, data)
}
