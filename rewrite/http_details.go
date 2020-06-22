package main

import (
	"net/http"

	"github.com/pciet/wichess/rules"
)

const (
	DetailsPath         = "/details"
	DetailsHTMLTemplate = "html/details.tmpl"
	DetailsPieceQuery   = "p"
)

type DetailsHTMLTemplateData struct {
	Name     string // readable piece name
	CodeName string // lowercase word used for URLs

	Basic         string // optional readable piece name for the basic kind of this piece
	BasicCodeName string

	// if Basic isn't defined then StartDescription and MovesDescription are required, otherwise
	// they are unused

	StartDescription string // description of the starting squares
	MovesDescription string // description of the piece moves

	CharacteristicA            string // optional, name of the first characteristic
	CharacteristicADescription string

	CharacteristicB            string // optional
	CharacteristicBDescription string
}

func DetailsTemplateData(pieceCodeName string) DetailsHTMLTemplateData {
	kind := rules.CodeNameKind(pieceCodeName)

	t := DetailsHTMLTemplateData{
		Name:     kind.Name(),
		CodeName: pieceCodeName,
	}

	basic := rules.BasicKind(kind)
	if basic != kind {
		t.Basic = basic.Name()
		t.BasicCodeName = basic.CodeName()
	}

	t.StartDescription = kind.StartDescription()
	t.MovesDescription = kind.MovesDescription()

	char := kind.CharacteristicAString()
	t.CharacteristicA = char
	t.CharacteristicADescription = rules.CharacteristicDescription(char)

	char = kind.CharacteristicBString()
	t.CharacteristicB = char
	t.CharacteristicBDescription = rules.CharacteristicDescription(char)

	return t
}

func DetailsGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		DebugPrintln(DetailsPath, "bad method", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	p, err := ParseURLQuery(r.URL.Query(), DetailsPieceQuery)
	if err != nil {
		DebugPrintln(DetailsPath, "bad URL:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if rules.IsPieceCodeName(p) == false {
		DebugPrintln(DetailsPath, "bad piece code name query argument:", p)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	WriteHTMLTemplate(w, DetailsHTMLTemplate, DetailsTemplateData(p))
}

func init() { ParseHTMLTemplate(DetailsHTMLTemplate) }
