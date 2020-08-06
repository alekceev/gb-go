package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// 4* Напишите функцию для вычисления корней квадратного уравнения (алгоритм можно найти в Википедии) и тесты к ней.

type Coeficient struct {
	A float64
	B float64
	C float64
}

type Result struct {
	X1 float64
	X2 float64
}

func parceInput(i int) float64 {
	val, err := strconv.ParseFloat(os.Args[i], 64)
	if err != nil {
		log.Fatal("Error input:", err)
	}
	return val
}

func main() {

	var c Coeficient

	flag.Float64Var(&c.A, "a", 0, "A int")
	flag.Float64Var(&c.B, "b", 0, "B int")
	flag.Float64Var(&c.C, "c", 0, "C int")
	flag.Parse()

	if c.A == 0 || c.B == 0 || c.C == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	res, err := RootX2(c)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Корни квадратного уравнения: x1=%v x2=%v\n", res.X1, res.X2)
	}
}

func Discriminant(c Coeficient) float64 {
	return c.B*c.B - 4*c.A*c.C
}

func RootX2(c Coeficient) (Result, error) {
	D := Discriminant(c)
	var x Result
	switch {
	case D > 0:
		x.X1 = (-c.B + math.Sqrt(D)) / 2 / c.A
		x.X2 = (-c.B - math.Sqrt(D)) / 2 / c.A
	case D == 0:
		x.X1 = -c.B / 2 / c.A
		x.X2 = x.X1
	case D < 0:
		return x, fmt.Errorf("Квадратное уравнение не имеет вещественных корней")
	}

	return x, nil
}
