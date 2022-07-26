package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aedwa038/transit_sql/converter"
	"github.com/aedwa038/transit_sql/parser"
)

var root = flag.String("root", "rail_data", "Input file for processing")

func main() {
	flag.Parse()
	var files []string
	err := filepath.Walk(*root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	var data []parser.Csvdata

	for _, file := range files {
		text := readFile(file)
		data = append(data, parser.Parse(text, file))
	}

	for _, c := range data {
		table, err := converter.GenerateTable(c)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(table)
	}

}

func readFile(filename string) []string {

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		t := scanner.Text()
		text = append(text, t)
	}

	return text
}
