package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ConvertStrToInt(columns []string) []int {
	ints := make([]int, len(columns))

	for _, column := range columns {
		i, err := strconv.Atoi(column)
		if err != nil {
			log.Fatalf("Failed to convert column to int: %v", err)
		}
		ints = append(ints, i)
	}
	return ints
}

func GetColumnsFromInput(input string, delimeter string) ([]string, []string) {
	// Split input by new line
	var lColumn []string
	var rColumn []string

	lines := strings.Split(input, "\n")

	// Split each line by space
	for _, line := range lines {
		columns := strings.Split(line, delimeter)

		lColumn = append(lColumn, columns[0])
		rColumn = append(rColumn, columns[1])
	}

	return lColumn, rColumn
}

func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return string(data)
}
