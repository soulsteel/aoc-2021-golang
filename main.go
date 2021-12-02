package main

import (
	"fmt"
	"github.com/soulsteel/aoc-2019-golang/day1"
	"github.com/soulsteel/aoc-2019-golang/day2"
	"strings"
)

func main() {
	day1Part1, day1Part2 := day1.GetResults()
	printReport(day1Part1, day1Part2, 1)

	day2Part1, day2Part2 := day2.GetResults()
	printReport(day2Part1, day2Part2, 2)
}

func printReport(v1 int, v2 int, day int) {
	fmt.Println(strings.Repeat("~", 30))
	fmt.Printf("Day #%d results.\n", day)
	fmt.Printf("Part 1: %d\n", v1)
	fmt.Printf("Part 2: %d\n", v2)
	fmt.Println(strings.Repeat("~", 30))
}
