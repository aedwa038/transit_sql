package parser

import (
	"path"
	"strings"
)

type Csvdata struct {
	Table   string
	Coloums []string
	Records [][]string
}

func Parse(text []string, file string) Csvdata {

	c := Csvdata{}
	header := parseHeader(text)
	records := parseData(text)
	c.Coloums = header
	c.Records = records
	c.Table = getTable(file)
	return c
}

func parseData(text []string) [][]string {

	var records [][]string
	for i := 1; i < len(text); i++ {
		f := strings.Split(text[i], ",")
		records = append(records, f)
	}
	return records
}

func parseHeader(text []string) []string {
	header := text[0]
	f := strings.Split(header, ",")
	return f
}

func getTable(file string) string {
	f := path.Base(file)
	h := strings.Split(f, ".")
	return h[0]
}
