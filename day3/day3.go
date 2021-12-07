package day3

import (
	"errors"
	"github.com/soulsteel/aoc-2019-golang/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

func GetResults() {
	report := utils.ReadLines("./input/day3.txt")
	calcOxygenAndCO2(report, 0, "co2")
	utils.PrintReport(partOne(report), 0, 3)

}

func calcConsumption(report []string, rating string) int {
	firstRow := strings.Split(report[0], "")
	bitsCounter := make([]int, 0)

	for _, v := range firstRow {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		bitsCounter = append(bitsCounter, n)
	}

	for j := 0; j < len(bitsCounter); j++ {
		bitsCounter[j] = countBits(j, report)
	}

	gammaAsStr := ""
	for _, v := range bitsCounter {
		if (v >= len(report) / 2) && (rating == "oxygen") {
			gammaAsStr += "1"
		} else {
			gammaAsStr += "0"
		}
	}

	gamma, err := strconv.ParseUint(gammaAsStr, 2, 16)
	if err != nil {
		log.Fatalln(err)
	}

	epsilon := ^gamma
	mask := uint64(0b111111111111)
	consumption := (gamma & mask) * (epsilon & mask)

	return int(consumption)
}

func calcOxygenAndCO2(report []string, index int, elem string) []string {
	if len(report) == 1 {
		return report
	}

	count := countBits(index, report)
	reportLen := int(math.Round(float64(len(report)) / 2.0))

	if count == reportLen && elem == "oxygen" {
		l := removeBinaries(index, report, 1)
		return calcOxygenAndCO2(l, index+1, elem)
	}

	if count == reportLen && elem == "co2" {
		l := removeBinaries(index, report, 0)
		return calcOxygenAndCO2(l, index+1, elem)
	}

	if count < reportLen {
		l := removeBinaries(index, report, 1)
		return calcOxygenAndCO2(l, index+1, elem)
	} else {
		l := removeBinaries(index, report, 0)
		return calcOxygenAndCO2(l, index+1, elem)
	}
}

func countBits(index int, report []string) (counter int) {
	for _, v := range report {
		asSlice := strings.Split(v, "")
		bit, err := strconv.Atoi(asSlice[index])
		if err != nil {
			log.Fatalln(err)
		}
		counter += bit
	}

	return counter
}

func removeBinaries(index int, report []string, least int) []string {
	c := report[:0]
	for _, n := range report {
		asSlice := strings.Split(n, "")
		bit, err := strconv.Atoi(asSlice[index])
		if err != nil {
			log.Fatalln(err)
		}
		if bit == least {
			c = append(c, n)
		}
	}

	return c
}

func countCommonBeat(arr []string, col int, side string) (int, error) {
	bits := make([]bool, 0)
	for _, v := range arr {
		asSlice := strings.Split(v, "")
		n, err := strconv.Atoi(asSlice[col])
		if err != nil {
			log.Fatalln(err)
		}
		switch side {
		case "least":
			bits = append(bits, n != 1)
		case "most":
			bits = append(bits, n == 1)
		default:
			return -1, errors.New("unknown argument: side provided. Accepted ones: `most`, `least`")
		}
	}

	counter := 0
	for _, v := range bits {
		if v {
			counter++
		}
	}

	return counter, nil
}

func partOne(report []string) int {
	result := ""
	for i:=0; i<len(report[0]);i++ {
		counter, err := countCommonBeat(report, i, "most")
		if err != nil {
			log.Fatalln(err)
		}
		result += strconv.Itoa(counter)
	}

	n, err := strconv.Atoi(result)
	if err != nil {
		log.Fatalln(err)
	}

	return n
}