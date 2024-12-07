package days

import (
	"regexp"
	"strings"

	"advent-of-code-2024/utils"
)

func Day03() (int, int) {
	input := utils.ReadInput("inputs/day03.txt")

	return GetMatches(input), GetConditionalMatches(input)
}

func GetMatches(input string) int {
	var total int

	matches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		total += utils.StrToInt(match[1]) * utils.StrToInt(match[2])
	}

	return total
}

func GetConditionalMatches(input string) int {
	var total int;

	dos := strings.Split(input, "do()");

	for _, do := range dos {
		withDontRemoved := strings.Split(do, "don't()")[0]
		total += GetMatches(withDontRemoved)
	}

	return total
}