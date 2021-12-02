// Package day2
package day2

import (
	"errors"
	"github.com/soulsteel/aoc-2019-golang/utils"
	"log"
	"strconv"
	"strings"
)

func GetResults() (int, int) {
	fileData := utils.ReadLines("./input/day2.txt")
	return partOne(fileData), partTwo(fileData)
}

type course struct {
	horizontal, depth, aim int
}

func (c *course) calc(direction string) error {
	command, val := parseDirection(direction)
	switch command {
	case "forward":
		c.horizontal += val
		c.depth += val*c.aim
	case "up":
		c.aim -= val
	case "down":
		c.aim += val
	default:
		return errors.New("Unknown command provided: " + command)
	}

	return nil
}

func parseDirection(direction string) (string, int) {
	s := strings.Split(direction, " ")
	val, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}

	return s[0], val
}

func partOne(fileData []string) int {
	c := course{depth: 0, horizontal: 0, aim: 0}
	for _, v := range fileData {
		if err := c.calc(v); err != nil {
			log.Fatal(err)
		}
	}

	return c.horizontal*c.aim
}

func partTwo(fileData []string) int {
	c := course{depth: 0, horizontal: 0, aim: 0}
	for _, v := range fileData {
		if err := c.calc(v); err != nil {
			log.Fatal(err)
		}
	}

	return c.horizontal*c.depth
}
