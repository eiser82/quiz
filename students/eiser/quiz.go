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
	"time"
)

type test struct {
	Question string
	Answer   string
}

type quiz struct {
	test     test
	duration time.Timer
}

func printWelcome(testDuration int, filePath string) {
	// Start test message
	fmt.Println("#######*****###############******#######")
	fmt.Println("Welcome to your Quiz.....")
	fmt.Printf("Duration of the test set to %v seconds...\n", testDuration)
	fmt.Printf("Quiz loaded from: %v...\n", filePath)
	fmt.Println("#######*****###############******#######")
	fmt.Println()
	fmt.Println("Press the [ENTER] key to begin the test.")
	fmt.Println("IMPORTANT: Time will start running after you press ENTER.")
	fmt.Println("Waiting for user...")
	fmt.Println()
	fmt.Println()
}

func createQuiz(filePath string, duration time.Timer) []test {
	var exam []test

	// Opens file from the filesystem
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening the file: ", err)
	}

	// Reads the content of the file and saves it
	// into a quiz type struct
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Error reading content of file: ", err)
		}

		exam = append(exam, test{
			Question: line[0],
			Answer:   line[1],
		})
	}

	return exam
}

func startQuiz(quiz []test) int {
	var count int

	for _, a := range quiz {
		fmt.Printf("%v: ", a.Question)
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)
		if answer == a.Answer {
			count++
		}
	}
	return count
}

func main() {

	var filePath string
	var testDuration int

	flag.StringVar(&filePath, "file", "./problems.csv", "Q&A file path")
	flag.IntVar(&testDuration, "duration", 30, "test duration in seconds")
	flag.Parse()

	printWelcome(testDuration, filePath)
	test := createQuiz(filePath, time.Timer(30*time.Second))
	result := startQuiz(test)

	fmt.Printf("Total questions: %v\n", len(test))
	fmt.Printf("Correct questions: %v\n", result)

}
