package main

import (
	"fmt"
	"os"
	datahandlers "src/data_handlers"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		err := fmt.Errorf("Usage: go run . <yaml_file_name>")
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	filename := args[0]
	logoData, err := datahandlers.HandleLogo(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for key, value := range *logoData {
		fmt.Printf("Key: %s, value: %v\n", key, value)
	}
}
