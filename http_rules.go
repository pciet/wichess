package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	RulesPath = "/rules"
	RulesHTML = "html/rules.html"
)

var rulesPage []byte

func RulesGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		DebugPrintln(RulesPath, "bad method", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, string(rulesPage))
}

func init() {
	var err error
	rulesPage, err = ioutil.ReadFile(RulesHTML)
	if err != nil {
		Panic(err)
	}
}
