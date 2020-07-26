package main

/*  4* Внести в телефонный справочник дополнительные данные.
Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных. */

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const filename = "phonebook.json"

type User struct {
	Name     string   `json:"name"`
	LastName string   `json:"last_name"`
	Phones   []string `json:"phones"`
}

func main() {

	// Загружаем данные из файла
	phoneBook := loadPhoneBook()

	// Добавление юзера, для примера
	// TODO сделать заполнение из консоли через os.Args[1:]
	phoneBook = append(phoneBook, User{
		Name:     "Вася",
		LastName: "Пупкин",
		Phones:   []string{"+79123456789", "+79123456780"},
	})

	// Сохраняем в файл
	savePhoneBook(phoneBook)

	for _, user := range phoneBook {
		fmt.Println(user.LastName, user.Name)
		for _, phone := range user.Phones {
			fmt.Printf("\t%v\n", phone)
		}
	}
}

// load phoneBook from file
func loadPhoneBook() []User {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Пустой файл
	if len(data) == 0 {
		data = []byte("[]")
	}

	var phoneBook []User

	err = json.Unmarshal(data, &phoneBook)
	if err != nil {
		log.Fatal(err)
	}
	return phoneBook
}

// save phoneBook to file
func savePhoneBook(phoneBook []User) {
	data, err := json.Marshal(phoneBook)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
