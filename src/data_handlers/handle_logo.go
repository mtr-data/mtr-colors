package datahandlers

import (
	"src/utils"
)

const RGB_KEY = "rgb"

func HandleLogo(data interface{}) (out *utils.NameToColorMap, err error) {
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
