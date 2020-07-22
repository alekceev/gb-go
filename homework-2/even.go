package main

import "fmt"

/* 1. Написать функцию, которая определяет, четное ли число. */

func main() {
	var num int
	fmt.Println("Введите число:")
	fmt.Scanln(&num)
	even := "нечётное"
	if isEven(num) {
		even = "чётное"
	}
	fmt.Println(num, even)
}

func isEven(i int) bool {
	return i%2 == 0
}
