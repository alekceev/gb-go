package main

/* 3. Дописать функцию, которая будет выводить справку к калькулятору.
Она должна вызываться при вводе слова help вместо выражения.
*/

import (
	"fmt"

	"github.com/alekceev/gb-go/homework-4/3-calc/calculator"
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		if input == "help" {
			printHelp()
			continue
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}

func printHelp() {
	fmt.Printf(`
		Калькурятор.
		Введите выражение для получения результата.
		Доступные бинарные операторы
			%v

		Доступные функции:
			%v

		Доступные константы:
			%v
		`, []string{"+", "-", "*", "/", "%", "&", "|", "^"}, calculator.GetFuncNames(), []string{"pi", "e", "phi"})

	fmt.Printf("\n")

}
