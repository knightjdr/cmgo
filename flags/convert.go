package flags

import "strconv"

// ConvertFloat changes an argument of interface type to a float.
func ConvertFloat(arg interface{}) float64 {
	var converted float64
	if arg != nil {
		converted, _ = strconv.ParseFloat(arg.(string), 64)
	}
	return converted
}

// ConvertString changes an argument of interface type to a string.
func ConvertString(arg interface{}) string {
	var converted string
	if arg != nil {
		converted = arg.(string)
	}
	return converted
}
