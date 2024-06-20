package main

import (
	//"bufio"
	"fmt"
	"encoding/csv"
	"log"
	"os"
)

var filepath string = "problems.csv"

func readCsvFile(filepath string) [][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Uneble open file: " + filepath + "\n %v", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("Cant read file as CSV: %v", err)
	}
	return records
}

func main() {
 fmt.Println(readCsvFile(filepath))
}
