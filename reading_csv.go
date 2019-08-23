package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func parseCsv(csvFile *os.File) []map[string]string {
	sheetParsed := make([]map[string]string, 0)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\t'

	headers, e := reader.Read()
	if e == io.EOF {
		return sheetParsed
	}
	if e != nil {
		log.Fatal(e)
	}

	for {
		row_map := make(map[string]string)
		record, e := reader.Read()
		if e == io.EOF {
			break
		}
		if e != nil {
			log.Fatal(e)
		}
		for num, cell := range record {
			row_map[headers[num]] = cell
		}
		sheetParsed = append(sheetParsed, row_map)
	}
	return sheetParsed
}

func main() {
	filepath := "/home/raleigh/Downloads/cleaned-BookstoreList_1.csv"
	csvFile, e := os.Open(filepath)
	if e != nil {
		log.Fatal(e)
	}
	sheetParsed := parseCsv(csvFile)
	fmt.Println(sheetParsed)
}
