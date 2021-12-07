// Package day1
package day1

import (
	"github.com/soulsteel/aoc-2019-golang/utils"
)

func GetResults() {
	nums := utils.GetInts("./input/day1.txt")
	utils.PrintReport(partOne(nums), partTwo(nums), 1)
}

func partOne(nums []int) (counter int) {
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			counter++
		}
	}

	return counter
}

func partTwo(nums []int) (counter int) {
	windows := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(nums) - i == 2 {
			break
		}
		windows = append(windows, nums[i] + nums[i+1] + nums[i+2])			
	}
	
	return partOne(windows)
}
