// Package rbo calculates the rank biased overlap between ordered lists
package rbo

// RBDext calculates the RBOext to depth k and returns the distance (1-RBOext).
// k should not be greater than max(len(s), len(t)), and will be set to this
// minimum if it is. Set k to 0 to comparse full length of lists.
// p is the persitence (weighting).
func RBDext(S, T []string, p float64, userK int) float64 {
	return 1 - RBOext(S, T, p, userK)
}
