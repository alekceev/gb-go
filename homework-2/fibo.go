package main

/* 3. Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.
   Числа Фибоначчи определяются соотношениями Fn=Fn-1 + Fn-2. */

import "fmt"

func main() {

	fibo := fibonachi(100)

	for i, n := range fibo {
		fmt.Printf("%-5d %.0f\n", i, n)
	}
}

// фунция получения первых maxLen элементов ряда фибоначи
func fibonachi(maxLen int) []float64 {

	// Начальные значения для ряда фибоначи
	var fibo = []float64{0, 1}

	for i := 1; i < maxLen; i++ {
		fibo = append(fibo, fibo[i]+fibo[i-1])
	}

	return fibo
}
