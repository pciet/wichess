package wichess

import (
	"encoding/json"
	"net/http"
)

// jsonMarshallable documents that an interface{} var should be able to be marshalled into
// JSON using encoding/json.
type jsonMarshallable interface{}

func jsonResponse(w http.ResponseWriter, content jsonMarshallable) {
	j, err := json.Marshal(content)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
