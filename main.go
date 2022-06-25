package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aedwa038/transit_sql/parser"
)

var file = flag.String("file", "input.txt", "Input file for processing")
var url = flag.String("url", "https://www.nowinstock.net/videogaming/consoles/sonyps5/full_history.php", "url input")

func main() {
	flag.Parse()
	text := readFile(*file)
	c := parser.Parse(text, *file)
	fmt.Println(c)
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
