package utils

import "html/template"

/*
LoadTemplates function
loads all the provided templates
*/
func LoadTemplates(tplname string, tpl ...string) (*template.Template, error) {
	t := template.New(tplname)
	t, err := t.ParseFiles(tpl...)
	return t, err
}
