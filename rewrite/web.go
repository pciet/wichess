package main

import (
	"html/template"
	"log"
	"net/http"
)

var HTMLTemplates = map[string]*template.Template{}

// All HTML templates are parsed with ParseHTMLTemplate in init functions.
func ParseHTMLTemplate(file string) {
	_, has := HTMLTemplates[file]
	if has {
		log.Panicln(file, "already parsed")
	}
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Panicln("failed to parse", file, "-", err)
	}
	HTMLTemplates[file] = t
}

func WriteWebTemplate(w http.ResponseWriter, file string, data interface{}) {
	t, has := HTMLTemplates[file]
	if has == false {
		log.Panicln(file, "template not parsed")
	}
	err := t.Execute(w, data)
	if err != nil {
		DebugPrintln("failed to execute template", file, "-", err)
	}
}
