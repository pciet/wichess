package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pciet/wichess/test"
)

func SaveCaseInCategory(category string, testCase test.MovesCaseJSON) {
	f := OpenCategoryFile(category)
	defer f.Close()

	mc, err := ioutil.ReadAll(f)
	if err != nil {
		Panic(err)
	}

	js := test.ParseMovesCategoryJSON(mc)

	// either replace the case with the same title, or insert at the end
	replaced := false
	for i := 0; i < len(js); i++ {
		if js[i].Name == testCase.Name {
			js[i] = testCase
			replaced = true
			break
		}
	}
	if replaced == false {
		js = append(js, testCase)
	}

	out, err := json.MarshalIndent(test.CategoryFileJSON{js}, "", "    ")
	if err != nil {
		Panic(err)
	}

	c, err := f.WriteAt(out, 0)
	if err != nil {
		Panic(err)
	}

	err = f.Truncate(int64(c))
	if err != nil {
		Panic(err)
	}
}

func OpenCategoryFile(name string) *os.File {
	f, err := os.OpenFile("../cases/moves_"+name+".json", os.O_RDWR, 0644)
	if err != nil {
		f.Close()
		Panic(err)
	}
	return f
}
