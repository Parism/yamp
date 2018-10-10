package create

import (
	"datastorage"
	"fmt"
	"middleware"
	"net/http"
	"strconv"
	"strings"
	"time"
	"utils"
	"views"
)

func init() {
	views.GetMux().HandleFunc("/cadeia", middleware.WithMiddleware(cadeia,
		middleware.Time(),
		middleware.CsrfProtection(),
		middleware.NeedsSession(),
		middleware.IsUser(),
	))
}

func cadeia(w http.ResponseWriter, r *http.Request) {
	personid := r.PostFormValue("personid")
	idint, _ := strconv.Atoi(personid)
	if !utils.CanActOnPerson(r, idint) {
		http.Redirect(w, r, "/notfound", http.StatusMovedPermanently)
		return
	}
	typosadeias := r.PostFormValue("typosadeias")
	start := r.PostFormValue("start")
	end := r.PostFormValue("end")
	if end == "" {
		end = start
	}
	start = strings.TrimSpace(start)
	end = strings.TrimSpace(end)
	startdate, err := time.Parse("02/01/2006", start)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Σφάλμα ανάκτησης ημερομηνίας αρχής", err)
		return
	}
	enddate, err := time.Parse("02/01/2006", end)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Σφάλμα ανάκτησης ημερομηνίας τέλους", err)
		return
	}
	if startdate.After(enddate) {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Η αρχική ημερομηνία δεν μπορεί να είναι μεγαλύτερη από την τελική ημερομηνία", err)
		return
	}
	startToStore := fmt.Sprintf("%d/%d/%d",
		startdate.Year(),
		startdate.Month(),
		startdate.Day(),
	)
	endToStore := fmt.Sprintf("%d/%d/%d",
		enddate.Year(),
		enddate.Month(),
		enddate.Day(),
	)
	db, _ := datastorage.GetDataRouter().GetDb("common")
	dbc := db.GetMysqlClient()
	res, err := dbc.Query("SELECT COUNT(*) from adeies where (end between ? and ? or start between ? and ?) and idperson = ?", startToStore, endToStore, startToStore, endToStore, personid)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Σφάλμα ανάκτησης συνόλου μεταβολών", err)
		return
	}
	var count int
	if res.Next() {
		_ = res.Scan(&count)
	}
	if count > 0 {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Υπάρχει ήδη μεταβολή στις συγκεκριμένες ημερομηνίες", nil)
		return
	}
	stmt := datastorage.GetDataRouter().GetStmt("create_adeia")
	_, err = stmt.Exec(typosadeias, personid, startToStore, endToStore)
	if err != nil {
		utils.RedirectWithError(w, r, "/retrieveproswpiko?id="+personid, "Σφάλμα προσθήκης μεταβολής", err)
		return
	}
	http.Redirect(w, r, "/retrieveproswpiko?id="+personid, http.StatusMovedPermanently)
}
