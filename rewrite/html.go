package main

import (
	"html/template"
	"net/http"
)

var HTMLTemplates = map[string]*template.Template{}

// All HTML templates are parsed with ParseHTMLTemplate in init functions.
func ParseHTMLTemplate(file string) {
	_, has := HTMLTemplates[file]
	if has {
		Panic(file, "already parsed")
	}
	t, err := template.ParseFiles(file)
	if err != nil {
		Panic("failed to parse", file, ":", err)
	}
	HTMLTemplates[file] = t
}

func WriteWebTemplate(w http.ResponseWriter, file string, data interface{}) {
	t, has := HTMLTemplates[file]
	if has == false {
		Panic(file, "template not parsed")
	}
	err := t.Execute(w, data)
	if err != nil {
		DebugPrintln("failed to execute template", file, ":", err)
	}
}
