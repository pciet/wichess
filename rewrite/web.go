package main

import (
	"html/template"
	"net/http"
)

var HTMLTemplates = map[string]*template.Template{}

// This function should be called for all HTML templates in init functions and not concurrently.
func ParseHTMLTemplate(file string) {
	_, has := HTMLTemplates[file]
	if has {
		panic(file, "already parsed")
	}
	t, err := template.ParseFiles(file)
	if err != nil {
		panic(file, "failed to parse", file, "-", err)
	}
	HTMLTemplates[file] = t
}

func WriteWebTemplate(w http.ResponseWriter, file string, data interface{}) {
	t, has := HTMLTemplates[file]
	if has == false {
		panic(file, "template not parsed")
	}
	err := t.Execute(w, data)
	if err != nil {
		DebugPrintln("failed to execute template", file, "-", err)
	}
}
