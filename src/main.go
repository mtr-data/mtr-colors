package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		err := fmt.Errorf("Usage: go run . <yaml_file_name>")
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	filename := args[0]
	str, err := ReadFileAsYamlStr(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(str)
}
