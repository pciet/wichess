package wichess

import (
	"fmt"
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
