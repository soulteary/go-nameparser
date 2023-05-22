package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	python3 "github.com/datadog/go-python3"
)

func LoadModule(dir string) *python3.PyObject {
	path := python3.PyImport_ImportModule("sys").GetAttrString("path")
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(dir))
	return python3.PyImport_ImportModule(filepath.Base(dir))
}

func Convert(input string) string {
	module := LoadModule("./convert")
	function := module.GetAttrString("Convert")
	args := python3.PyTuple_New(1)
	python3.PyTuple_SetItem(args, 0, python3.PyUnicode_FromString(input))
	return python3.PyUnicode_AsUTF8(function.Call(args, python3.Py_None))
}

type HumanName struct {
	Text   string `json:"text"`
	Detail struct {
		Title    string `json:"title"`
		First    string `json:"first"`
		Middle   string `json:"middle"`
		Last     string `json:"last"`
		Suffix   string `json:"suffix"`
		Nickname string `json:"nickname"`
	} `json:"detail"`
}

func main() {
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	if !python3.Py_IsInitialized() {
		log.Fatalln("Failed to initialize Python environment")
	}

	ret := Convert("Dr. Juan Q. Xavier de la Vega III (Doc Vega)")

	var name HumanName
	err := json.Unmarshal([]byte(ret), &name)
	if err != nil {
		fmt.Println("Parsing JSON failed:", err)
		return
	}

	fmt.Println("Name:", name.Text)
	fmt.Println("Detail:", name.Detail)
}
