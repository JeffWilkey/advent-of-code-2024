package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ConvertStrSliceToInt(columns []string) []int {
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

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to convert string to int: %v", err)
	}
	return i
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
