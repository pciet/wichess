package wichess

import (
	"html/template"
	"net/http"

	"github.com/pciet/wichess/piece"
)

// DetailsPieceQuery is the URL query key for a string with the codename of the piece for
// DetailsPath.
const DetailsPieceQuery = "p"

type DetailsHTMLTemplateData struct {
	Name     string // readable piece name
	CodeName string // lowercase word used for URLs

	Description template.HTML

	CharacteristicA            string // optional, name of the first characteristic
	CharacteristicADescription string

	CharacteristicB            string // optional
	CharacteristicBDescription string
}

func detailsTemplateData(pieceCodeName string, k piece.Kind) DetailsHTMLTemplateData {
	t := DetailsHTMLTemplateData{
		Name:        piece.Name(k),
		CodeName:    pieceCodeName,
		Description: template.HTML(piece.DetailsHTML(k)),
	}

	charA, charB := piece.Characteristics(k)

	if charA == piece.NoCharacteristic {
		return t
	}

	t.CharacteristicA = piece.CharacteristicName(charA)
	t.CharacteristicADescription = piece.CharacteristicDescription(charA)

	if charB == piece.NoCharacteristic {
		return t
	}

	t.CharacteristicB = piece.CharacteristicName(charB)
	t.CharacteristicBDescription = piece.CharacteristicDescription(charB)

	return t
}

func detailsGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		debug(DetailsPath, "bad method", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	p, err := parseURLQuery(r.URL.Query(), DetailsPieceQuery)
	if err != nil {
		debug(DetailsPath, "bad URL:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	kind := piece.KindForCodeName(p)
	if kind == piece.NoKind {
		debug(DetailsPath, "bad piece code name query argument:", p)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeHTMLTemplate(w, DetailsHTMLTemplate, detailsTemplateData(p, kind))
}
