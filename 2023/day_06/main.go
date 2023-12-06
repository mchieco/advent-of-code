package day_06

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

type game struct {
	Duration int
	Record   int
}

func parseInputPartOne(input []string) []game {
	games := make([]game, 0)
	for i, line := range input {
		fields := strings.Fields(strings.Split(line, ":")[1])
		for j, field := range fields {
			field, _ := strconv.Atoi(field)
			if i == 0 {
				games = append(games, game{Duration: field})
			}
			if i == 1 {
				games[j].Record = field
			}
		}
	}
	return games
}

func calculateNumWon(game game) int {
	numWon := 0
	for held := 1; held < game.Duration; held++ {
		distance := (game.Duration - held) * held
		if distance > game.Record {
			numWon++
		}
	}
	return numWon
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	games := parseInputPartOne(input)
	score := 1
	for _, game := range games {
		numWon := calculateNumWon(game)
		score *= numWon
	}
	return strconv.Itoa(score)
}

func parseInputPartTwo(input []string) game {
	game := game{}
	for i, line := range input {
		numStr := strings.Join(strings.Fields(strings.Split(line, ":")[1]), "")
		num := convertToInt(numStr)
		if i == 0 {
			game.Duration = num
		}
		if i == 1 {
			game.Record = num
		}
	}
	return game
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	game := parseInputPartTwo(input)
	score := calculateNumWon(game)
	return strconv.Itoa(score)
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
