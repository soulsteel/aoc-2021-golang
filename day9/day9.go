package day9

import (
	"github.com/soulsteel/aoc-2019-golang/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

func GetResults() {
	input := utils.ReadLines("./input/day9.txt")
	grid := createGrid(input)
	lowPoints := findLowPoints(grid)

	utils.PrintReport(partOne(grid, lowPoints), partTwo(grid, lowPoints), 9)
}

func partOne(grid [][]int, lowPoints [][]int) (count int) {
	for _, coords := range lowPoints {
		count += grid[coords[0]][coords[1]] + 1
	}
	return count
}

func partTwo(grid [][]int, lowPoints [][]int) int {
	visited := make([][]bool, 0)
	for i:=0;i<len(grid);i++ {
		row := make([]bool, 0)
		for j:=0;j<len(grid[0]);j++ {
			row = append(row, false)
		}
		visited = append(visited, row)
	}

	adjX := []int{-1, 0, 1, 0}
	adjY := []int{0, 1, 0, -1}
	largeBasins := make([]int, 0)

	for _, v := range lowPoints {
		queue := make([][]int, 0)
		queue = append(queue, v)
		visited[v[0]][v[1]] = true
		basinCounter := 1

		for len(queue) > 0 {
			firstOut := queue[0]
			queue = queue[1:]

			for k:=0; k<4; k++ {
				adjI, adjJ := firstOut[0] + adjX[k], firstOut[1] + adjY[k]

				if adjI >= 0 && adjI < len(grid) &&
					adjJ >= 0 && adjJ < len(grid[0]) &&
					!visited[adjI][adjJ] &&
					grid[adjI][adjJ] != 9 {
					queue = append(queue, []int{adjI, adjJ})
					visited[adjI][adjJ] = true
					basinCounter++
				}
			}
		}

		if len(largeBasins) == 3 {
			sort.Ints(largeBasins)
			if largeBasins[0] < basinCounter {
				largeBasins[0] = basinCounter
			}
		} else {
			largeBasins = append(largeBasins, basinCounter)
		}
	}

	return largeBasins[0] * largeBasins[1] * largeBasins[2]
}

func createGrid(input []string) [][]int {
	grid := make([][]int, 0)

	for _, line := range input {
		row := strings.Split(line, "")
		numbers := make([]int, 0)
		for _, numberStr := range row {
			n, err := strconv.Atoi(numberStr)
			if err != nil {
				log.Fatalln(err)
			}
			numbers = append(numbers, n)
		}
		grid = append(grid, numbers)
	}

	return grid
}

func findLowPoints(grid [][]int) [][]int {
	adjX := []int{-1, 0, 1, 0}
	adjY := []int{0, 1, 0, -1}

	lowPoints := make([][]int, 0)

	for i:=0;i<len(grid);i++ {
		for j:=0;j<len(grid[0]);j++ {
			isLowPoint := true
			for k:=0;k<4;k++ {
				adjI, adjJ := i + adjX[k], j + adjY[k]
				if adjI >= 0 && adjI < len(grid) && adjJ >= 0 && adjJ < len(grid[0]) {
					if grid[adjI][adjJ] <= grid[i][j] {
						isLowPoint = false
					}
				}
			}
			if isLowPoint {
				lowPoints = append(lowPoints, []int{i, j})
			}
		}
	}

	return lowPoints
}
