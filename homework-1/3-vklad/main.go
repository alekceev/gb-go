package main

import (
	"fmt"
)

const years = 5

func main() {
	var amount, percent float64
	fmt.Println("Введите сумму вклада:")
	fmt.Scanln(&amount)

	fmt.Println("Введите процент вклада:")
	fmt.Scanln(&percent)

	if percent > 1 {
		percent /= 100
	}

	fmt.Printf("\n%-10s %-10s\n", "Год", "Сумма")
	for i := 1; i <= years; i++ {
		amount += amount * percent
		fmt.Printf("%-10v %-10.2f\n", i, amount)
	}

	fmt.Printf("\nСумма вклада за %v лет под %v%% годовых составит %.2f\n", years, percent*100, amount)
}
