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

var (
	schemaFile  = "scripts/schema.sql"
	queriesFile = "scripts/queries.sql"
	transitFile = "scripts/transit.sql"
)

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

	var tables []string
	for _, c := range data {
		table, err := converter.GenerateTable(c)
		if err != nil {
			fmt.Println(err)
		}
		tables = append(tables, table)
	}

	var inserts []string
	for _, c := range data {
		tableInserts, err := converter.GenerateInserts(c)
		if err != nil {
			fmt.Println(err)
		}
		for _, ins := range tableInserts {
			inserts = append(inserts, ins)
		}
	}
	//Get the base file dir
	path, err := os.Getwd()
	if err != nil {
		log.Println("error msg", err)
	}

	//Create output path
	outPath := filepath.Join(path, "scripts")
	if _, err := os.Stat(outPath); !os.IsNotExist(err) {
		err := os.RemoveAll(outPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := writeFile("schema.sql", tables); err != nil {
		fmt.Println(err)
	}

	if err := writeFile("queries.sql", inserts); err != nil {
		fmt.Println(err)
	}

	if err := writeFile("transit.sql", append(tables, inserts...)); err != nil {
		fmt.Println(err)
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

func writeFile(filename string, data []string) error {
	f, err := os.Create("scripts/" + filename)
	if err != nil {
		f.Close()
		return err
	}

	for _, line := range data {
		if _, err := fmt.Fprintln(f, line); err != nil {
			return err
		}
	}
	defer f.Close()
	return nil

}
