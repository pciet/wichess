package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/pciet/wichess/piece"
	"github.com/pciet/wichess/rules"
)

type Piece struct {
	Address     rules.Address     `json:"addr"`
	Kind        piece.Kind        `json:"k"`
	Orientation rules.Orientation `json:"o"`
	Moved       bool              `json:"m"`
}

type MovesCaseJSON struct {
	Name         string            `json:"case"`
	Active       rules.Orientation `json:"active"`
	PreviousMove rules.Move        `json:"prev"`
	State        rules.State       `json:"state"`
	Position     []Piece           `json:"pos"`
	Moves        []rules.MoveSet   `json:"moves"`
}

type CategoryFileJSON struct {
	Cases []MovesCaseJSON `json:"cases"`
}

func ParseMovesCategoryJSON(t []byte) []MovesCaseJSON {
	var out CategoryFileJSON
	err := json.Unmarshal(t, &out)
	if err != nil {
		panic(err.Error())
	}
	return out.Cases
}

func MovesCategoryFilenames(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err.Error())
	}

	categoryFiles := make([]string, 0, 2)

	for _, f := range files {
		n := f.Name()
		if (strings.Contains(n, "moves") == false) || (filepath.Ext(n) != ".json") {
			continue
		}
		categoryFiles = append(categoryFiles, n)
	}

	if len(categoryFiles) == 0 {
		panic("no moves test case category files found")
	}

	return categoryFiles
}

const CaseDir = "./cases"

func LoadMovesCases() []MovesCaseJSON {
	categoryFiles := MovesCategoryFilenames(CaseDir)
	cases := make([]MovesCaseJSON, 0, 8)

	for _, f := range categoryFiles {
		c, err := ioutil.ReadFile(CaseDir + "/" + f)
		if err != nil {
			panic(err.Error())
		}
		cases = append(cases, ParseMovesCategoryJSON(c)...)
	}

	return cases
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
