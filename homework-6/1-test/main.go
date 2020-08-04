package main

import (
	"fmt"

	"github.com/alekceev/gb-go/homework-6/1-test/statistic"
)

// 1. Дополните код из раздела «Тестирование» функцией подсчета суммы переданных элементов и тестом для этой функции.

func main() {
	fmt.Println("avg:", statistic.Average([]float64{2, 4, 2, 4}))
}
