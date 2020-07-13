package main

import (
	"fmt"
	"log"
	"strconv"
)

const years = 5

func main() {
	fmt.Println("Введите сумму вклада:")
	var str string
	fmt.Scanln(&str)
	amount, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Введите процент вклада:")
	fmt.Scanln(&str)
	percent, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatalln(err)
	}

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
