package main

import (
	"encoding/json"

	"github.com/ckndave/hclparser/convert"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Module.Get("exports").Set("parseToString", parseToString)
	js.Module.Get("exports").Set("parseToObject", parseToObject)
}

// Parse a HCL string into a JSON string
func parseToString(input string) (output string, err error) {
	convertedBytes, err := convert.String(input)
	if err != nil {
		return "", err
	}
	bytes, _ := json.Marshal(convertedBytes)
	output = string(bytes)
	return output, nil
}

// Parse a HCL string into a JSON object
func parseToObject(input string) (output *js.Object, err error) {
	jsonString, err := parseToString(input)
	if err != nil {
		empty := js.Object{}
		return &empty, err
	}

	obj := js.Global.Get("JSON").Call("parse", string(jsonString))
	return obj, nil
}
