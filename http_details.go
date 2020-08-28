package wichess

import (
	"html/template"
	"net/http"

	"github.com/pciet/wichess/piece"
)

const (
	DetailsPath         = "/details"
	DetailsHTMLTemplate = "html/details.tmpl"
	DetailsPieceQuery   = "p"
)

type DetailsHTMLTemplateData struct {
	Name     string // readable piece name
	CodeName string // lowercase word used for URLs

	Description template.HTML

	CharacteristicA            string // optional, name of the first characteristic
	CharacteristicADescription string

	CharacteristicB            string // optional
	CharacteristicBDescription string
}

func DetailsTemplateData(pieceCodeName string, k piece.Kind) DetailsHTMLTemplateData {
	t := DetailsHTMLTemplateData{
		Name:        piece.Names[k],
		CodeName:    pieceCodeName,
		Description: template.HTML(piece.DetailsHTML[k]),
	}

	chars := piece.CharacteristicList[k]

	if chars.A == piece.NoCharacteristic {
		return t
	}

	t.CharacteristicA = piece.CharacteristicNames[chars.A]
	t.CharacteristicADescription = piece.CharacteristicDescriptions[chars.A]

	if chars.B == piece.NoCharacteristic {
		return t
	}

	t.CharacteristicB = piece.CharacteristicNames[chars.B]
	t.CharacteristicBDescription = piece.CharacteristicDescriptions[chars.B]

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

	kind := piece.KindForCodeName(p)
	if kind == piece.NoKind {
		DebugPrintln(DetailsPath, "bad piece code name query argument:", p)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	WriteHTMLTemplate(w, DetailsHTMLTemplate, DetailsTemplateData(p, kind))
}
