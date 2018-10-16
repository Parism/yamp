package views

import (
	"fmt"
	"middleware"
	"models"
	"net/http"
	"strconv"
	"utils"
	"variables"
)

func init() {
	GetMux().HandleFunc("/dynamologia", middleware.WithMiddleware(dynamologia,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

func dynamologia(w http.ResponseWriter, r *http.Request) {
	clabels := make(chan []models.Ierarxia)
	role := utils.GetSessionValue(r, "role")
	roleint, _ := strconv.Atoi(role)
	labelredis, _ := strconv.Atoi(utils.GetSessionValue(r, "label"))
	if roleint == variables.ADMIN {
		go utils.GetIerarxia(clabels)
	} else {
		go utils.GetLabels(labelredis, clabels)
	}
	var navbar string
	if roleint >= variables.ADMIN {
		navbar = "./templates/adminviews/navbar.html"
	} else {
		navbar = "./templates/userviews/navbar.html"
	}
	t, err := utils.LoadTemplates("dynamologia",
		"./templates/adminviews/dynamologia.html",
		"./templates/adminviews/header.html",
		"./templates/adminviews/footer.html",
		navbar,
	)
	if err != nil {
		fmt.Fprintf(w, "Error -> %s", err)
		return
	}
	datamap := make(map[string]interface{})
	data := utils.Data{}
	datamap["labels"] = <-clabels
	data.Data = datamap
	data.Context = utils.LoadContext(r)
	t.ExecuteTemplate(w, "dynamologia", data)
}
