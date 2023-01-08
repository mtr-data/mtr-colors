package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadFileAsBytes(filename string) (data []byte, err error) {
	return ioutil.ReadFile(filename)
}

func ReadFileAsYaml(filename string) (out interface{}, err error) {
	data, err := ReadFileAsBytes(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &out)
	return out, err
}
