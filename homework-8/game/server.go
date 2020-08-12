package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// Клиент имеет имя и канал
type client struct {
	Name string
	Ch   chan string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

var (
	answered        bool
	currentQuestion string
	currentAnswer   int
	operators       = []func(a, b int) (int, string){
		func(a, b int) (int, string) {
			return a + b, fmt.Sprintf("%d + %d = ?", a, b)
		},
		func(a, b int) (int, string) {
			return a - b, fmt.Sprintf("%d - %d = ?", a, b)
		},
		func(a, b int) (int, string) {
			return a * b, fmt.Sprintf("%d * %d = ?", a, b)
		},
	}
)

func main() {
	rand.Seed(time.Now().Unix())

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.Ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			if len(clients) == 1 {
				go nextQuestion()
			} else {
				cli.Ch <- fmt.Sprintf("The current question is: %s", currentQuestion)
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Ch)
		}
	}
}

func handleConn(conn net.Conn) {
	cli := client{
		//Name: conn.RemoteAddr().String(),
		Ch: make(chan string),
	}
	go clientWriter(conn, cli.Ch)

	// После подключения запрашиваем имя игрока
	cli.Ch <- "Enter you name: "

	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()
		// Пустую строку не воспринимаем
		if len(text) == 0 {
			continue
		}
		// set name
		if len(cli.Name) == 0 {
			cli.Name = text
			cli.Ch <- "Hi " + cli.Name + "! Let's play"
			messages <- cli.Name + " has arrived"
			entering <- cli
			continue
		}

		answer, err := strconv.Atoi(text)
		if err != nil {
			cli.Ch <- "incorrect input"
			continue
		}
		if answer != currentAnswer {
			cli.Ch <- "incorrect answer"
			continue
		}
		if answered {
			cli.Ch <- "question is already answered"
			continue
		}

		answered = true
		messages <- fmt.Sprintf("%s wins! The answer is: %d", cli.Name, answer)

		go nextQuestion()
	}
	leaving <- cli
	if len(cli.Name) == 0 {
		messages <- cli.Name + " has left"
	}
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func nextQuestion() {
	log.Println("generating next question")

	operator := operators[rand.Intn(len(operators))]

	currentAnswer, currentQuestion = operator(rand.Intn(10), rand.Intn(10))

	answered = false

	messages <- fmt.Sprintf("The next question is: %s", currentQuestion)
}
