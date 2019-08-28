package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
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

func findCsvs(filepath string) []string {
	var csvs []string
	files, e := ioutil.ReadDir(filepath)
	if e != nil {
		log.Fatal(e)
	}
	for _, file := range files {
		fullpath := filepath + "/" + file.Name()
		csvs = append(csvs, fullpath)
	}

	return csvs
}

func readCsv(filepath string, ch chan<- []map[string]string) {
	csvFile, e := os.Open(filepath)
	defer csvFile.Close()
	if e != nil {
		log.Fatal(e)
	}
	ch <- parseCsv(csvFile)
	return
}

func main() {
	ch := make(chan []map[string]string)
	filepath := "/home/raleigh/Desktop/trash/csv_trash"
	csvfiles := findCsvs(filepath)
	for _, file := range csvfiles {
		go readCsv(file, ch)
		fmt.Println(<-ch)
	}
}
