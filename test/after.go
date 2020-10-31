package test

import (
	"encoding/json"
	"fmt"

	"github.com/pciet/wichess/rules"
)

type AfterMoveCaseJSON struct {
	Name       string `json:"case"`
	rules.Move `json:"mov"`
	Previous   rules.Move     `json:"prev"`
	Position   []testPiece    `json:"pos"`
	Changes    []rules.Square `json:"cha"`
}

type AfterMoveCategoryFileJSON struct {
	Cases []AfterMoveCaseJSON `json:"cases"`
}

func loadAllAfterMoveCases() []AfterMoveCaseJSON {
	caseFiles := parseAllCases(afterMoveParser, "after")
	out := make([]AfterMoveCaseJSON, 0, 8)

	for _, cf := range caseFiles {
		f := cf.(AfterMoveCategoryFileJSON)
		for _, c := range f.Cases {
			out = append(out, c)
		}
	}

	return out
}

func afterMoveParser(out *caseJSON, in []byte) error {
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
