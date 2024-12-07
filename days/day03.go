package days

import (
	"fmt"
	"regexp"
	"strings"

	"advent-of-code-2024/utils"
)

func Day03() (int, int) {
	input := utils.ReadInput("inputs/day03.txt")
	lines := utils.GetLinesFromInput(input)

	return CalculateMulFromCorruptedData(lines), CalculateMulFromCorruptedData2(lines)
}

func CalculateMulFromCorruptedData(lines []string) int {
	var total int

	for _, line := range lines {
		matches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			total += utils.StrToInt(match[1]) * utils.StrToInt(match[2])
		}
	}
	return total
}

func CalculateMulFromCorruptedData2(lines []string) int {
	var total int

	for _, line := range lines {
		do := strings.Split(line, "do()")

		// loop over every other element
		fmt.Println(do)
		for i := 0; i < len(do); i += 2 {
			matches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(do[i], -1)
			for _, match := range matches {
				total += utils.StrToInt(match[1]) * utils.StrToInt(match[2])
			}
		}
	}
	return total
}
