package test

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

const CaseDir = "./cases"

// CategoryFilenames lists all filenames in the directory that contain the tag string and have
// the .json extension.
func CategoryFilenames(dir, tag string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err.Error())
	}

	categoryFiles := make([]string, 0, 2)

	for _, f := range files {
		n := f.Name()
		if (strings.Contains(n, tag) == false) || (filepath.Ext(n) != ".json") {
			continue
		}
		categoryFiles = append(categoryFiles, n)
	}

	if len(categoryFiles) == 0 {
		panic("no " + tag + " test case category files found")
	}

	return categoryFiles
}

func LoadAllCases(tag string) [][]byte {
	categoryFiles := CategoryFilenames(CaseDir, tag)
	out := make([][]byte, 0, 8)
	for _, f := range categoryFiles {
		out = append(out, LoadFile(CaseDir+"/"+f))
	}

	return out
}

func LoadCategoryFile(dir, tag, name string) []byte {
	return LoadFile(dir + "/" + tag + "_" + name + ".json")
}

func LoadFile(name string) []byte {
	c, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err.Error())
	}
	return c
}
