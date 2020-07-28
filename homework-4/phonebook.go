package main

/*
	2. Создать тип, описывающий контакт в телефонной книге.
	   Создать псевдоним типа телефонной книги (массив контактов) и реализовать для него интерфейс Sort{}.
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

const filename = "phonebook.json"

type User struct {
	LastName string   `json:"last_name"`
	Name     string   `json:"name"`
	Phones   []string `json:"phones"`
}

func (u *User) FullName() string {
	return u.LastName + " " + u.Name
}

type PhoneBook []User

func (p *PhoneBook) Len() int {
	return len(*p)
}

func (p *PhoneBook) Less(i, j int) bool {
	return (*p)[i].LastName < (*p)[j].LastName
}

func (p *PhoneBook) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PhoneBook) FindByName(names []string) []int {
	var ids []int

	for i, user := range *p {
		if len(names) > 1 {
			if (user.LastName == names[0] && user.Name == names[1]) || (user.LastName == names[1] && user.Name == names[0]) {
				ids = append(ids, i)
				break
			}
		} else if len(names) == 1 {
			if user.LastName == names[0] || user.Name == names[0] {
				ids = append(ids, i)
			}
		}
	}
	return ids
}

func (p *PhoneBook) AddPhone(row []string) {

	// find ids
	ids := (*p).FindByName(row)
	// if found
	if ids != nil {
		// update first
		(*p)[ids[0]].Phones = row[2:]
		fmt.Println("Updated")
	} else { // or add new
		*p = append(*p, User{
			Name:     row[0],
			LastName: row[1],
			Phones:   row[2:],
		})
		fmt.Println("Added")
	}
}

func main() {

	// Загружаем данные из файла
	phoneBook := loadPhoneBook()

	var command string
	if len(os.Args) < 2 {
		command = "help"
	} else {
		command = os.Args[1]
	}

	switch command {
	case "help":
		printHelp()
	case "list":
		printBook(phoneBook)
	case "add":
		phoneBook.AddPhone(os.Args[2:])
		savePhoneBook(phoneBook)
	case "find":
		ids := phoneBook.FindByName(os.Args[2:])
		if ids == nil {
			fmt.Println("Not found")
		}
		printResult(phoneBook, ids)
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println(`
	commands:
		help - Help
		add  - Add row "add Lastname Name +8123445677"
		find - Find "find Name"
		list - List phonebook
	`)
}

func printResult(phoneBook PhoneBook, ids []int) {
	fmt.Println("Result:")
	for _, i := range ids {
		user := phoneBook[i]
		fmt.Printf("%d) %s\n", i+1, user.FullName())
		for _, phone := range user.Phones {
			fmt.Printf("\t%v\n", phone)
		}
	}
}

func printBook(phoneBook PhoneBook) {
	fmt.Println("PhoneBook:")
	// sort phoneBook
	sort.Sort(&phoneBook)
	for i, user := range phoneBook {
		fmt.Printf("%d) %s\n", i+1, user.FullName())
		for _, phone := range user.Phones {
			fmt.Printf("\t%v\n", phone)
		}
	}
}

// load phoneBook from file
func loadPhoneBook() PhoneBook {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	// Пустой файл
	if len(data) == 0 {
		data = []byte("[]")
	}

	var phoneBook PhoneBook

	err = json.Unmarshal(data, &phoneBook)
	if err != nil {
		log.Fatal(err)
	}
	return phoneBook
}

// save phoneBook to file
func savePhoneBook(phoneBook PhoneBook) {
	data, err := json.Marshal(phoneBook)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
