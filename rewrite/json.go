package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// able to be marshalled to JSON
type JSONMarshallable interface{}

func JSONResponse(w http.ResponseWriter, content JSONMarshallable) {
	j, err := json.Marshal(content)
	if err != nil {
		log.Panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
