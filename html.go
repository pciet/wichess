package wichess

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func loadHTMLTemplates() {
	parseHTMLTemplate(DetailsHTMLTemplate, GameHTMLTemplate, IndexHTMLTemplate, LoginHTMLTemplate,
		MatchHTMLTemplate, RewardHTMLTemplate)

	var err error
	rulesPage, err = ioutil.ReadFile(RulesHTML)
	if err != nil {
		panic(err.Error())
	}
}

var htmlTemplates = map[string]*template.Template{}

func parseHTMLTemplate(filenames ...string) {
	for _, file := range filenames {
		_, has := htmlTemplates[file]
		if has {
			log.Panicln(file, "already parsed")
		}
		t, err := template.ParseFiles(file)
		if err != nil {
			log.Panicln("failed to parse", file, ":", err)
		}
		htmlTemplates[file] = t
	}
}

func writeHTMLTemplate(w http.ResponseWriter, file string, data interface{}) {
	t, has := htmlTemplates[file]
	if has == false {
		log.Panicln(file, "template not parsed")
	}
	err := t.Execute(w, data)
	if err != nil {
		debug("failed to execute template", file, ":", err)
	}
}
