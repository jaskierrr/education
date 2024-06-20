package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)


func readCsvFile(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Unable open file: " + filepath + "\n %v", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Unable read file as CSV: %v", err)
	}
	return records
}


func main() {
	csvPath := flag.String("csv", "problems.csv", "choose file with quiz data")
	timeLimit := flag.Int("limit", 30, "set time limit to quiz in seconds")
	flag.Parse()

	filePath, err := filepath.Abs(*csvPath)
	if err != nil {
		log.Fatalf("File not found: %v", err)
	}

	records := readCsvFile(filePath)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	score := 0

	labelloop:
	for i := range records {
		fmt.Println(records[i][0])

		answerCh := make(chan string)

		go func () {
			reader := bufio.NewReader(os.Stdin)
			answer, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Unable read your answer")
			}
			answerCh <- strings.TrimSuffix(answer, "\r\n")
		}()

		select {
			case <- timer.C:
				fmt.Printf("You answered %d questions out of %d correctly\n", score, len(records))
				break labelloop
			case answer := <- answerCh:
				if answer == records[i][1] {
					score++
				}
		}
	}
}
