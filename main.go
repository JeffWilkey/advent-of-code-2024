package main

import (
	"fmt"

	"advent-of-code-2024/days"
)

func main() {
	fmt.Println("Advent of Code 2024")
	distance, similarities := days.Day01()
	fmt.Printf("Day 01 - Distance: %d, Similarities: %d", distance, similarities)
}
