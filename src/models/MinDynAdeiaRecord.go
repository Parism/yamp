package models

/*
MinDynAdeiaRecord struct
is a record holding the sum of a category
*/
type MinDynAdeiaRecord struct {
	Rank     string `json:"rank"`
	Category string `json:"category"`
	Count    int    `json:"count"`
}
