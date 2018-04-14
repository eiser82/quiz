package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type quiz struct {
	Question string
	Answer   string
}

func main() {

	var filePath string
	var answer string
	var count int
	var exam []quiz

	flag.StringVar(&filePath, "file", "./problems.csv", "Q&A file path")
	flag.Parse()

	// Welcome message
	fmt.Println("Welcome to your Quiz.....")
	fmt.Println("Loading questions from: ", filePath)

	// Opens file from the filesystem and logs
	// an error if something goes wrong
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening the file: ", err)
	}

	// Reads the content of the file and saves it
	// into a quiz type variable
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error reading content of file: ", err)
		}

		exam = append(exam, quiz{
			Question: line[0],
			Answer:   line[1],
		})
	}

	// Prints each question from the quiz and
	// save the answer into a variable
	for _, a := range exam {
		fmt.Printf("%v: ", a.Question)
		fmt.Scan(&answer)

		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)

		if answer == a.Answer {
			count++
		}

	}

	fmt.Printf("Total questions: %v\n", len(exam))
	fmt.Printf("Correct questions: %v\n", count)

}
