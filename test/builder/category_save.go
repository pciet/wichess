package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pciet/wichess/test"
)

// TODO: reduce repetition between SaveMovesCase and SaveAfterMoveCase

func SaveMovesCase(category string, c test.MovesCaseJSON) {
	f := OpenCategoryFile(MovesTag, category)
	defer f.Close()

	fc, err := ioutil.ReadAll(f)
	if err != nil {
		Panic(err)
	}

	var caseFile test.MovesCategoryFileJSON
	err = json.Unmarshal(fc, &caseFile)
	if err != nil {
		Panic(err)
	}

	// either replace the case with the same title, or insert at the end
	replaced := false
	for i := 0; i < len(caseFile.Cases); i++ {
		if caseFile.Cases[i].Name == c.Name {
			caseFile.Cases[i] = c
			replaced = true
			break
		}
	}
	if replaced == false {
		caseFile.Cases = append(caseFile.Cases, c)
	}

	out, err := json.MarshalIndent(caseFile, "", "    ")
	if err != nil {
		Panic(err)
	}

	WriteCategoryFile(f, out)
}

func SaveAfterMoveCase(category string, c test.AfterMoveCaseJSON) {
	f := OpenCategoryFile(AfterMoveTag, category)
	defer f.Close()

	fc, err := ioutil.ReadAll(f)
	if err != nil {
		Panic(err)
	}

	var caseFile test.AfterMoveCategoryFileJSON
	err = json.Unmarshal(fc, &caseFile)
	if err != nil {
		Panic(err)
	}

	// either replace the case with the same title, or insert at the end
	replaced := false
	for i := 0; i < len(caseFile.Cases); i++ {
		if caseFile.Cases[i].Name == c.Name {
			caseFile.Cases[i] = c
			replaced = true
			break
		}
	}
	if replaced == false {
		caseFile.Cases = append(caseFile.Cases, c)
	}

	out, err := json.MarshalIndent(caseFile, "", "    ")
	if err != nil {
		Panic(err)
	}

	WriteCategoryFile(f, out)
}
