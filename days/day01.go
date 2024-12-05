package days

import (
	"math"
	"sort"
	"strconv"

	"advent-of-code-2024/utils"
)

func Day01() (int, int) {
	input := utils.ReadInput("inputs/day01.txt")
	lColumn, rColumn := utils.GetColumnsFromInput(input, "   ")

	sort.Strings(lColumn)
	sort.Strings(rColumn)

	return calculateTotalDistance(lColumn, rColumn), calculateSimilarity(lColumn, rColumn)
}

func calculateTotalDistance(lColumn, rColumn []string) int {
	var distance int
	for i := range lColumn {
		l, _ := strconv.Atoi(lColumn[i])
		r, _ := strconv.Atoi(rColumn[i])

		distance += int(math.Abs(float64(l - r)))
	}
	return distance
}

func calculateSimilarity(lColumn, rColumn []string) int {
	var similarity int
	for i := range lColumn {
		l, _ := strconv.Atoi(lColumn[i])

		appearances := 0

		for j := range rColumn {
			r, _ := strconv.Atoi(rColumn[j])
			if l == r {
				appearances++
			}
		}

		similarity += appearances * l
	}
	return similarity
}
