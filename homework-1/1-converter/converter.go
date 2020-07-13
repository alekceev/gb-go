package main

import (
	"fmt"
	"log"
	"strconv"
)

/* 1. Написать программу для конвертации рублей в доллары.
Программа запрашивает сумму в рублях и выводит сумму в долларах.
Курс доллара задайте константой.
*/

const usd float64 = 71.1

func init() {
	fmt.Println("Обмен рублей по курсу ", usd, " за доллар")
}

func main() {
	var str string
	fmt.Println("Введите сумму в рублях:")
	// читаем строку
	fmt.Scanln(&str)

	// конвертим в float
	rub, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatalln(err)
	}

	// округление до центов в пользу банка
	fmt.Printf("Вы получите $%.2f\n", rub/usd-0.005)
}
