package wichess

import (
	"database/sql"
	"net/http"
)

const CollectionPath = "/collection"

var CollectionHandler = AuthenticRequestHandler{
	Get: CollectionGet,
}

type CollectionJSON struct {
	Collection `json:"c"`
}

func CollectionGet(w http.ResponseWriter, r *http.Request, tx *sql.Tx, requester Player) {
	coll := PlayerCollection(tx, requester.ID)
	tx.Commit()
	JSONResponse(w, CollectionJSON{coll})
}
