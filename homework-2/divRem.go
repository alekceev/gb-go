package main

import "fmt"

/* 2. Написать функцию, которая определяет, делится ли число без остатка на 3. */

const div = 3

func main() {
	var num int
	fmt.Println("Введите число:")
	fmt.Scanln(&num)
	res := " не"
	if isDivisibleBy(num, div) {
		res = ""
	}
	fmt.Printf("Число %d%s делится без остатся на %d\n", num, res, div)
}

func isDivisibleBy(dividend, divider int) bool {
	return dividend%divider == 0
}
