package day8

import (
	"fmt"
	"github.com/soulsteel/aoc-2019-golang/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

func GetResults() {
	input := utils.ReadLines("./input/day8.txt")
	utils.PrintReport(partOne(input), partTwo(input), 8)
}

func partOne(input []string) (counter int) {
	easyDigits := map[int]struct{}{2: struct{}{}, 3: struct{}{}, 4: struct{}{}, 7: struct{}{}}
	for _, v := range input {
		asSlice := strings.Split(v, "|")
		splitted := strings.Split(asSlice[1], " ")
		for _, s := range splitted {
			if _, ok := easyDigits[len(s)]; ok {
				counter++
			}
		}
	}

	return counter
}

func partTwo(input []string) (sum int) {
	for _, l := range input {
		asSlice := strings.Split(l, "|")
		mappers := generateMappers(asSlice[0])
		A := filterCorrectMapper(mappers, asSlice[0])
		sum += mergeLetters(A, asSlice[1])
	}

	return sum
}

func mergeLetters(mapper map[int][]string, outcome string) int {
	splitOutcome := strings.Split(outcome, " ")
	trimmed := make([]string, 0)

	for _, v := range splitOutcome {
		if v != "" {
			trimmed = append(trimmed, v)
		}
	}
	numbers := make([]int, 0)

	for _, word := range trimmed {
		n := getDigit(mapper, word)
		if n == -1 {
			panic("Logic error!")
		}
		numbers = append(numbers, n)
	}

	numberStr := ""
	for _, n := range numbers {
		numberStr += strconv.Itoa(n)
	}

	result, err := strconv.Atoi(numberStr)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}

func generateMappers(income string) []map[int][]string {
	// all the possible ways to align numbers. In total 8
	combinations := make([]map[int]string, 0)
	combinations = append(combinations, map[int]string{})
	combinations = append(combinations, map[int]string{})
	combinations = append(combinations, map[int]string{})
	combinations = append(combinations, map[int]string{})
	combinations = append(combinations, map[int]string{})
	combinations = append(combinations, map[int]string{})
	combinations = append(combinations, map[int]string{})
	combinations = append(combinations, map[int]string{})

	splitIncome := strings.Split(income, " ")

	// Loop through simple digits and populate all the possible combinations
	// digit #1 (contains 2 lines)
	for _, s := range splitIncome {
		if len(s) == 2 {
			fillLetterMap(combinations, s)
		}
	}

	// digit #7 (contains 3 lines)
	for _, s := range splitIncome {
		if len(s) == 3 {
			fillLetterMap(combinations, s)
		}
	}

	// digit #4 (contains 4 lines)
	for _, s := range splitIncome {
		if len(s) == 4 {
			fillLetterMap(combinations, s)
		}
	}

	// digit #8 (contains 7 lines)
	for _, s := range splitIncome {
		if len(s) == 7 {
			fillLetterMap(combinations, s)
		}
	}

	mappers := make([]map[int][]string, 0)
	for _, p := range combinations {
		mappers = append(mappers, generateDigits(p))
	}

	return mappers
}

func filterCorrectMapper(mappers []map[int][]string, income string) map[int][]string {
	trimIncome := make([]string, 0)
	splitIncome := strings.Split(income, " ")
	for _, v := range splitIncome {
		if v != "" {
			trimIncome = append(trimIncome, v)
		}
	}

	correctness := make(map[int]int)
	for i, m := range mappers {
		for _, w := range trimIncome {
			if containsDigit(m, w) {
				correctness[i] += 1
			}
		}
	}

	correctKey := 0
	for i, v := range correctness {
		if v == 10 {
			correctKey = i
		}
	}

	return mappers[correctKey]
}

func fillLetterMap(letterMaps []map[int]string, word string) []map[int]string  {
	splitWord := strings.Split(word, "")

	filled := make(map[string]int)

	for i, v := range letterMaps[0] {
		filled[v] = i
	}

	if len(splitWord) == 2 {
		for i := 0; i < 4; i++ {
			letterMaps[i][2] = splitWord[0]
			letterMaps[i][5] = splitWord[1]
		}

		for i := 4; i < 8; i++ {
			letterMaps[i][2] = splitWord[1]
			letterMaps[i][5] = splitWord[0]
		}
	}

	if len(splitWord) == 3 {
		for _, l := range splitWord {
			_, ok := filled[l]; if !ok {
				for i:=0; i<8;i++ {
					letterMaps[i][1] = l
				}
			}
		}
	}

	if len(splitWord) == 4 {
		notAdded := make([]string, 0)
		for _, l := range splitWord {
			_, ok := filled[l]; if !ok {
				notAdded = append(notAdded, l)
			}
		}
		letterMaps[0][0] = notAdded[0]
		letterMaps[0][3] = notAdded[1]

		letterMaps[1][0] = notAdded[0]
		letterMaps[1][3] = notAdded[1]

		letterMaps[2][0] = notAdded[1]
		letterMaps[2][3] = notAdded[0]

		letterMaps[3][0] = notAdded[1]
		letterMaps[3][3] = notAdded[0]

		letterMaps[4][0] = notAdded[0]
		letterMaps[4][3] = notAdded[1]

		letterMaps[5][0] = notAdded[0]
		letterMaps[5][3] = notAdded[1]

		letterMaps[6][0] = notAdded[1]
		letterMaps[6][3] = notAdded[0]

		letterMaps[7][0] = notAdded[1]
		letterMaps[7][3] = notAdded[0]
	}

	if len(splitWord) == 7 {
		notAdded := make([]string, 0)
		for _, l := range splitWord {
			_, ok := filled[l]; if !ok {
				notAdded = append(notAdded, l)
			}
		}

		for i := 0; i < 8; i++ {
			if i % 2 == 0 {
				letterMaps[i][4] = notAdded[1]
				letterMaps[i][6] = notAdded[0]
			} else {
				letterMaps[i][4] = notAdded[0]
				letterMaps[i][6] = notAdded[1]
			}
		}
	}

	return letterMaps
}

func generateDigits(letters map[int]string) map[int][]string {
	mapper := make(map[int][]string)
	mapper[0] = []string{letters[0], letters[1], letters[2], letters[5], letters[6], letters[4]}				// 6 lines
	mapper[1] = []string{letters[2], letters[5]}															// 2 lines
	mapper[2] = []string{letters[1], letters[2], letters[3], letters[4], letters[6]}							// 5 lines
	mapper[3] = []string{letters[1], letters[2], letters[3], letters[5], letters[6]}							// 5 lines
	mapper[4] = []string{letters[0], letters[3], letters[2], letters[5]}										// 4 lines
	mapper[5] = []string{letters[0], letters[1], letters[3], letters[5], letters[6]}							// 5 lines
	mapper[6] = []string{letters[0], letters[1], letters[3], letters[4], letters[5], letters[6]}				// 6 lines
	mapper[7] = []string{letters[1], letters[2], letters[5]}												// 3 lines
	mapper[8] = []string{letters[0], letters[1], letters[2], letters[3], letters[4], letters[6], letters[5]}	// 7 lines
	mapper[9] = []string{letters[0], letters[1], letters[2], letters[3], letters[5], letters[6]}				// 6 lines

	for _, v := range mapper {
		sort.Strings(v)
	}

	return mapper
}

func getDigit(digitMapper map[int][]string, pattern string) int {
	splitPattern := strings.Split(pattern, "")
	sort.Strings(splitPattern)
	for digit, word := range digitMapper {
		if utils.StringSliceEqual(splitPattern, word) {
			return digit
		}
	}

	fmt.Println(digitMapper, pattern)
	return -1
}

func containsDigit(digitMapper map[int][]string, pattern string) bool {
	splitPattern := strings.Split(pattern, "")
	sort.Strings(splitPattern)
	for _, word := range digitMapper {
		if utils.StringSliceEqual(splitPattern, word) {
			return true
		}
	}
	return false
}
