package day7

import (
	"github.com/soulsteel/aoc-2019-golang/utils"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func GetResults() {
	input := utils.ReadLines("./input/day7.txt")
	splitted := strings.Split(input[0], ",")
	positions := make([]int, 0)
	for _, v := range splitted {
		numb, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		positions = append(positions, numb)
	}

	utils.PrintReport(partOne(positions), partTwo(positions), 7)
}

func partOne(positions []int) int {
	sort.Ints(positions)

	medians := make([]int, 0)
	if len(positions) % 2 == 0 {
		left := len(positions)/2
		right := len(positions)/2-1
		medians = append(medians, left, right)
	} else {
		medians = append(medians, (len(positions)-1)/2)
	}

	distancesSums := make([]int, 0)
	for _, m := range medians {
		distances := make([]int, 0)
		for _, v := range positions {
			d := int(math.Abs(float64(v - positions[m])))
			distances = append(distances, d)
		}
		distancesSums = append(distancesSums, utils.IntSum(distances))
	}

	min, _ := utils.MinMax(distancesSums)

	return min
}

func partTwo(positions []int) int {
	floor := int(math.Floor(float64(utils.IntSum(positions)) / float64(len(positions))))
	ceil := int(math.Ceil(float64(utils.IntSum(positions)) / float64(len(positions))))

	min, _ := utils.MinMax([]int{
		calculateCrabFuel(positions, floor),
		calculateCrabFuel(positions, ceil)})
	return min
}

func calculateCrabFuel(positions []int, point int) int {
	cumulative := make([]int, 0)
	for _, p := range positions {
		n := int(math.Abs(float64(p)-float64(point)))
		total := n * (n + 1) / 2
		cumulative = append(cumulative, total)
	}

	return utils.IntSum(cumulative)
}