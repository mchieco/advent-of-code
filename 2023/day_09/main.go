package day_09

import (
	"fmt"
	"strconv"
	"strings"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

func convertToInts(line string) []int {
	vals := strings.Fields(line)
	ints := make([]int, len(vals))
	for i, val := range vals {
		ints[i], _ = strconv.Atoi(val)
	}
	return ints
}

func calculateRightDifference(ints []int) int {
	if allEqualZero(ints) {
		return 0
	}
	nextRow := make([]int, 0)
	for i := 0; i < len(ints)-1; i++ {
		nextRow = append(nextRow, ints[i+1]-ints[i])
	}
	return ints[(len(ints)-1)] + calculateRightDifference(nextRow)
}

func calculateLeftDifference(ints []int) int {
	if allEqualZero(ints) {
		return 0
	}
	nextRow := make([]int, 0)
	for i := 0; i < len(ints)-1; i++ {
		nextRow = append(nextRow, ints[i+1]-ints[i])
	}
	return ints[0] - calculateLeftDifference(nextRow)
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	sum := 0
	for _, line := range input {
		ints := convertToInts(line)
		sum += calculateRightDifference(ints)
	}
	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	sum := 0
	for _, line := range input {
		ints := convertToInts(line)
		sum += calculateLeftDifference(ints)
	}
	return strconv.Itoa(sum)
}

func allEqualZero(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != 0 {
			return false
		}
	}
	return true
}
