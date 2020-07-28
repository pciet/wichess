package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/pciet/wichess/test"
)

func CategoriesHandler(tag string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		categoryFiles := test.CategoryFilenames("../cases", tag)
		categories := make([]string, len(categoryFiles))
		for i, filename := range categoryFiles {
			c, _ := fmt.Sscanf(
				strings.TrimSuffix(filename, ".json"), tag+"_%s", &(categories[i]))
			if c != 1 {
				Panic("couldn't parse", filename)
			}
		}
		j, err := json.Marshal(categories)
		if err != nil {
			Panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	}
}

const CategoryNameQuery = "name"

func CategoryHandler(tag string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query()[CategoryNameQuery][0]
		if name == "" {
			fmt.Println("bad category URL", r.URL)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, err := w.Write(test.LoadCategoryFile("../"+test.CaseDir, tag, name))
		if err != nil {
			Panic(err)
		}
	}
}

const SaveCategoryQuery = "cat"

func SaveCategoryHandler(tag string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println(r.Method, "not POST")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		category := r.URL.Query()[SaveCategoryQuery][0]
		if category == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("no " + SaveCategoryQuery + " query in URL " + r.URL.String())
			return
		}

		switch tag {
		case MovesTag:
			var c test.MovesCaseJSON
			err := json.NewDecoder(r.Body).Decode(&c)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println(err.Error())
				return
			}
			SaveMovesCase(category, c)
		case AfterMoveTag:
			var c test.AfterMoveCaseJSON
			err := json.NewDecoder(r.Body).Decode(&c)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println(err.Error())
				return
			}
			SaveAfterMoveCase(category, c)
		default:
			Panic("bad tag", tag)
		}
		fmt.Println("saved", category)
	}
}

func WriteCategoryFile(f *os.File, with []byte) {
	count, err := f.WriteAt(with, 0)
	if err != nil {
		Panic(err)
	}

	err = f.Truncate(int64(count))
	if err != nil {
		Panic(err)
	}
}

func OpenCategoryFile(tag, name string) *os.File {
	f, err := os.OpenFile("../cases/"+tag+"_"+name+".json", os.O_RDWR, 0644)
	if err != nil {
		f.Close()
		Panic(err)
	}
	return f
}
