package datahandlers

import (
	"src/utils"
)

const RGB_KEY = "rgb"

func HandleLogo(filename string) (out *utils.NameToColorMap, err error) {
	data, err := utils.ReadFileAsYaml(filename)
	if err != nil {
		return nil, err
	}

	result := make(utils.NameToColorMap)

	for colorName, colorObject := range data.(utils.StringToAnyMap) {
		rgb := colorObject.(utils.StringToAnyMap)[RGB_KEY].([]interface{})
		rgbArray, err := utils.GetRgbArray(rgb)
		if err != nil {
			return nil, err
		}

		result[colorName] = *rgbArray
	}
	return &result, err
}
