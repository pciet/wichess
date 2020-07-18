package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pciet/wichess/test"
)

const SaveCaseCategoryQuery = "cat"

func SaveCaseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	category := r.URL.Query()[SaveCaseCategoryQuery][0]
	if category == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("no " + SaveCaseCategoryQuery + " query in URL " + r.URL.String())
		return
	}

	var c test.MovesCaseJSON
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err.Error())
		return
	}

	SaveCaseInCategory(category, c)
}
