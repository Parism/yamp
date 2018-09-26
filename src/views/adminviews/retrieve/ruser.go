package retrieve

import (
	"datastorage"
	"fmt"
	"messages"
	"middleware"
	"models"
	"net/http"
	"utils"
	"variables"
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
	cdeltas := make(chan []models.Groupld)
	clambdas := make(chan []models.Groupld)
	cstring := make(chan string)
	id := r.URL.Query().Get("id")
	go getLd("lambdas", clambdas)
	go getLd("deltas", cdeltas)
	go getLabel(id, cstring)
	deltas := <-cdeltas
	lambdas := <-clambdas
	label := <-cstring
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT id,username,rolestring from accounts join roles on accounts.role=roles.role where id =?", id)
	if err != nil {
		messages.SetMessage(r, "Invalid query")
		http.Redirect(w, r, "/listusers", http.StatusMovedPermanently)
	}
	var user models.User
	if res.Next() {
		_ = res.Scan(
			&user.ID,
			&user.Username,
			&user.Role,
		)
	}
	res.Close()
	user.Label = label
	datamap := make(map[string]interface{})
	datamap["deltas"] = deltas
	datamap["lambdas"] = lambdas
	datamap["user"] = user
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = datamap
	t, err := utils.LoadTemplates("ruser",
		"templates/adminviews/ruser.html",
		"templates/adminviews/navbar.html",
		"templates/adminviews/header.html",
		"templates/adminviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	t.ExecuteTemplate(w, "ruser", data)
}

func getLd(q string, c chan []models.Groupld) {
	ldArray := []models.Groupld{}
	ld := models.Groupld{}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	query := "SELECT * FROM " + q + ";"
	res, _ := dbc.Query(query)
	for res.Next() {
		_ = res.Scan(
			&ld.ID,
			&ld.Name,
		)
		ldArray = append(ldArray, ld)
	}
	res.Close()
	c <- ldArray
}

func getLabel(id string, c chan string) {
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	var role int
	res, _ := dbc.Query("SELECT role from accounts where id = ?", id)
	if res.Next() {
		res.Scan(&role)
	}
	res.Close()
	var q string
	if role == variables.CAPTAIN {
		q = "select name from lambdaname where id = ?"
	} else if role == variables.USER {
		q = "select name from deltaname where id = ?"
	}
	res, _ = dbc.Query(q, id)
	var label string
	if res.Next() {
		_ = res.Scan(&label)
	}
	c <- label
}
