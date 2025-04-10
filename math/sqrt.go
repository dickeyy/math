package math

// Sqrt returns the square root of x.
//
// It implements the Babylonian method, or Heron's method:
//
//	z -= (z*z - x) / (2 * z)
//
// The initial value z is chosen to be 1, which is not a good guess.
// However, it converges to the correct result.
func Sqrt(x float64) float64 {
	if x < 0 {
		return NaN()
	}
	if x == 0 {
		return 0
	}

	// Use the Babylonian method to approximate the square root.
	z := 1.0
	for range 1000 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
