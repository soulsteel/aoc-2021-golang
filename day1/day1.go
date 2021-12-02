// Package day1
package day1

import (
	"github.com/soulsteel/aoc-2019-golang/utils"
)

func GetResults() (int, int) {
	nums := utils.GetInts("./input/day1.txt")
	return partOne(nums), partTwo(nums)
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
