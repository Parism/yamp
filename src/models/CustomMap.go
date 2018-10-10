package models

/*
CustomMap struct
is a way to remember the insertion order in the map
iterate the map in insertion order by iterating the map
*/
type CustomMap struct {
	Map  map[string][]interface{} `json:"map"`
	Keys []string
}

/*
Init function
think of it as a builder. Startup function of the object
*/
func (rm *CustomMap) Init() {
	rm.Map = make(map[string][]interface{})
}

/*
SetKey function
sets a key in the map and the key slice
*/
func (rm *CustomMap) SetKey(key string) {
	rm.Keys = append(rm.Keys, key)
	rm.Map[key] = nil
}

/*
GetKeys function
returns a list with the keys of the
CustomMap in order to be able to iterate
the map in insertion order
*/
func (rm *CustomMap) GetKeys() []string {
	return rm.Keys
}

/*
Get function
returns the corresponding slice according to the key
*/
func (rm *CustomMap) Get(key string) []interface{} {
	return rm.Map[key]
}

/*
Set function
appends a value to the corresponding slice
*/
func (rm *CustomMap) Set(key string, value interface{}) {
	rm.Map[key] = append(rm.Map[key], value)
}
