package main

import (
	"fmt"

	"advent-of-code-2024/days"
)

func main() {
	fmt.Println("Advent of Code 2024")
	distance, similarities := days.Day01()
	fmt.Printf("Day 01 - Distance: %d, Similarities: %d\n", distance, similarities)
	safeLevels, dSafeLevels := days.Day02()
	fmt.Printf("Day 02 - Safe levels: %d, Safe levels(dampened): %d\n", safeLevels, dSafeLevels)
	allMuls, onlyDos := days.Day03()
	fmt.Printf("Day 03 - All Muls: %d, Only do's: %d\n", allMuls, onlyDos)
}
