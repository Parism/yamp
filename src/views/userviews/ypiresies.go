package userviews

import (
	"fmt"
	"middleware"
	"models"
	"net/http"
	"strconv"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/ypiresies", middleware.WithMiddleware(ypiresies,
		middleware.Time(),
		middleware.NoCache(),
		middleware.NeedsSession(),
		middleware.CsrfProtection(),
		middleware.IsUser(),
	))
}

func ypiresies(w http.ResponseWriter, r *http.Request) {
	clabels := make(chan []models.Ierarxia)
	ctypoiypiresiwn := make(chan []models.TyposYpiresias)
	labelredis, _ := strconv.Atoi(utils.GetSessionValue(r, "label"))
	go utils.GetLabels(labelredis, clabels)
	if r.Method == "POST" {
		label := r.PostFormValue("label")
		labelform, _ := strconv.Atoi(label)
		if utils.CheckLabelAuthed(r, labelform) {
			go utils.GetTypoiYpiresiwn(labelform, ctypoiypiresiwn)
		}
	}
	t, err := utils.LoadTemplates("ypiresies",
		"templates/userviews/ypiresies.html",
		"templates/userviews/navbar.html",
		"templates/userviews/header.html",
		"templates/userviews/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Err->%s", err)
		return
	}
	datamap := make(map[string]interface{})
	datamap["ypiresies"] = <-ctypoiypiresiwn
	datamap["labels"] = <-clabels
	data := utils.Data{}
	data.Data = datamap
	data.Context = utils.LoadContext(r)
	t.ExecuteTemplate(w, "ypiresies", data)
}
