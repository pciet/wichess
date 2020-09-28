// Builder is a small web app to visually inspect and make test cases for package tests.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	MovesTag     = "moves"
	AfterMoveTag = "after"
)

func main() {
	http.HandleFunc("/", HTMLHandler(indexPage))

	http.HandleFunc("/moves", HTMLHandler(movesPage))
	http.HandleFunc("/moves/categories", CategoriesHandler(MovesTag))
	http.HandleFunc("/moves/category", CategoryHandler(MovesTag))
	http.HandleFunc("/moves/save", SaveCategoryHandler(MovesTag))

	http.HandleFunc("/after", HTMLHandler(afterPage))
	http.HandleFunc("/after/categories", CategoriesHandler(AfterMoveTag))
	http.HandleFunc("/after/category", CategoryHandler(AfterMoveTag))
	http.HandleFunc("/after/save", SaveCategoryHandler(AfterMoveTag))

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("../../web/img"))))

	// app JavaScript is reused in the /wichess/ path
	http.Handle("/wichess/", http.StripPrefix("/wichess/", http.FileServer(http.Dir("../../web"))))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		Panic(err)
	}
}

var movesPage, afterPage, indexPage []byte

func HTMLHandler(b []byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, string(b))
	}
}

func init() {
	var err error
	movesPage, err = ioutil.ReadFile("html/moves.html")
	if err != nil {
		Panic(err)
	}
	afterPage, err = ioutil.ReadFile("html/after.html")
	if err != nil {
		Panic(err)
	}
	indexPage, err = ioutil.ReadFile("html/index.html")
	if err != nil {
		Panic(err)
	}
}

func Panic(a ...interface{}) { panic(fmt.Sprintln(a...)) }
