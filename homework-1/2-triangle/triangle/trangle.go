package triangle

import "math"

// площадь треугольника
func Area(a, h float64) float64 {
	return 0.5 * a * h
}

// гипотенуза треугольника
func Hypotenuse(a, b float64) float64 {
	return math.Hypot(a, b)
}

// периметр треугольника
func Perimeter(a, b, c float64) float64 {
	return a + b + c
}
