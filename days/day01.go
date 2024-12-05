package days

import (
	"math"
	"sort"

	"advent-of-code-2024/utils"
)

func Day01() (int, int) {
	input := utils.ReadInput("inputs/day01.txt")
	lColumn, rColumn := utils.GetColumnsFromInput(input, "   ")

	sort.Strings(lColumn)
	sort.Strings(rColumn)

	lColumnInts := utils.ConvertStrToInt(lColumn)
	rColumnInts := utils.ConvertStrToInt(rColumn)

	return calculateTotalDistance(lColumnInts, rColumnInts), calculateSimilarity(lColumnInts, rColumnInts)
}

func calculateTotalDistance(lColumn, rColumn []int) int {
	var distance int
	for i := range lColumn {
		distance += int(math.Abs(float64(lColumn[i] - rColumn[i])))
	}
	return distance
}

func calculateSimilarity(lColumn, rColumn []int) int {
	var similarity int
	for _, l := range lColumn {
		appearances := 0

		for _, r := range rColumn {
			if l == r {
				appearances++
			}
		}

		similarity += appearances * l
	}
	return similarity
}
