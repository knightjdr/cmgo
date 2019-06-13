// Package math defines common math operations
package math

// MinInt returns the minimum of two intergers.
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
