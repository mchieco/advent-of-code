package day_05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type row struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

type seed struct {
	Value int
	Range int
	Seen  bool
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	seeds := []seed{}
	almanac := make([][]row, 0)
	group := make([]row, 0)
	firstEmpty := false
	for i, line := range input {
		if strings.Split(line, ":")[0] == "seeds" {
			seedsStrings := strings.Fields(strings.Split(line, ":")[1])
			for _, nseed := range seedsStrings {
				seeds = append(seeds, seed{
					Value: convertToInt(nseed),
				})
			}
		}
		if len(strings.Split(line, ":")) > 1 {
			continue
		}
		if line == "" || i == len(input)-1 {
			if !firstEmpty {
				firstEmpty = true
				continue
			}
			almanac = append(almanac, group)
			group = make([]row, 0)
			continue
		}
		rowVals := strings.Fields(line)
		group = append(group, row{DestinationRangeStart: convertToInt(rowVals[0]), SourceRangeStart: convertToInt(rowVals[1]), RangeLength: convertToInt(rowVals[2])})
	}
	for _, convMap := range almanac {
		for i := range seeds {
			seeds[i].Seen = false
		}
		for _, row := range convMap {
			for i, seed := range seeds {
				if seed.Value >= row.SourceRangeStart && seed.Value < row.SourceRangeStart+row.RangeLength && !seed.Seen {
					seeds[i].Seen = true
					seeds[i].Value = row.DestinationRangeStart + (seed.Value - row.SourceRangeStart)
				}
			}
		}
	}
	ans := 0
	for _, seed := range seeds {
		if seed.Value < ans || ans == 0 {
			ans = seed.Value
		}
	}
	return strconv.Itoa(ans)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	seeds := []seed{}
	almanac := make([][]row, 0)
	group := make([]row, 0)
	firstEmpty := false
	for i, line := range input {
		if strings.Split(line, ":")[0] == "seeds" {
			seedsStrings := strings.Fields(strings.Split(line, ":")[1])
			for index, nseed := range seedsStrings {
				if (index+1)%2 == 0 {
					seeds = append(seeds, seed{
						Range: convertToInt(nseed),
						Value: convertToInt(seedsStrings[index-1]),
					})
				}

			}
		}
		if len(strings.Split(line, ":")) > 1 {
			continue
		}
		if line == "" || i == len(input)-1 {
			if !firstEmpty {
				firstEmpty = true
				continue
			}
			almanac = append(almanac, group)
			group = make([]row, 0)
			continue
		}
		rowVals := strings.Fields(line)
		group = append(group, row{DestinationRangeStart: convertToInt(rowVals[0]), SourceRangeStart: convertToInt(rowVals[1]), RangeLength: convertToInt(rowVals[2])})
	}
	slices.Reverse(almanac)
	ans := 0
	for {
		value := ans
		for _, convMap := range almanac {
			for _, row := range convMap {
				if value >= row.DestinationRangeStart && value < row.DestinationRangeStart+row.RangeLength {
					value = row.SourceRangeStart + (value - row.DestinationRangeStart)
					break
				}
			}
		}
		for _, seed := range seeds {
			if value >= seed.Value && value < seed.Range+seed.Value {
				return strconv.Itoa(ans)
			}
		}
		ans += 1
	}
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
