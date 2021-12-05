package day4

import (
	"fmt"
	"github.com/soulsteel/aoc-2019-golang/utils"
	"log"
	"strconv"
	"strings"
)

func GetRes() (int, int) {
	raw := utils.ReadLines("./input/day4.txt")
	boards := genBoardNums(raw)

	//fmt.Println(calcMult(boards))

	calcLastMult(boards)

	return 0, 0
}

func genBoardNums(input []string) [][]int {
	boards := make([][]int, 0)
	for i := 0; i <= len(input) - 5; i += 5 {
		nums := make([]int, 0)

		for j := i; j < (i+5); j++ {
			asStrs := strings.Split(input[j], " ")

			for _, v := range asStrs {
				if v == "" {
					continue
				}

				n, err := strconv.Atoi(v)
				if err != nil {
					log.Fatalln(err)
				}
				nums = append(nums, n)
			}
		}
		boards = append(boards, nums)
	}

	return boards
}

func getMarkedIndexes(board []int, seq []int) []int {
	markedIndexes := make([]int, 0)
	for index, num := range board {
		for _, M := range seq {
			if num == M {
				markedIndexes = append(markedIndexes, index)
			}
		}
	}

	return markedIndexes
}

func isBingo(markedIndexes []int, d string) bool {
	limit, step, next := 21, 5, 1
	if d == "column" {
		limit, step, next = 5, 1, 5
	}

	for i := 0; i < limit; i += step {
		bingo := []int{i, i+1*next, i+2*next, i+3*next, i+4*next}
		diff := utils.IntDiff(markedIndexes, bingo)
		if len(markedIndexes) - len(diff) == 5 {
			return true
		}
	}

	return false
}

func calcMult(boards [][]int) int {
	data := []int{50,68,2,1,69,32,87,10,31,21,78,23,62,98,16,99,65,35,27,96,66,26,74,72,45,52,81,60,38,57,54,19,18,77,71,29,51,41,22,6,58,5,42,92,85,64,94,12,83,11,17,14,37,36,59,33,0,93,34,70,97,7,76,20,3,88,43,47,8,79,80,63,9,25,56,75,15,4,82,67,39,30,89,86,46,90,48,73,91,55,95,28,49,61,44,84,40,53,13,24}
	//data := []int{7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1}
	seq := make([]int, 0)
	sum, last := 0, 0

out:
	for i := 0; i < len(data); i++ {
		seq = append(seq, data[i])

		for _, board := range boards {
			if len(seq) >= 5 {
				marked := getMarkedIndexes(board, seq)

				if isBingo(marked, "column") {
					sum = utils.IntSum(utils.IntDiff(board, getMVals(marked, board)))
					last = seq[len(seq)-1]
					break out
				}

				if isBingo(marked, "row") {
					sum = utils.IntSum(utils.IntDiff(board, getMVals(marked, board)))
					last = seq[len(seq)-1]
					break out
				}
			}
		}
	}

	return sum * last
}

func calcLastMult(boards [][]int) int {
	data := []int{50,68,2,1,69,32,87,10,31,21,78,23,62,98,16,99,65,35,27,96,66,26,74,72,45,52,81,60,38,57,54,19,18,77,71,29,51,41,22,6,58,5,42,92,85,64,94,12,83,11,17,14,37,36,59,33,0,93,34,70,97,7,76,20,3,88,43,47,8,79,80,63,9,25,56,75,15,4,82,67,39,30,89,86,46,90,48,73,91,55,95,28,49,61,44,84,40,53,13,24}
	//data := []int{7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1}
	seq := make([]int, 0)
	sum, last := 0, 0

	if (len(boards)) == 1 {
		res := calcMult([][]int{0: boards[0]})
		fmt.Println(res)
		return res
	}

	for i := 0; i < len(data); i++ {
		seq = append(seq, data[i])

		for bI, board := range boards {
			if len(seq) >= 5 {
				marked := getMarkedIndexes(board, seq)

				if isBingo(marked, "column") {
					sum = utils.IntSum(utils.IntDiff(board, getMVals(marked, board)))
					last = seq[len(seq)-1]

					remainig := utils.RemoveIndex(boards, bI)

					//fmt.Println(remainig);
					return calcLastMult(remainig)
				}

				if isBingo(marked, "row") {
					sum = utils.IntSum(utils.IntDiff(board, getMVals(marked, board)))
					last = seq[len(seq)-1]

					remainig := utils.RemoveIndex(boards, bI)

					//fmt.Println(remainig)
					return calcLastMult(remainig)
				}
			}
		}
	}

	return sum * last
}

func getMVals(indexes []int, board []int) (vals []int) {
	for _, i := range indexes {
		vals = append(vals, board[i])
	}

	return vals
}