package mymath

import "math"

// Abs возвращает модуль числа
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Ceil возвращает наименьшее целое число, не менее x
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Floor возвращает наибольшее целое число, не более x
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Max возвращает максимальное из двух чисел
func Max(x, y float64) float64 {
	return math.Max(x, y)
}

// Min возвращает минимальное из двух чисел
func Min(x, y float64) float64 {
	return math.Min(x, y)
}

// Sin возвращает синус угла x (в радианах)
func Sin(x float64) float64 {
	return math.Sin(x)
}

// Cos возвращает косинус угла x (в радианах)
func Cos(x float64) float64 {
	return math.Cos(x)
}

// Tan возвращает тангенс угла x (в радианах)
func Tan(x float64) float64 {
	return math.Tan(x)
}

// Exp возвращает экспоненту, e^x
func Exp(x float64) float64 {
	return math.Exp(x)
}

// Log возвращает натуральный логарифм x
func Log(x float64) float64 {
	return math.Log(x)
}

// Log10 возвращает десятичный логарифм x
func Log10(x float64) float64 {
	return math.Log10(x)
}

// Pow возвращает x в степени y
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Sqrt возвращает квадратный корень из x
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Cbrt возвращает кубический корень из x
func Cbrt(x float64) float64 {
	return math.Cbrt(x)
}
