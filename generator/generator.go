//Package generator generates go struct definitions from map[string]interface{},
//which is what json is unmarshalled to when not given a type.
//excuse the early state of this code
package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"reflect"
	"unicode"
	"unicode/utf8"
)

//WriteGo is a function which takes the output of json.Unmarshal and writes out
//some nicely formatted go code to an io.Writer. Run gofmt on it anyway.
func WriteGo(input interface{}, name string, w io.Writer) error {
	_ = "breakpoint"

	code, err := generateGo(name, input)
	fset := token.NewFileSet()
	ast, err := parser.ParseFile(fset, "", code, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("Looks like the code we generated wasn't valid go:\n %s \n The code generated looked like:\n %s", err.Error(), code)
	}
	printer.Fprint(w, fset, ast)
	return nil
}

//GetType will take the output of json.Unmarshal and return some barely formatted go.
//This is useful if you are going to just copy the defined struct somewhere and run gofmt on it.
func GetType(input interface{}, name string, w io.Writer) error {
	//If there is no data returned, then the field is null, and we shall return early
	if input == nil {
		return nil
	}
	switch input.(type) {
	default:
		w.Write([]byte(getExportableName(name) + " " + reflect.TypeOf(input).Name()))
		writeJSONtag(name, w)
		return nil
	case map[string]interface{}:
		return getTypes(input.(map[string]interface{}), name, w)
	case []interface{}:
		return getArrayTypes(input.([]interface{}), name, w)
	}

}

func getTypes(input map[string]interface{}, name string, w io.Writer) error {
	w.Write([]byte(getExportableName(name) + " struct {\n"))
	for k, v := range input {
		err := GetType(v, k, w)
		if err != nil {
			return err
		}

	}
	w.Write([]byte("}"))
	writeJSONtag(name, w)
	return nil
}

func getArrayTypes(input []interface{}, name string, w io.Writer) error {
	if len(input) == 0 {
		w.Write([]byte(fmt.Sprintf("//Empty array found with name: %s \n", name)))
		return nil
	}
	w.Write([]byte(getExportableName(name) + " []"))
	err := GetType(input[0], "", w)
	if err != nil {
		return err
	}
	writeJSONtag(name, w)
	return nil

}

func getExportableName(name string) string {
	if name == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(name)
	return string(unicode.ToUpper(r)) + name[n:]

}

func generateGo(name string, thing interface{}) (string, error) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	w.Write([]byte("package generatedCode \ntype " + name + " "))
	err := GetType(thing, "", w)
	if err != nil {
		return "", err
	}
	w.Flush()
	return b.String(), nil
}

func writeJSONtag(name string, w io.Writer) {
	if name != "" {
		w.Write([]byte("  `json:\"" + name + "\"`\n"))
	}
}
