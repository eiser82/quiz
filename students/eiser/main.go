package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {

	var filePath string
	flag.StringVar(&filePath, "file", "./problems.csv", "Q&A file path")

	flag.Parse()
	fmt.Println("Welcome to your Quiz.....")
	fmt.Println("Loading questions from: ", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening the file: ", err)
	}

	reader := csv.NewReader(bufio.NewReader(file))

	var exam []Quiz

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error reading content of file: ", err)
		}

		exam = append(exam, Quiz{
			Question: line[0],
			Answer:   line[1],
		})
	}

	fmt.Println(exam)

}
