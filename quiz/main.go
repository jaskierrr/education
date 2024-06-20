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
)

var score int

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
	csvFlag := flag.String("csv", "problems.csv", "help message for csv flag")
	flag.Parse()

	filePath, err := filepath.Abs(*csvFlag)
	if err != nil {
		log.Fatalf("File not found: %v", err)
	}

	records := readCsvFile(filePath)

	for i := range records {
		fmt.Println(records[i][0])

		reader := bufio.NewReader(os.Stdin)
		answear, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Unable read your answear")
		}

		answear = strings.TrimSuffix(answear, "\r\n")

		if answear == records[i][1] {
			score += 1
		}
	}
	fmt.Printf("You answered %d questions out of %d correctly\n", score, len(records))
}
