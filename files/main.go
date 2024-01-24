package main

import (
	"log"
	"os"
)

func WriteToFile() {
	file, err := os.Create("text.txt")
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("test\nhello")
}

func ReadFromFile() {
	b, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
}

func main() {
	WriteToFile()
	ReadFromFile()
}
