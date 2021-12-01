package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Part 1. Answer: %d\n", partOne())
	fmt.Printf("Part 2. Answer: %d\n", partTwo())
}

func partOne() (counter int) {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	isFirst := true
	var prev int
	
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		next, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		
		if isFirst {
			prev = next
			isFirst = false
		}
		
		if next > prev {
			counter++
		}
		
		prev = next	
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	return counter
}

func partTwo() (counter int) {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums := make([]int, 0)
	
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		next, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, next)	
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	windows := make([]int, 0)
	
	for i := 0; i < len(nums); i++ {
		if len(nums) - i == 2 {
			break
		}
		windows = append(windows, nums[i] + nums[i+1] + nums[i+2])			
	}
	
	for j := 1; j < len(windows); j++ {
		if windows[j] > windows[j-1] {
			counter++
		}
	}
	
	return counter
}
