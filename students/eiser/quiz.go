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

var (
	score        int
	filePath     string
	testDuration int
)

type quiz struct {
	question string
	answer   string
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

func createQuiz(filePath string) []quiz {
	var exam []quiz

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

		exam = append(exam, quiz{
			question: line[0],
			answer:   line[1],
		})
	}

	return exam
}

func startQuiz(test []quiz, duration int) {

	for _, a := range test {
		fmt.Printf("%v: ", a.question)
		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)
		if answer == a.answer {
			score++
		}
	}
}

func main() {
	flag.StringVar(&filePath, "file", "./problems.csv", "Q&A file path")
	flag.IntVar(&testDuration, "duration", 30, "test duration in seconds")
	flag.Parse()

	printWelcome(testDuration, filePath)
	test := createQuiz(filePath)

	ticker := time.NewTicker(time.Second * time.Duration(testDuration))
	completed := make(chan bool)

	go func() {
		startQuiz(test, testDuration)
		completed <- true
	}()

	select {
	case <-completed:
	case <-ticker.C:
		fmt.Println("Test completed, time it's up!")
		fmt.Printf("Total questions: %v\n", len(test))
		fmt.Printf("Correct questions: %v\n", score)
	}
}
