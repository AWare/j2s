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
	if response.Status() != 200 {
		fmt.Println("Could not download JSON. Sorry.")
		fmt.Println(response.Error)
	}
	err = generator.WriteGo(result, name, w)
	if err != nil {
		fmt.Print(err)
	}

}

type writeAndPrint int

func (w writeAndPrint) Write(b []byte) (n int, err error) {
	fmt.Print(string(b))
	return len(b), nil
}
