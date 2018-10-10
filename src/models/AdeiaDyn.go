package models

import "fmt"

/*
AdeiaDyn struct
holds information about a single persons
leaves
Not using embedded structs
to be easier to parse when JSONified
*/
type AdeiaDyn struct {
	PersonID int
	Name     string
	Surname  string
	Start    string
	End      string
	Typos    string
	Days     int
	Monada   string
	Rank     string
	Repr     string
	Category string
}

/*
BuildRepr function
builds the string representation of the object
*/
func (ad *AdeiaDyn) BuildRepr() {
	ad.Start = DateBuilder(ad.Start)
	ad.End = DateBuilder(ad.End)
	if ad.Days == 1 {
		ad.Repr = fmt.Sprintf("σε %s 1 ημέρας, την %s", ad.Typos, ad.Start)
	} else {
		ad.Repr = fmt.Sprintf("σε %s %d ημερών, από %s έως %s", ad.Typos, ad.Days, ad.Start, ad.End)
	}
}
