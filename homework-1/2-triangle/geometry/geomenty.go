package geometry

import "math"

// площадь треугольника
func TriangleArea(a, h float64) float64 {
	return 0.5 * a * h
}

// гипотенуза треугольника
func TriangleHypotenuse(a, b float64) float64 {
	return math.Hypot(a, b)
}

// периметр треугольника
func TrianglePerimeter(a, b, c float64) float64 {
	return a + b + c
}
