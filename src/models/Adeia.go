package models

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

/*
Adeia struct
holds the data for adeia objects retrieved from db
*/
type Adeia struct {
	ID       int
	Start    string
	End      string
	Typos    string
	Days     int
	IDPerson int
	Repr     string
}

/*
BuildRepr function of object Adeia
builds the Repr string to be displayed in the templates
*/
func (ad *Adeia) BuildRepr() {
	ad.Start = stringBuilder(ad.Start)
	ad.End = stringBuilder(ad.End)
	log.Println(ad.Start, ad.End)
	if ad.Days == 1 {
		ad.Repr = fmt.Sprintf("Σε %s 1 ημέρας, την %s", ad.Typos, ad.Start)
	} else {
		ad.Repr = fmt.Sprintf("Σε %s %d ημερών, από %s έως %s", ad.Typos, ad.Days, ad.Start, ad.End)
	}
}

func stringBuilder(s string) string {
	sarray := strings.Split(s, "-")
	var buffer bytes.Buffer
	for i := len(sarray) - 1; i >= 0; i-- {
		buffer.WriteString(sarray[i])
		if i > 0 {
			buffer.WriteString("-")
		}
	}
	return buffer.String()
}
