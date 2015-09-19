//Package generator generates go struct definitions from map[string]interface{},
//which is what json is unmarshalled to when not given a type.
//excuse the early state of this code
package generator

import (
	"io"
	"reflect"
	"unicode"
	"unicode/utf8"
)

func GetType(input interface{}, name string, w io.Writer) error {
	//If there is no data returned, then the field is null, and we shall return early
	if input == nil {
		return nil
	}
	w.Write([]byte(getExportableName(name)))
	if name != "" {
		defer w.Write([]byte("  `json:\"" + name + "\"`\n"))
	}
	switch input.(type) {
	default:
		w.Write([]byte(" " + reflect.TypeOf(input).Name()))
		return nil
	case map[string]interface{}:
		return getTypes(input.(map[string]interface{}), name, w)
	case []interface{}:
		return getArrayTypes(input.([]interface{}), name, w)
	}

}

func getTypes(input map[string]interface{}, name string, w io.Writer) error {
	w.Write([]byte(" struct {\n"))
	for k, v := range input {
		err := GetType(v, k, w)
		if err != nil {
			return err
		}

	}
	w.Write([]byte("}"))
	return nil
}

func getArrayTypes(input []interface{}, name string, w io.Writer) error {
	w.Write([]byte("[]"))
	return GetType(input[0], "", w)
}

func getExportableName(name string) string {
	if name == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(name)
	return string(unicode.ToUpper(r)) + name[n:]

}
