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
		if levelsAreSafe(levels) || (dampen && dampenedLevelsAreSafe(levels)) {
			safeLevels++
		}
	}
	return safeLevels
}

func dampenedLevelsAreSafe(levels []int) bool {
	for i := range levels {
		newLevels := make([]int, 0, len(levels)-1)
		newLevels = append(newLevels, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)
		if levelsAreSafe(newLevels) {
			return true
		}
	}
	return false
}

func levelsAreSafe(levels []int) bool {
	var isIncreasing bool

	for i := 0; i < len(levels)-1; i++ {
		curr := levels[i]
		next := levels[i+1]

		if curr == next {
			return false
		}

		if i == 0 {
			isIncreasing = curr < next
		}

		// Check if the sequence is not following the increasing/decreasing pattern or the difference is greater than 3
		if (isIncreasing && curr > next) || (!isIncreasing && curr < next) || utils.Abs(curr-next) > 3 {
			return false
		}
	}

	return true
}
