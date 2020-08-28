package wichess

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

func LoadHTMLTemplates() {
	ParseHTMLTemplate(DetailsHTMLTemplate, GameHTMLTemplate, IndexHTMLTemplate, LoginHTMLTemplate,
		MatchHTMLTemplate, RewardHTMLTemplate)

	var err error
	rulesPage, err = ioutil.ReadFile(RulesHTML)
	if err != nil {
		Panic(err)
	}
}

var HTMLTemplates = map[string]*template.Template{}

func ParseHTMLTemplate(filenames ...string) {
	for _, file := range filenames {
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
}

func WriteHTMLTemplate(w http.ResponseWriter, file string, data interface{}) {
	t, has := HTMLTemplates[file]
	if has == false {
		Panic(file, "template not parsed")
	}
	err := t.Execute(w, data)
	if err != nil {
		DebugPrintln("failed to execute template", file, ":", err)
	}
}
