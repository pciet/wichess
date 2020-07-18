package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pciet/wichess/test"
)

const CategoryNameQuery = "name"

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query()[CategoryNameQuery][0]
	if name == "" {
		fmt.Println("bad category URL", r.URL)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	f := OpenCategoryFile(name)
	defer f.Close()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := io.Copy(w, f)
	if err != nil {
		Panic(err)
	}
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categoryFiles := test.MovesCategoryFilenames("../cases")

	categories := make([]string, len(categoryFiles))
	for i, filename := range categoryFiles {
		c, _ := fmt.Sscanf(strings.TrimSuffix(filename, ".json"), "moves_%s", &(categories[i]))
		if c != 1 {
			Panic("couldn't parse", filename)
		}
	}

	j, err := json.Marshal(categories)
	if err != nil {
		Panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
