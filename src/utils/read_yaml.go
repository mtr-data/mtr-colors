package utils

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

func ReadFileAsYaml(filename string) (out interface{}, err error) {
	data, err := ReadFileAsBytes(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &out)
	return out, err
}
