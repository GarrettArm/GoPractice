package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func parseCsv() []map[string]string {
	sheetParsed := make([]map[string]string, 0)
	csvFile, _ := os.Open("/home/raleigh/Downloads/cleaned-BookstoreList_1.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\t'
	headers, err := reader.Read()
	if err == io.EOF {
		return sheetParsed
	}
	if err != nil {
		log.Fatal(err)
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		row_map := make(map[string]string)
		for num, cell := range record {
			row_map[headers[num]] = cell
		}
		sheetParsed = append(sheetParsed, row_map)
	}
	return sheetParsed
}

func main() {
	sheetParsed := parseCsv()
	fmt.Println(sheetParsed)
}
