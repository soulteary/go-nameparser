package bridge

import (
	"path/filepath"

	python3 "github.com/go-python/cpy3"
)

func LoadModule(dir string) *python3.PyObject {
	path := python3.PyImport_ImportModule("sys").GetAttrString("path")
	python3.PyList_Insert(path, 0, python3.PyUnicode_FromString(dir))
	return python3.PyImport_ImportModule(filepath.Base(dir))
}
