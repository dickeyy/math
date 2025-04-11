package math

import (
	"math"
	"testing"
)

func TestNegativeSqrt(t *testing.T) {
	got := Sqrt(-1)
	if !math.IsNaN(got) {
		t.Errorf("Sqrt(-1) = %v, want NaN", got)
	}
}

func TestZeroSqrt(t *testing.T) {
	got := Sqrt(0)
	if got != 0 {
		t.Errorf("Sqrt(0) = %v, want 0", got)
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		x    float64
		want float64
	}{
		{4, 2},
	}
	for _, test := range tests {
		if got := Sqrt(test.x); got != test.want {
			t.Errorf("Sqrt(%v) = %v, want %v", test.x, got, test.want)
		}
	}
}

func TestSqrtCompareMath(t *testing.T) {
	ours := Sqrt(24)
	theirs := math.Sqrt(24)
	if ours != theirs {
		t.Errorf("Sqrt(24) = %v, want %v", ours, theirs)
	}
}
