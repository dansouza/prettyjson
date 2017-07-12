package main

import (
	"fmt"

	"io"
	"os"

	"encoding/json"

	"github.com/hokaccha/go-prettyjson"
	"github.com/mattn/go-colorable"
)

func main() {
	var content map[string]interface{}
	var file io.ReadCloser
	var err error

	if len(os.Args) == 2 {
		if os.Args[1] != "-" {
			file, err = os.Open(os.Args[1])
			if err != nil {
				fmt.Printf("error: %s: %s", os.Args[1], err)
				os.Exit(-1)
			}
			defer file.Close()
		}
	} else {
		file = os.Stdin
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&content)
	if err != nil {
		fmt.Printf("error decoding json: %s", err)
		os.Exit(-1)
	}

	pretty, err := prettyjson.Marshal(content)
	if err != nil {
		fmt.Printf("error pretty-printing json: %s", err)
		os.Exit(-1)
	}

	colorable.NewColorableStdout().Write(pretty)
	fmt.Print("\n")
}
