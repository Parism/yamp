package models

/*
MinDynRecord struct
holds the records required
for the minimum dynamologio object
*/
type MinDynRecord struct {
	Rank     string `json:"rank"`
	Metaboli string `json:"metaboli"`
	Count    int    `json:"count"`
}
