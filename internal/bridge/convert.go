package bridge

import (
	"github.com/datadog/go-python3"
)

// LoadModule is a wrapper for python.LoadModule
func Convert(input string) string {
	module := LoadModule("./convert")
	function := module.GetAttrString("Convert")
	args := python3.PyTuple_New(1)
	python3.PyTuple_SetItem(args, 0, python3.PyUnicode_FromString(input))
	return python3.PyUnicode_AsUTF8(function.Call(args, python3.Py_None))
}
