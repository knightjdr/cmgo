package flags

// SetFloat sets the field value from a command line argument, then from the fileOptions
// map if no argument is defined and finally to the defaultValue.
func SetFloat(field string, args, fileOptions map[string]interface{}, defaultValue float64) float64 {
	var value float64

	if args[field] != nil {
		value = convertFloat(args[field])
	} else if fileOptions[field] != nil {
		value = convertFloat(fileOptions[field])
	} else {
		value = defaultValue
	}

	return value
}

// SetInt sets the field value from a command line argument, then from the fileOptions
// map if no argument is defined and finally to the defaultValue.
func SetInt(field string, args, fileOptions map[string]interface{}, defaultValue int) int {
	var value int

	if args[field] != nil {
		value = convertInt(args[field])
	} else if fileOptions[field] != nil {
		value = convertInt(fileOptions[field])
	} else {
		value = defaultValue
	}

	return value
}

// SetString sets the field value from a command line argument, then from the fileOptions
// map if no argument is defined and finally to the defaultValue.
func SetString(field string, args, fileOptions map[string]interface{}, defaultValue string) string {
	var value string

	if args[field] != nil {
		value = convertString(args[field])
	} else if fileOptions[field] != nil {
		value = convertString(fileOptions[field])
	} else {
		value = defaultValue
	}

	return value
}
