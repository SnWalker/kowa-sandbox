// Package calc is a tiny fixture used by Kowa to exercise real GitHub workflows.
package calc

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Clamp limits v to the inclusive range [lo, hi].
func Clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
