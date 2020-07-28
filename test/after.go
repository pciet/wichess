package test

import (
	"encoding/json"
	"fmt"

	"github.com/pciet/wichess/rules"
)

type AfterMoveCaseJSON struct {
	Name       string `json:"case"`
	rules.Move `json:"mov"`
	Position   []Piece                 `json:"pos"`
	Changes    []rules.AddressedSquare `json:"cha"`
}

type AfterMoveCategoryFileJSON struct {
	Cases []AfterMoveCaseJSON `json:"cases"`
}

func LoadAllAfterMoveCases() []AfterMoveCaseJSON {
	caseFiles := ParseAllCases(AfterMoveParser, "after")
	out := make([]AfterMoveCaseJSON, 0, 8)

	for _, cf := range caseFiles {
		f := cf.(AfterMoveCategoryFileJSON)
		for _, c := range f.Cases {
			out = append(out, c)
		}
	}

	return out
}

func AfterMoveParser(out *CaseJSON, in []byte) error {
	var value AfterMoveCategoryFileJSON
	err := json.Unmarshal(in, &value)
	*out = value
	return err
}

func (a AfterMoveCaseJSON) String() string {
	out := a.Name + "\n"
	out += "Position\n"
	out += fmt.Sprintf("%v\n", a.Position)
	out += "Move " + a.Move.String() + "\n"
	out += "Changes\n"
	out += fmt.Sprintf("%v\n", a.Changes)

	return out
}
