package day5

import (
	"fmt"
	"github.com/soulsteel/aoc-2019-golang/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

func GetResults() (int, int) {
	input := utils.ReadLines("./input/day5.txt")
	lines := readLines(input)
	points := getAllPoints(lines)
	countOverlaps := countOverlaps(points)
	fmt.Println(countOverlaps)
	return 0, 0
}

type line struct {
	x1, y1, x2, y2 int
	diff map[int]int
	key string
}

type point struct{
	x, y int
}

func readLines(input []string) (lines []line) {
	for _, row := range input {
		parts := strings.Split(row, "->")
		part1, part2 := strings.Trim(parts[0], " "), strings.Trim(parts[1], " ")
		left, right := strings.Split(part1, ","), strings.Split(part2, ",")
		x1, err := strconv.Atoi(left[0])
		if err != nil {
			log.Fatalln(err)
		}
		y1, err := strconv.Atoi(left[1])
		if err != nil {
			log.Fatalln(err)
		}
		x2, err := strconv.Atoi(right[0])
		if err != nil {
			log.Fatalln(err)
		}
		y2, err := strconv.Atoi(right[1])
		if err != nil {
			log.Fatalln(err)
		}

		lines = append(lines, line{x1:x1, y1:y1, x2:x2, y2:y2})
	}

	return lines
}

func getPointsFromLine(l line) []point {
	points := make([]point, 0)
	if l.x1 == l.x2 {
		min, max := l.y2, l.y1
		if l.y1 < l.y2 {
			min, max = l.y1, l.y2
		}
		for i := min; i<= max; i++ {
			p := point{l.x1, i}
			points = append(points, p)
		}
	}

	if l.y1 == l.y2 {
		min, max := l.x2, l.x1
		if l.x1 < l.x2 {
			min, max = l.x1, l.x2
		}
		for i := min; i<= max; i++ {
			p := point{i, l.y1}
			points = append(points, p)
		}
	}

	if math.Abs(float64(l.x1 - l.x2)) == math.Abs(float64(l.y1 - l.y2)) {
		points = append(points, point{l.x1, l.y1})
		if l.x1 > l.x2  && l.y1 < l.y2 {
			counter := 1
			for i := l.x1; i > l.x2; i-- {
				p := point{l.x1-counter, l.y1+counter}
				points = append(points, p)
				counter++
			}
		}

		if l.x1 > l.x2  && l.y1 > l.y2 {
			counter := 1
			for i := l.x1; i > l.x2; i-- {
				p := point{l.x1-counter, l.y1-counter}
				points = append(points, p)
				counter++
			}
		}

		if l.x1 < l.x2 && l.y1 < l.y2 {
			counter := 1
			for i := l.x1; i < l.x2; i++ {
				p := point{l.x1+counter, l.y1+counter}
				points = append(points, p)
				counter++
			}
		}

		if l.x1 < l.x2 && l.y1 > l.y2 {
			counter := 1
			for i := l.x1; i < l.x2; i++ {
				p := point{l.x1+counter, l.y1-counter}
				points = append(points, p)
				counter++
			}
		}
	}

	return points
}

func getAllPoints(lines []line) []point {
	points := make([]point, 0)

	for _, l := range lines {
		points = append(points, getPointsFromLine(l)...)
	}

	return points
}

func countOverlaps(points []point) (count int) {
	dict := make(map[point]int)
	for _ , p := range points {
		dict[p] = dict[p] +1
	}

	for _, v := range dict {
		if v > 1 {
			count++
		}
	}
	return count
}