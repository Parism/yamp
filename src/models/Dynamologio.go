package models

/*
Dynamologio struct
holds the final information
to be parsed into json
for this object
*/
type Dynamologio struct {
	Rankmap      CustomMap           `json:"rankmap"`
	Metaboles    []AdeiaDyn          `json:"metaboles"`
	MetabolesMin []MinDynAdeiaRecord `json:"metabolesmin"`
}
