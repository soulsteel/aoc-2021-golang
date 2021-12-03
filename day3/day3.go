package day3

import (
	"github.com/soulsteel/aoc-2019-golang/utils"
	"log"
	"strconv"
	"strings"
)

func GetResults() (int, int) {
	report := utils.ReadLines("./input/day3small.txt")
	return calcConsumption(report, "oxygen"), calcOxygenAndCO2(report)
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
		//for i, v := range report {
		//	if i == 0 {
		//		continue
		//	}
		//	asSlice := strings.Split(v, "")
		//	bit, err := strconv.Atoi(asSlice[j])
		//	if err != nil {
		//		log.Fatalln(err)
		//	}
		//	bitsCounter[j] += bit
		//}
		bitsCounter[j] = countBits(j, report)
	}

	//for i, v := range report {
	//	if i == 0 {
	//		continue
	//	}
	//
	//	for j := 0; j < len(bitsCounter); j++ {
	//		asSlice := strings.Split(v, "")
	//		bit, err := strconv.Atoi(asSlice[j])
	//		if err != nil {
	//			log.Fatalln(err)
	//		}
	//		bitsCounter[j] += bit
	//	}
	//}

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

func calcOxygenAndCO2(report []string) int {
	innerReport := make([]string, len(report))
	copy(report, innerReport)


	for i := 0; i < len(report[1]); i++ {
		count := countBits(i, innerReport)

		if count >= len(report) / 2 {
			removeBinaries(i, innerReport, 1)
		} else {
			removeBinaries(i, innerReport, 1)
		}
	}

	println("test")
	println(innerReport)

	return 0
}

func countBits(index int, report []string) (counter int) {
	for i, v := range report {
		if i == 0 {
			continue
		}
		asSlice := strings.Split(v, "")
		bit, err := strconv.Atoi(asSlice[index])
		if err != nil {
			log.Fatalln(err)
		}
		counter += bit
	}

	return counter
}

func removeBinaries(index int, report []string, least int) {
	for i, v := range report {
		asSlice := strings.Split(v, "")
		bit, err := strconv.Atoi(asSlice[index])
		if err != nil {
			log.Fatalln(err)
		}
		if bit == least {
			report[i] = report[len(report)-1]
			report = report[:len(report)-1]
		}
	}
}