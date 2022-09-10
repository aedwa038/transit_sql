package converter

import (
	"strings"

	"github.com/aedwa038/transit_sql/parser"
)

var coloumMap = map[string]string{
	"agency_id":           "VARCHAR (255) NOT NULL",
	"agency_name":         "VARCHAR (255) NOT NULL",
	"agency_url":          "VARCHAR (255) NOT NULL",
	"agency_timezone":     "VARCHAR (255) NOT NULL",
	"agency_lang":         "VARCHAR (255) NOT NULL",
	"agency_phone":        "VARCHAR (255)",
	"route_id":            "INT NOT NULL",
	"route_short_name":    "VARCHAR (255)",
	"route_long_name":     "VARCHAR (255) NOT NULL",
	"route_type":          "INT NOT NULL",
	"route_url":           "VARCHAR (255)",
	"route_color":         "VARCHAR (255) NOT NULL",
	"trip_id":             "INT NOT NULL",
	"arrival_time":        "VARCHAR (255) NOT NULL",
	"departure_time":      "VARCHAR (255) NOT NULL",
	"stop_sequence":       "INT NOT NULL",
	"pickup_type":         "INT NOT NULL",
	"drop_off_type":       "INT NOT NULL",
	"shape_dist_traveled": "real",
	"stop_id":             "INT NOT NULL",
	"stop_code":           "INT",
	"stop_name":           "VARCHAR (255) NOT NULL",
	"stop_desc":           "VARCHAR (255)",
	"stop_lat":            "DECIMAL(10,6)",
	"stop_lon":            "DECIMAL(10,6)",
	"zone_id":             "INT NOT NULL",
	"service_id":          "INT NOT NULL",
	"trip_headsign":       "VARCHAR (255) NOT NULL",
	"direction_id":        "INT NOT NULL",
	"block_id":            "VARCHAR (255) NOT NULL",
	"shape_id":            "INT NOT NULL",
	"date":                "INT NOT NULL",
	"exception_type":      "INT NOT NULL",
	"shape_pt_lat":        "DECIMAL(10,6)",
	"shape_pt_lon":        "DECIMAL(10,6)",
	"shape_pt_sequence":   "INT NOT NULL",
}

func GenerateTable(data parser.Csvdata) (string, error) {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE ")
	sb.WriteString("public.")
	sb.WriteString(data.Table)
	sb.WriteString("\n")
	sb.WriteString("(\n")
	for i, col := range data.Coloums {
		sb.WriteString(col)
		sb.WriteString(" ")
		sb.WriteString(coloumMap[col])
		if i < len(data.Coloums)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}
	sb.WriteString(");\n")
	return sb.String(), nil
}

func GenerateInserts(data parser.Csvdata) ([]string, error) {
	var records []string

	for _, r := range data.Records {
		var sb strings.Builder
		sb.WriteString("INSERT INTO ")
		sb.WriteString("public.")
		sb.WriteString(data.Table)
		sb.WriteString("(")
		for i, _ := range r {
			sb.WriteString(data.Coloums[i])
			if i < len(r)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString(") ")
		sb.WriteString("VALUES ")
		sb.WriteString("( ")
		for i, d := range r {
			if strings.Contains(d, "\"") {
				d = strings.ReplaceAll(d, "\"", "'")
			}
			if data.Coloums[i] == "agency_id" && data.Table == "agency" {
				d = "'" + d + "'"
			}
			if data.Coloums[i] == "route_color" && data.Table == "routes" {
				d = "'" + d + "'"
			}
			if data.Coloums[i] == "arrival_time" && data.Table == "stop_times" {
				d = "'" + d + "'"
			}
			if data.Coloums[i] == "departure_time" && data.Table == "stop_times" {
				d = "'" + d + "'"
			}
			if data.Coloums[i] == "stop_desc" && data.Table == "stops" {
				d = "'" + d + "'"
			}
			sb.WriteString(d)
			if i < len(r)-1 {
				sb.WriteString(", ")
			}
		}
		sb.WriteString("); ")
		sb.WriteString("\n")
		records = append(records, sb.String())
	}

	return records, nil
}
