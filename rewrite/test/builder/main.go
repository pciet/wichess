// github.com/pciet/wichess/test/builder is a program to
// visually make test cases for the position's moves and
// position after move tests.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var builderPage []byte

func init() {
	var err error
	builderPage, err = ioutil.ReadFile("builder.html")
	if err != nil {
		Panic(err)
	}
}

func BuilderPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, string(builderPage))
}

func main() {
	http.HandleFunc("/", BuilderPageHandler)

	http.Handle("/web/",
		http.StripPrefix("/web/",
			http.FileServer(http.Dir("./"))))

	http.Handle("/img/",
		http.StripPrefix("/img/",
			http.FileServer(http.Dir("../../web/img"))))

	// app JavaScript is reused in the /wichess/ path
	http.Handle("/wichess/",
		http.StripPrefix("/wichess/",
			http.FileServer(http.Dir("../../web/js"))))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		Panic(err)
	}
}

func Panic(a ...interface{}) { panic(fmt.Sprintln(a...)) }
