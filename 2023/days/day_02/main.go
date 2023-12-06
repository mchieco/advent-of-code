package day_02

import (
	"fmt"
	"log"
	"slices"
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

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	sumOfGames := 0
	for _, game := range input {
		gameSplit := strings.Split(game, ":")
		gameIDParsed := strings.Split(gameSplit[0], " ")[1]
		gameID, err := strconv.Atoi(gameIDParsed)
		if err != nil {
			log.Fatal(err)
		}
		games := strings.Split(gameSplit[1], ";")
		for i, game := range games {
			colorPulls := strings.Split(game, ",")
			if !IsAllowedTurn(colorPulls) {
				break
			}
			if i == len(games)-1 {
				sumOfGames += gameID
			}
		}
	}
	return strconv.Itoa(sumOfGames)
}

func IsAllowedTurn(colorPulls []string) bool {
	allowedColorSizes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	for _, colorPull := range colorPulls {
		colorPull = strings.TrimSpace(colorPull)
		identifiers := strings.Split(colorPull, " ")
		numberParsed := identifiers[0]
		color := identifiers[1]
		number, err := strconv.Atoi(numberParsed)
		if err != nil {
			log.Fatal(err)
		}
		if number > allowedColorSizes[color] {
			return false
		}
	}
	return true
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	sumOfGames := 0
	for _, game := range input {
		gameSplit := strings.Split(game, ":")
		red := make([]int, 0)
		green := make([]int, 0)
		blue := make([]int, 0)
		games := strings.Split(gameSplit[1], ";")
		for _, game := range games {
			colorPulls := strings.Split(game, ",")
			for _, colorPull := range colorPulls {
				info := strings.TrimSpace(colorPull)
				identifiers := strings.Split(info, " ")
				numberParsed := identifiers[0]
				color := identifiers[1]
				number, err := strconv.Atoi(numberParsed)
				if err != nil {
					log.Fatal(err)
				}
				if color == "red" {
					red = append(red, number)
				} else if color == "green" {
					green = append(green, number)
				} else if color == "blue" {
					blue = append(blue, number)
				}
			}
		}
		redMax := slices.Max(red)
		greenMax := slices.Max(green)
		blueMax := slices.Max(blue)
		sumOfGames += (redMax * greenMax * blueMax)
	}
	return strconv.Itoa(sumOfGames)
}
