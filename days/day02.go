package days

import (
	"strings"

	"advent-of-code-2024/utils"
)

func Day02() (int, int) {
	input := utils.ReadInput("inputs/day02.txt")
	lines := utils.GetLinesFromInput(input)

	return calculateSafeLevels(lines, false), calculateSafeLevels(lines, true)
}

func calculateSafeLevels(lines []string, dampen bool) int {
	var safeLevels int

	for _, line := range lines {
		levels := utils.ConvertStrToInt(strings.Split(line, " "))
		if levelsAreSafe(levels) {
			safeLevels++
		} else if dampen && dampenedLevelsAreSafe(levels) {
			safeLevels++
		}
	}
	return safeLevels
}

func dampenedLevelsAreSafe(levels []int) bool {
	for i := range levels {
		newLevels := append([]int(nil), levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)
		if levelsAreSafe(newLevels) {
			return true
		}
	}
	return false
}

func levelsAreSafe(levels []int) bool {
	var isIncreasing bool

	for i := range len(levels) - 1 {
		curr := levels[i]
		next := levels[i+1]

		if curr == next {
			return false
		}

		if i == 0 {
			if curr < next {
				isIncreasing = true
			} else {
				isIncreasing = false
			}
		}

		if isIncreasing && curr > next || next-curr > 3 {
			return false
		}

		if !isIncreasing && curr < next || curr-next > 3 {
			return false
		}
	}

	return true
}
