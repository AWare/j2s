package generator

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	j := `{"array":[1,2,3,4,5]}`
	testJSONString(j, t)
}

func TestEmptyArray(t *testing.T) {
	j := `{"empty":[]}`
	testJSONString(j, t)
}

func TestStruct(t *testing.T) {
	j := `{"fieldA":"a","fieldB":0}`
	testJSONString(j, t)
}
func TestNullField(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err := GetType(nil, "name", w)
	if err != nil {
		t.Error(err)
	}
	w.Flush()
	if len(b.Bytes()) != 0 {
		t.Errorf("Expected no written data, got: %s", b.String())
	}

}
func testJSONString(j string, t *testing.T) {
	fmt.Printf("From:\n%s\n", j)

	m, err := unmarshallFromString(j)
	if err != nil {
		t.Error(err)
	}
	code, err := generateCode(m)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("The following go was generated: %s", code)
}

func generateCode(input map[string]interface{}) (string, error) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err := WriteGo(input, "x", w)
	if err != nil {
		return "", err
	}
	w.Flush()
	return b.String(), nil
}

func unmarshallFromString(input string) (map[string]interface{}, error) {
	byt := []byte(input)
	var m map[string]interface{}
	err := json.Unmarshal(byt, &m)
	return m, err

}
