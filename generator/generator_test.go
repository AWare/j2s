package generator

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"go/parser"
	"go/printer"
	"go/token"

	"github.com/jmcvetta/napping"
)

//TODO: Write test cases.
func TestLeaf(t *testing.T) {
	/*var b bytes.Buffer
	w := bufio.NewWriter(&b)*/
	var a pw
	type leaf int
	var x leaf
	err := GetType(x, "x", a)
	if err != nil {
		t.Error(err)
	}
	//w.Flush()
	//a := b.String()
	//if a != "x generator.leaf\n" {
	//		t.Error("Expecting 'x generator.leaf', got " + a)
	//}

}

func TestWriter(t *testing.T) {
	/*var b bytes.Buffer
	w := bufio.NewWriter(&b)*/
	var a pw
	m := make(map[string]interface{})
	m["something"] = "hello"
	m["cat"] = 5
	n := make(map[string]interface{})
	n["a"] = "b"
	m["a"] = n
	err := WriteGo(m, "x", a)
	if err != nil {
		t.Error(err)
	}
	//w.Flush()
	//a := b.String()
	//if a != "x generator.leaf\n" {
	//		t.Error("Expecting 'x generator.leaf', got " + a)
	//}

}

//Ideally this should test some JSON that's stored here as a file.
//Write the generated code to a file
//And try and unmarshall against this.
func TestComplexJSON(t *testing.T) {
	var result map[string]interface{}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	url := "https://alexwa.re/somejson.json"
	resp, err := napping.Get(url, nil, &result, nil)
	if err != nil {
		t.Error(err)
	}
	if resp.Status() == 200 {
		w.Write([]byte("package main \ntype resp "))
		err := GetType(result, "", w)
		if err != nil {
			t.Error(err)
		}
		w.Flush()
		a := b.String()
		//	e, err := parser.ParseExpr(a)
		fset := token.NewFileSet()
		e, err := parser.ParseFile(fset, "", a, parser.AllErrors)
		if err != nil {
			t.Errorf("Source code generated was not valid go. %s", err.Error())
		}

		var p pw
		printer.Fprint(p, fset, e)
	}
}
func TestEmptyArray(t *testing.T) {
	byt := []byte(`{"empty":[]}`)

	var m map[string]interface{}
	if err := json.Unmarshal(byt, &m); err != nil {
		panic(err)
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	WriteGo(m, "x", w)
	w.Flush()

	fmt.Println(b.String())

}
func TestStruct(t *testing.T) {
	m := make(map[string]interface{})
	m["something"] = "hello"
	m["cat"] = 5
	n := make(map[string]interface{})
	n["a"] = "b"
	m["a"] = n
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	GetType(m, "x", w)
	w.Flush()

	fmt.Println(b.String())

}

type pw int

func (p pw) Write(b []byte) (n int, err error) {
	fmt.Print(string(b))
	return len(b), nil
}
