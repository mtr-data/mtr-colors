package main

import (
	"fmt"
	"os"
	datahandlers "src/data_handlers"
	"src/utils"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		err := fmt.Errorf("Usage: go run . <yaml_file_name>")
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	filename := args[0]
	data, err := utils.ReadFileAsYaml(filename)
	checkErr(err)

	handler := datahandlers.HandleLogo
	result, err := handler(data)
	checkErr(err)

	for key, value := range *result {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}
}
