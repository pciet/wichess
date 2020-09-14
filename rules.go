package wichess

import (
	"fmt"
	"net/http"
)

var rulesPage []byte

func rulesGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		debug(RulesPath, "bad method", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, string(rulesPage))
}
