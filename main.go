package main

import (
	"fmt"
	"github.com/soulsteel/aoc-2019-golang/day6"
	"strings"
)

func main() {
	//day1Part1, day1Part2 := day1.GetResults()
	//printReport(day1Part1, day1Part2, 1)
	//
	//day2Part1, day2Part2 := day2.GetResults()
	//printReport(day2Part1, day2Part2, 2)

	//day3Part1, day3Part2 := day3.GetResults()
	//printReport(day3Part1, day3Part2, 3)

	//day4Part1, day4Part2 := day4.GetRes()
	//printReport(day4Part1, day4Part2, 4)

	//day5Part1, day5Part2 := day5.GetResults()
	//printReport(day5Part1, day5Part2, 5)

	day6Part1, day6Part2 := day6.GetResults()
	printReport(day6Part1, day6Part2, 6)

}

func printReport(v1 int, v2 int, day int) {
	fmt.Println(strings.Repeat("~", 30))
	fmt.Printf("Day #%d results.\n", day)
	fmt.Printf("Part 1: %d\n", v1)
	fmt.Printf("Part 2: %d\n", v2)
	fmt.Println(strings.Repeat("~", 30))
}
