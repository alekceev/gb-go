package main

/* 3. Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.
   Числа Фибоначчи определяются соотношениями Fn=Fn-1 + Fn-2. */

import "fmt"

func main() {

	var fibo []float64

	// заполняем ряд
	fibonachi(&fibo, 100)

	for i, n := range fibo {
		fmt.Printf("%-5d %.0f\n", i, n)
	}
}

// фунция получения первых maxLen элементов ряда фибоначи
func fibonachi(fibo *[]float64, maxLen int) {

	// Начальные значения для ряда фибоначи
	if *fibo == nil {
		*fibo = append(*fibo, 0, 1)
	}

	lastID := len(*fibo) - 1

	// Набрали нужную длину, выходим
	if lastID >= 100 {
		return
	}

	n1 := (*fibo)[lastID]
	n2 := (*fibo)[lastID-1]

	*fibo = append(*fibo, n1+n2)

	fibonachi(fibo, maxLen)
}
