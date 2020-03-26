package main

import (
	"encoding/json"
	"net/http"
)

// JSONMarshallable documents that an interface{} var should be
// able to be marshalled into JSON using encoding/json.
type JSONMarshallable interface{}

func JSONResponse(w http.ResponseWriter, content JSONMarshallable) {
	j, err := json.Marshal(content)
	if err != nil {
		Panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
