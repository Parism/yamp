package models

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
	GetStart func() string
	GetEnd   func() string
}

func (ad *Adeia) GetStr() string {

}

/*
GetStartDate function
returns the start date
Test feature
*/
func (ad *Adeia) GetStartDate() string {
	return "test start date"
}

/*
GetEndDate function
returns the end date
Test feature
*/
func (ad *Adeia) GetEndDate() string {
	return "test end date"
}
