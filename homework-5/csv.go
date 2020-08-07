package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

// 3. Самостоятельно изучите пакет encoding/csv. Напишите пример, иллюстрирующий его применение.

func main() {

	phones := [][]string{
		{"name", "phone"},
		{"Вася Пупкин", "+79123456688"},
		{"Коля Моржин", "+791122333454"},
		{"Витя Скалкин", "+79333123123"},
	}

	// WRITE
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	w.Comma = ';'

	w.WriteAll(phones)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	s := buf.String()
	log.Printf("%v\n", s)

	// READ
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = ';'

	phones, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(phones)
	for _, p := range phones {
		fmt.Println(p[0], p[1])
	}

}
