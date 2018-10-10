package retrieve

import (
	"datastorage"
	"fmt"
	"messages"
	"middleware"
	"models"
	"net/http"
	"strconv"
	"utils"
	"variables"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/retrieveproswpiko", middleware.WithMiddleware(rproswpiko,
		middleware.Time(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

/*
rproswpiko function
retrieves a single personel object
*/
func rproswpiko(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idint, _ := strconv.Atoi(id)
	if !utils.CanActOnPerson(r, idint) {
		http.Redirect(w, r, "/notfound", http.StatusMovedPermanently)
		return
	}
	cypiresies := make(chan []models.Ypiresia)
	cierarxia := make(chan []models.Ierarxia)
	ctypoiadeiwn := make(chan []models.TyposAdeias)
	ctypoiypiresiwn := make(chan []models.TyposYpiresias)
	cadeies := make(chan []models.Adeia)
	go utils.GetTypoiAdeiwn(ctypoiadeiwn)
	go utils.GetDimoiries(cierarxia)
	go utils.GetAdeies(idint, cadeies)
	go utils.GetPersonYpiresies(idint, cypiresies)
	label, _ := strconv.Atoi(utils.GetSessionValue(r, "label"))
	go utils.GetTypoiYpiresiwn(label, ctypoiypiresiwn)
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
			&proswpiko.Label,
		)
		if err != nil {
			utils.RedirectWithError(w, r, "/proswpiko", "Ανεπιτυχής ανάκτηση προσωπικού", err)
		}
	}
	res.Close()
	typoiadeiwn := <-ctypoiadeiwn
	ierarxia := <-cierarxia
	adeies := <-cadeies
	typoiypiresiwn := <-ctypoiypiresiwn
	datamap := make(map[string]interface{})
	datamap["Ierarxia"] = ierarxia
	datamap["Proswpiko"] = proswpiko
	datamap["TypoiAdeiwn"] = typoiadeiwn
	datamap["Adeies"] = adeies
	datamap["TypoiYpiresiwn"] = typoiypiresiwn
	datamap["Ypiresies"] = <-cypiresies
	data := utils.Data{}
	data.Context = utils.LoadContext(r)
	data.Data = datamap
	role := utils.GetSessionValue(r, "role")
	roleint, _ := strconv.Atoi(role)
	var navbar string
	if roleint >= variables.ADMIN {
		navbar = "./templates/adminviews/navbar.html"
	} else {
		navbar = "./templates/userviews/navbar.html"
	}
	t, err := utils.LoadTemplates("rproswpiko",
		"./templates/adminviews/rproswpiko.html",
		"./templates/adminviews/header.html",
		"./templates/adminviews/footer.html",
		navbar,
	)
	if err != nil {
		fmt.Fprintf(w, "Error -> %s", err)
		return
	}
	t.ExecuteTemplate(w, "rproswpiko", data)
}
