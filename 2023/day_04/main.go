package day_04

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// card struct representing a scratchcard
type card struct {
	index   int
	winners []int
	all     []int
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
	ans := 0
	for _, c := range input {
		cardScore := 0
		splitCard := strings.Split(c, "|")
		ourCard := strings.Split(strings.TrimSpace(splitCard[1]), " ")
		ourCard = removeWhitespace(ourCard)
		winningCard := strings.Split(strings.TrimSpace(strings.Split(splitCard[0], ":")[1]), " ")
		winningCard = removeWhitespace(winningCard)
		for _, number := range ourCard {
			if slices.Contains(winningCard, number) {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore *= 2
				}
			}
		}
		ans += cardScore
	}
	return strconv.Itoa(ans)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	answerMap := make(map[int]int)
	for _, c := range input {
		numWon := 0
		splitCard := strings.Split(c, "|")
		ourCard := strings.Split(strings.TrimSpace(splitCard[1]), " ")
		ourCard = removeWhitespace(ourCard)
		cardNumberSlice := removeWhitespace(strings.Split(strings.Split(splitCard[0], ":")[0], " "))
		cardNumber, _ := strconv.Atoi(cardNumberSlice[1])
		winningCard := strings.Split(strings.TrimSpace(strings.Split(splitCard[0], ":")[1]), " ")
		winningCard = removeWhitespace(winningCard)
		for _, number := range ourCard {
			if slices.Contains(winningCard, number) {
				numWon++
			}
		}
		answerMap[cardNumber] += 1
		for i := 1; i <= numWon; i++ {
			answerMap[cardNumber+i] += answerMap[cardNumber]
		}
	}
	ans := 0
	for _, v := range answerMap {
		ans += v
	}
	return strconv.Itoa(ans)
}

func removeWhitespace(ss []string) (ret []string) {
	for _, s := range ss {
		if s == "" {
			continue
		}
		ret = append(ret, s)
	}
	return
}
