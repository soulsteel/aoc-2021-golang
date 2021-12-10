package day9

import (
	"fmt"
	"github.com/soulsteel/aoc-2019-golang/utils"
	"log"
	"strconv"
	"strings"
)

func GetResults() {
	input := utils.ReadLines("./input/day9.txt")
	utils.PrintReport(partOne(input),0,9)
}

func partOne(input []string) int {
	lowPoints := 0
	grid, lineLength := makeGrid(input)

	for i:=0;i<len(grid);i++{
		if i == 0 {	// top left
			if grid[i] < grid[i+1] && grid[i] < grid[i+lineLength] {
				lowPoints += grid[i]+1
				//fmt.Println(i, grid[i], "top left")
			}
		} else if i == lineLength-1 { // top right
			if grid[i] < grid[i-1] && grid[i] < grid[i+lineLength] {
				lowPoints += grid[i]+1
				//fmt.Println(i, grid[i], "top right")
			}
		} else if i == len(grid) - lineLength { // bottom left
			if grid[i] < grid[i-lineLength] && grid[i] < grid[i+1] {
				lowPoints += grid[i]+1
				//fmt.Println(i, grid[i], "bottom left")
			}
		} else if i == len(grid) - 1 { // bottom right
			if grid[i] < grid[i-lineLength] && grid[i] < grid[i-1] {
				lowPoints += grid[i]+1
				//fmt.Println(i, grid[i], "bottom right")
			}
		} else if i > 0 && i < lineLength { // top margin
			if grid[i] < grid[i-1] && grid[i] < grid[i+1] && grid[i] < grid[i+lineLength] {
				lowPoints += grid[i]+1
				//fmt.Println(i, grid[i], "top margin")
			}
		}  else if i > len(grid) - lineLength && i < len(grid) - 1 { // bottom margin
			if grid[i] < grid[i-lineLength] && grid[i] < grid[i-1] && grid[i] < grid[i+1] {
				lowPoints += grid[i]+1
				//fmt.Println(i, grid[i], "bottom margin")
			}
		} else if i % lineLength == 0 { // left margin
			if grid[i] < grid[i+1] && grid[i] < grid[i-lineLength]  && grid[i] < grid[i+lineLength] {
				lowPoints += grid[i]+1
				//fmt.Println(i, grid[i], "left margin")
			}
		} else if (i - lineLength + 1) % lineLength == 0 { // right margin
			if grid[i] < grid[i-1] && grid[i] < grid[i-lineLength] && grid[i] < grid[i+lineLength] {
				lowPoints += grid[i]+1
				//fmt.Println(i, grid[i], "right margin")
			}
		} else { // every other point on the grid
			if grid[i] < grid[i+1] && grid[i] < grid[i-1] && grid[i] < grid[i-lineLength] && grid[i] < grid[i+lineLength] {
				lowPoints += grid[i]+1
				//fmt.Println(i, grid[i], "grid")
			}
		}
	}

	fmt.Println(lowPoints)

	return 0
}

func makeGrid(input []string) (A []int, lineLength int) {
	lineLength = len(input[0])
	for _, v := range input {
		splitLine := strings.Split(v, "")
		for _, numberStr := range splitLine {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				log.Fatalln(err)
			}
			A = append(A, number)
		}
	}

	return A, lineLength
}
