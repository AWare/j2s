package generator

import (
	"io"
	"reflect"
	"unicode"
	"unicode/utf8"
)

func getType(input interface{}, name string, w io.Writer) error {
	//If there is no data returned, then the field is null, and we shall return early
	if input == nil {
		return nil
	}
	w.Write([]byte(getExportableName(name)))
	switch input.(type) {
	default:
		w.Write([]byte(" " + reflect.TypeOf(input).Name() + "  `json:\"" + name + "\"`\n"))
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
		err := getType(v, k, w)
		if err != nil {
			return err
		}

	}
	w.Write([]byte("}\n"))
	return nil
}

func getArrayTypes(input []interface{}, name string, w io.Writer) error {
	w.Write([]byte("[]"))
	return getType(input[0], "", w)
}

func getExportableName(name string) string {
	if name == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(name)
	return string(unicode.ToUpper(r)) + name[n:]

}
