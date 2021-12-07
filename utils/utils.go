// Package utils provides shared functionality to read from input files
package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadLines takes filesPath as argument and returns a slice of line contents
func ReadLines(filePath string) (lines []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	return lines
}

// GetInts takes file path and returns a slice of ints for each line of the file
func GetInts(filePath string) (ints []int) {
	for _, v := range ReadLines(filePath)  {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, i)
	}

	return ints
}

func IntDiff(a, b []int) []int {
	mb := make(map[int]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func IntSum(s []int) (result int) {
	for _, v := range s {
		result += v
	}

	return result
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func RemoveIndex(s [][]int, index int) [][]int {
	ret := make([][]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func PrintReport(v1 int, v2 int, day int) {
	fmt.Println(strings.Repeat("~", 30))
	fmt.Printf("Day #%d results.\n", day)
	fmt.Printf("Part 1: %d\n", v1)
	fmt.Printf("Part 2: %d\n", v2)
	fmt.Println(strings.Repeat("~", 30))
}
