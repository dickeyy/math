package math

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		x    float64
		want float64
	}{
		{-1, 1},
		{0, 0},
		{1, 1},
	}
	for _, test := range tests {
		if got := Abs(test.x); got != test.want {
			t.Errorf("Abs(%v) = %v, want %v", test.x, got, test.want)
		}
	}
}

func TestAbsNaN(t *testing.T) {
	got := Abs(NaN())
	if !math.IsNaN(got) {
		t.Errorf("Abs(NaN()) = %v, want NaN", got)
	}
}
