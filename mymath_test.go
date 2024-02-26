package mymath

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{1, 1},
		{-1, 1},
		{2.5, 2.5},
		{-2.5, 2.5},
	}

	for _, tc := range testCases {
		result := Abs(tc.input)
		if result != tc.expected {
			t.Errorf("Abs(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}

func TestCeil(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{1.2, 2},
		{-3.8, -3},
		{5.5, 6},
	}

	for _, tc := range testCases {
		result := Ceil(tc.input)
		if result != tc.expected {
			t.Errorf("Ceil(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}

// По аналогии добавьте тесты для остальных функций

func TestFloor(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{1.6, 1},
		{-2.4, -3},
		{3.8, 3},
	}

	for _, tc := range testCases {
		result := Floor(tc.input)
		if result != tc.expected {
			t.Errorf("Floor(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		x        float64
		y        float64
		expected float64
	}{
		{0, 0, 0},
		{1.5, -2.5, 1.5},
		{3, 3, 3},
		{-4, -1, -1},
	}

	for _, tc := range testCases {
		result := Max(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("Max(%v, %v) expected: %v, got: %v", tc.x, tc.y, tc.expected, result)
		}
	}
}
func TestMin(t *testing.T) {
	testCases := []struct {
		x        float64
		y        float64
		expected float64
	}{
		{0, 0, 0},
		{1.5, -2.5, -2.5},
		{3, 3, 3},
		{-4, -1, -4},
	}

	for _, tc := range testCases {
		result := Min(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("Min(%v, %v) expected: %v, got: %v", tc.x, tc.y, tc.expected, result)
		}
	}
}

func TestSin(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{math.Pi / 2, 1},
		{-math.Pi / 2, -1},
		{math.Pi, 0},
	}

	for _, tc := range testCases {
		result := Sin(tc.input)
		if result != tc.expected {
			t.Errorf("Sin(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}

func TestCos(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 1},
		{math.Pi / 2, 0},
		{-math.Pi / 2, 0},
		{math.Pi, -1},
	}

	for _, tc := range testCases {
		result := Cos(tc.input)
		if result != tc.expected {
			t.Errorf("Cos(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}
func TestTan(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{math.Pi / 4, 1},
		{math.Pi / 2, math.Inf(1)},
		{-math.Pi / 4, -1},
	}

	for _, tc := range testCases {
		result := Tan(tc.input)
		if result != tc.expected {
			t.Errorf("Tan(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}

func TestExp(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 1},
		{1, math.E},
		{2, math.E * math.E},
		{-1, 1 / math.E},
	}

	for _, tc := range testCases {
		result := Exp(tc.input)
		if result != tc.expected {
			t.Errorf("Exp(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}

func TestLog(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{1, 0},
		{math.E, 1},
		{100, math.Log(100)},
	}

	for _, tc := range testCases {
		result := Log(tc.input)
		if result != tc.expected {
			t.Errorf("Log(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}

// Продолжайте добавлять тесты для остальных функций

func TestLog10(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{1, 0},
		{10, 1},
		{100, 2},
	}

	for _, tc := range testCases {
		result := Log10(tc.input)
		if result != tc.expected {
			t.Errorf("Log10(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}

func TestPow(t *testing.T) {
	testCases := []struct {
		x        float64
		y        float64
		expected float64
	}{
		{2, 3, 8},
		{10, 0, 1},
		{3, 0.5, math.Sqrt(3)},
	}

	for _, tc := range testCases {
		result := Pow(tc.x, tc.y)
		if result != tc.expected {
			t.Errorf("Pow(%v, %v) expected: %v, got: %v", tc.x, tc.y, tc.expected, result)
		}
	}
}

func TestSqrt(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{4, 2},
		{16, 4},
	}

	for _, tc := range testCases {
		result := Sqrt(tc.input)
		if result != tc.expected {
			t.Errorf("Sqrt(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}

func TestCbrt(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
	}{
		{0, 0},
		{8, 2},
		{27, 3},
	}

	for _, tc := range testCases {
		result := Cbrt(tc.input)
		if result != tc.expected {
			t.Errorf("Cbrt(%v) expected: %v, got: %v", tc.input, tc.expected, result)
		}
	}
}
