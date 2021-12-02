// Package utils provides shared functionality to read from input files
package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
