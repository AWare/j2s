package main

import (
	"fmt"
	"os"

	"github.com/AWare/j2s/generator"
	"github.com/jmcvetta/napping"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Give me a URL, and a name and off we go.")
		return
	}
	url := args[1]
	name := args[2]
	var result map[string]interface{}
	var w writeAndPrint
	response, err := napping.Get(url, nil, &result, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	if response.Status() == 200 {
		fmt.Print("HEY")
		generator.WriteGo(result, name, w)
	}

}

type writeAndPrint int

func (w writeAndPrint) Write(b []byte) (n int, err error) {
	fmt.Print(string(b))
	return len(b), nil
}
