package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func basicCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func parseCsv(csvFile *os.File) []map[string]string {
	sheetParsed := make([]map[string]string, 0)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\t'

	headers, err := reader.Read()
	if err == io.EOF {
		return sheetParsed
	}
	basicCheck(err)

	for {
		row_map := make(map[string]string)
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		basicCheck(err)
		for num, cell := range record {
			row_map[headers[num]] = cell
		}
		sheetParsed = append(sheetParsed, row_map)
	}
	return sheetParsed
}

func main() {
	filepath := "/home/raleigh/Downloads/cleaned-BookstoreList_1.csv"
	csvFile, err := os.Open(filepath)
	basicCheck(err)
	sheetParsed := parseCsv(csvFile)
	fmt.Println(sheetParsed)
}
