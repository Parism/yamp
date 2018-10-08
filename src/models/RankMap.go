package models

/*
RankMap struct
is a way to remember the insertion order in the map
iterate the map in insertion order by iterating the map
*/
type RankMap struct {
	Rankmap map[string][]Proswpiko `json:"rankmap"`
	keys    []string
}

/*
Init function
think of it as a builder. Startup function of the object
*/
func (rm *RankMap) Init() {
	rm.Rankmap = make(map[string][]Proswpiko)
}

/*
SetKey function
sets a key in the map and the key slice
*/
func (rm *RankMap) SetKey(key string) {
	rm.keys = append(rm.keys, key)
	rm.Rankmap[key] = nil
}

/*
GetKeys function
returns a list with the keys of the
RankMap in order to be able to iterate
the map in insertion order
*/
func (rm *RankMap) GetKeys() []string {
	return rm.keys
}

/*
Get function
returns the corresponding slice according to the key
*/
func (rm *RankMap) Get(key string) []Proswpiko {
	return rm.Rankmap[key]
}

/*
Set function
appends a value to the corresponding slice
*/
func (rm *RankMap) Set(key string, value Proswpiko) {
	rm.Rankmap[key] = append(rm.Rankmap[key], value)
}
