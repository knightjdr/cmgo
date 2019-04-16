package flags

// ConvertString changes an argument of interface type to a string.
func ConvertString(arg interface{}) string {
	var converted string
	if arg != nil {
		converted = arg.(string)
	}
	return converted
}
