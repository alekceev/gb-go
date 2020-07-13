package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/alekceev/gb-go/homework-1/2-triangle/triangle"
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

	s := triangle.Area(a, b)

	fmt.Println("Площадь треугольника равна", s)

	c := triangle.Hypotenuse(a, b)

	fmt.Println("Гипотенуза треугольника равна", c)
	fmt.Println("Периметр треугольника равен", triangle.Perimeter(a, b, c))
}

func readLeg(str string) (float64, error) {
	fmt.Println(str)
	fmt.Scanln(&str)
	return strconv.ParseFloat(str, 64)
}
