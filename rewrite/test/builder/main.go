// wichess/test/builder is a program to visually make test cases for tests in the test directory.
// For now this is just a test of expected moves for a position.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", BuilderPageHandler)
	http.HandleFunc("/categories", CategoriesHandler)
	http.HandleFunc("/category", CategoryHandler)
	http.HandleFunc("/savecase", SaveCaseHandler)

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("./"))))

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("../../web/img"))))

	// app JavaScript is reused in the /wichess/ path
	http.Handle("/wichess/", http.StripPrefix("/wichess/",
		http.FileServer(http.Dir("../../web"))))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		Panic(err)
	}
}

var builderPage []byte

func BuilderPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, string(builderPage))
}

func init() {
	var err error
	builderPage, err = ioutil.ReadFile("builder.html")
	if err != nil {
		Panic(err)
	}
}

func Panic(a ...interface{}) { panic(fmt.Sprintln(a...)) }
