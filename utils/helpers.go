package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ConvertStrToInt(columns []string) []int {
	ints := make([]int, len(columns))

	for index, column := range columns {
		i, err := strconv.Atoi(column)
		if err != nil {
			log.Fatalf("Failed to convert column to int: %v", err)
		}
		ints[index] = i
	}
	return ints
}

func GetLinesFromInput(input string) []string {
	return strings.Split(input, "\n")
}

func GetColumnsFromInput(input string, delimiter string) [][]string {
	lines := GetLinesFromInput(input)
	var columns [][]string

	for _, line := range lines {
		items := strings.Split(line, delimiter)

		for i, item := range items {
			if len(columns) <= i {
				columns = append(columns, []string{})
			}
			columns[i] = append(columns[i], item)
		}
	}

	return columns
}

func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return string(data)
}
