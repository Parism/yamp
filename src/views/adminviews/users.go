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
	views.GetMux().HandleFunc("/listusers", middleware.WithMiddleware(listusers,
		middleware.Time(),
		middleware.NoCache(),
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
	t, err := utils.LoadTemplates("users",
		"templates/adminviews/users.html",
		"templates/adminviews/navbar.html",
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	t.ExecuteTemplate(w, "users", data)
}
