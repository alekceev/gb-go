package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// 4* Напишите утилиту для копирования файлов, используя пакет flag.

func main() {
	from := flag.String("from", "", "from file")
	to := flag.String("to", "", "to file")

	flag.Parse()

	fmt.Printf("Copy from %s to %s\n", *from, *to)

	// Copy file
	if err := fileCopy(*from, *to); err != nil {
		log.Fatal("Error fileCopy: ", err)
	}
	log.Println("Ok")
}

func fileCopy(from, to string) error {
	fromStat, err := os.Stat(from)
	if err != nil {
		return err
	}

	// only file
	if !fromStat.Mode().IsRegular() {
		return fmt.Errorf("%v not a file", fromStat.Name())
	}

	// error if files exist
	toStat, err := os.Stat(to)
	if err == nil {
		return fmt.Errorf("File %v exists", toStat.Name())
	}

	// open "from" file
	fromFh, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromFh.Close()

	// crate "to" file
	toFh, err := os.Create(to)
	if err != nil {
		return err
	}
	defer toFh.Close()

	// copy data
	_, err = io.Copy(toFh, fromFh)
	if err != nil {
		return fmt.Errorf("Error copy: %v", err)
	}

	return toFh.Sync()
}
