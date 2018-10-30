package models

import (
	"fmt"
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
	ad.Start = DateBuilder(ad.Start)
	ad.End = DateBuilder(ad.End)
	if ad.Days == 1 {
		ad.Repr = fmt.Sprintf("Σε %s 1 ημέρας, την %s", ad.Typos, ad.Start)
	} else {
		ad.Repr = fmt.Sprintf("Σε %s %d ημερών, από %s έως %s", ad.Typos, ad.Days, ad.Start, ad.End)
	}
}
