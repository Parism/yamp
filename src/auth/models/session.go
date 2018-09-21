package models

import (
	"encoding/json"
)

/*
Session struct
holds a map containing everything
associated with a sessionid
Is converted to json and stored within redis
Can be retrieved by sessionid
{sessionid -> sessionstruct (json)}
*/
type Session struct {
	Sessionmap map[string]string `json:"session"`
}

/*
GetKey function
is a getter function for the session struct
*/
func (s *Session) GetKey(key string) string {
	if value, exists := s.Sessionmap[key]; exists {
		return value
	}
	return ""
}

/*
SetKey function
is a setter function for the session struct
*/
func (s *Session) SetKey(key string, value string) {
	s.Sessionmap[key] = value
}

/*
DeleteKey function
deletes a key from the session
chances are this will never be used
*/
func (s *Session) DeleteKey(key string) {
	delete(s.Sessionmap, key)
}

/*
ToJSON function
returns a json encoded, string representation of the session object
*/
func (s *Session) ToJSON() string {
	data, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(data)
}

/*
FromJSON function gets a json string and builds the session
*/
func (s *Session) FromJSON(res string) {
	json.Unmarshal([]byte(res), s)
}
