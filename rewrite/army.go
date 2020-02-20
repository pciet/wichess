package main

import (
	"encoding/json"
	"io"
)

// When a player requests a new game they specify which of their pieces to include in an ArmyRequest.
type ArmyRequest [16]PieceIdentifier

func DecodeArmyRequest(jsonBody io.Reader) (ArmyRequest, error) {
	var a ArmyRequest
	err := json.NewDecoder(jsonBody).Decode(&a)
	return a, err
}
