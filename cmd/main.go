// The wichess command line program is separated from package wichess to share application symbols
// with the test package.
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/pciet/wichess"
	"github.com/pciet/wichess/memory"
)

var portFlag = flag.Int("port", 8080, "Internet protocol port `number`")

func main() {
	flag.Parse()
	port := *portFlag
	if (port <= 0) || (port > 65535) {
		fmt.Println("invalid port number", port)
		os.Exit(1)
	}

	memory.Initialize()

	// HTTP handlers are how people get the game interface in their web browser and how game
	// interactions are done with this host program.
	wichess.InitializeHTTP()

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("http.ListenAndServe:", err.Error())
		os.Exit(1)
	}
}
