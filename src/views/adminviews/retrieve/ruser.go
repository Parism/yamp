package retrieve

import (
	"datastorage"
	"fmt"
	"log"
	"messages"
	"middleware"
	"models"
	"net/http"
	"strconv"
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
	cierarxia := make(chan []models.Groupld)
	id := r.URL.Query().Get("id")
	idint, _ := strconv.Atoi(id)
	go getIerarxia(cierarxia)
	ierarxia := <-cierarxia
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT id,username,roles.rolestring from accounts join roles on accounts.role=roles.role where id = ?;", idint)
	if err != nil {
		messages.SetMessage(r, "Invalid query")
		http.Redirect(w, r, "/listusers", http.StatusMovedPermanently)
	}
	var user models.User
	if res.Next() {
		err = res.Scan(
			&user.ID,
			&user.Username,
			&user.Role,
		)
		if err != nil {
			log.Println(err)
			return
		}
	}
	res.Close()
	datamap := make(map[string]interface{})
	datamap["ierarxia"] = ierarxia
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

func getIerarxia(c chan []models.Groupld) {
	ldArray := []models.Groupld{}
	ld := models.Groupld{}
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	query := "SELECT id,perigrafi FROM ierarxia;"
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
