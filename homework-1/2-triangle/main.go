package main

import (
	"fmt"
	"log"

	"github.com/alekceev/gb-go/homework-1/2-triangle/geometry"
)

/*
Даны катеты прямоугольного треугольника.
Найти его площадь, периметр и гипотенузу.
Используйте тип данных float64 и функции из пакета math.
*/

func main() {

	a, err := readLeg("Введите катет 'a' прямоугольного треугольника:")
	if err != nil {
		log.Fatalln(err)
	}

	b, err := readLeg("Введите катет 'b' прямоугольного треугольника:")
	if err != nil {
		log.Fatalln(err)
	}

	s := geometry.TriangleArea(a, b)

	fmt.Println("Площадь треугольника равна", s)

	c := geometry.TriangleHypotenuse(a, b)

	fmt.Println("Гипотенуза треугольника равна", c)
	fmt.Println("Периметр треугольника равен", geometry.TrianglePerimeter(a, b, c))
}

func readLeg(str string) (float64, error) {
	fmt.Println(str)
	var f float64
	_, err := fmt.Scanln(&f)
	return f, err
}
