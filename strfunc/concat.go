// Package strfunc contains functions for handling and modifying strings
package strfunc

import "bytes"

// Concat concatenates an array of strings.
func Concat(arr []string) string {
	var buffer bytes.Buffer
	for _, value := range arr {
		buffer.WriteString(value)
	}
	return buffer.String()
}
