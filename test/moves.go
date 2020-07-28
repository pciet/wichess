package test

import (
	"encoding/json"
	"fmt"

	"github.com/pciet/wichess/rules"
)

type MovesCaseJSON struct {
	Name         string            `json:"case"`
	Active       rules.Orientation `json:"active"`
	PreviousMove rules.Move        `json:"prev"`
	State        rules.State       `json:"state"`
	Position     []Piece           `json:"pos"`
	Moves        []rules.MoveSet   `json:"moves"`
}

type MovesCategoryFileJSON struct {
	Cases []MovesCaseJSON `json:"cases"`
}

func LoadAllMovesCases() []MovesCaseJSON {
	caseFiles := ParseAllCases(MovesParser, "moves")
	out := make([]MovesCaseJSON, 0, 8)

	for _, cf := range caseFiles {
		f := cf.(MovesCategoryFileJSON)
		for _, c := range f.Cases {
			out = append(out, c)
		}
	}

	return out
}

func MovesParser(out *CaseJSON, in []byte) error {
	var value MovesCategoryFileJSON
	err := json.Unmarshal(in, &value)
	*out = value
	return err
}

func (a MovesCaseJSON) String() string {
	out := a.Name + "\n"
	out += "Active " + a.Active.String() + "\n"
	out += "Previous Move " + a.PreviousMove.String() + "\n"
	out += "State " + a.State.String() + "\n"
	out += "Position\n"
	out += fmt.Sprintf("%v\n", a.Position)
	out += fmt.Sprintf("%v\n", a.Moves)
	return out
}
