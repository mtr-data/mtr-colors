package utils

func ToRgbValue(arg interface{}) uint16 {
	return uint16(arg.(int))
}

func GetRgbArray(arg []interface{}) (out *RgbArray, err error) {
	if len(arg) < 3 {
		return nil, err
	}

	return &RgbArray{
		ToRgbValue(arg[0]),
		ToRgbValue(arg[1]),
		ToRgbValue(arg[2]),
	}, nil
}
