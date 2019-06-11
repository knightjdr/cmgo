package strfunc

import (
	"strings"
)

// BeforeLast returns everything before the last delimiter in a string.
func BeforeLast(str, delim string) string {
	delimIndex := strings.LastIndex(str, delim)
	if delimIndex >= 0 {
		return str[:delimIndex]
	}
	return str
}
