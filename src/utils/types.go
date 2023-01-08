package utils

// string to any
type StringToAnyMap = map[string]interface{}

// RGB
type RgbArray = [3]uint8
type NameToColorMap = map[string]RgbArray

// color handler functions
type ColorHandler func(data interface{}) (out *NameToColorMap, err error)
