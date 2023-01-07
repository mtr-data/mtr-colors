package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFileAsBytes(filename string) (data []byte, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	size := fileInfo.Size()
	buf := make([]byte, size)
	_, err = file.Read(buf)

	return buf, err
}

func ReadFileAsYaml(filename string) (out *yaml.Node, err error) {
	var root yaml.Node

	data, err := ReadFileAsBytes(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &root)
	return &root, err
}

func ReadFileAsYamlStr(filename string) (out string, err error) {
	root, err := ReadFileAsYaml(filename)
	if err != nil {
		return "", err
	}

	data, err := yaml.Marshal(root)
	str := string(data)

	return str, err
}
