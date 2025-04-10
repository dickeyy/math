package math

import (
	"math"
	"testing"
)

func TestNaN(t *testing.T) {
	got := NaN()
	if !math.IsNaN(got) {
		t.Errorf("NaN() = %v, want NaN", got)
	}
}
