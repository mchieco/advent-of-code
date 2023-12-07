package day_07

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// hand strength:
// 0: high card
// 1: pair
// 2: two pair
// 3: three of a kind
// 4: full house
// 5: four of a kind
// 6: five of a kind
type hand struct {
	cards    []int
	bid      int
	strength int
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

func calculateThreeCardOptions(mapOfCards map[int]int, cardNum int) int {
	for cardNum2, amount2 := range mapOfCards {
		if cardNum2 == cardNum {
			continue
		}
		if amount2 == 2 {
			return 4
		}
	}
	return 3
}

func calculateTwoCardOptions(mapOfCards map[int]int, cardNum int) int {
	for cardNum2, amount2 := range mapOfCards {
		if cardNum2 == cardNum {
			continue
		}
		if amount2 == 3 {
			return 4

		}
		if amount2 == 2 {
			return 2
		}
	}
	return 1
}

func orderHands(hands []hand) []hand {
	for i := 0; i < len(hands); i++ {
		for j := 0; j < len(hands)-i-1; j++ {
			if hands[j].strength > hands[j+1].strength {
				hands[j], hands[j+1] = hands[j+1], hands[j]
			}
			if hands[j].strength == hands[j+1].strength {
				for k := 0; k < len(hands[j].cards); k++ {
					if hands[j].cards[k] < hands[j+1].cards[k] {
						break
					}
					if hands[j].cards[k] > hands[j+1].cards[k] {
						hands[j], hands[j+1] = hands[j+1], hands[j]
						break
					}
				}
			}
		}
	}
	return hands
}

func calculateTotalWinnings(hands []hand) int {
	total := 0
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return total
}

var faceCards = map[rune]int{
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func getHandsWithStrengthPart1(input []string) []hand {
	hands := make([]hand, 0)
	for _, line := range input {
		split := strings.Split(line, " ")
		hand := hand{
			bid:   convertToInt(split[1]),
			cards: make([]int, 0),
		}
		for _, card := range split[0][0:5] {
			if value, ok := faceCards[card]; ok {
				hand.cards = append(hand.cards, value)
			} else {
				hand.cards = append(hand.cards, convertToInt(string(card)))
			}
		}
		hand.strength = calculateHandStrengthPart1(hand)
		hands = append(hands, hand)
	}
	return hands
}

func calculateHandStrengthPart1(hand hand) int {
	mapOfCards := make(map[int]int)
	for _, card := range hand.cards {
		mapOfCards[card] += 1
	}
	if len(mapOfCards) == 5 {
		return 0
	}
	for cardNum, amount := range mapOfCards {
		if amount >= 4 {
			return amount + 1
		}
		if amount == 3 {
			return calculateThreeCardOptions(mapOfCards, cardNum)
		}
		if amount == 2 {
			return calculateTwoCardOptions(mapOfCards, cardNum)
		}
	}
	return 0
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	hands := getHandsWithStrengthPart1(input)
	hands = orderHands(hands)
	return strconv.Itoa(calculateTotalWinnings(hands))
}

var faceCards2 = map[rune]int{
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func getHandsWithStrengthPart2(input []string) []hand {
	hands := make([]hand, 0)
	for _, line := range input {
		split := strings.Split(line, " ")
		hand := hand{
			bid:   convertToInt(split[1]),
			cards: make([]int, 0),
		}
		for _, card := range split[0][0:5] {
			if value, ok := faceCards2[card]; ok {
				hand.cards = append(hand.cards, value)
			} else {
				hand.cards = append(hand.cards, convertToInt(string(card)))
			}
		}
		hand.strength = calculateHandStrengthPart2(hand)
		hands = append(hands, hand)
	}
	return hands
}

func calculateHandStrengthPart2(hand hand) int {
	mapOfCards := make(map[int]int)
	numberOfJacks := 0
	for _, card := range hand.cards {
		if card == 1 {
			numberOfJacks++
		} else {
			mapOfCards[card] += 1
		}
	}
	if numberOfJacks == 5 {
		return 6
	}
	if len(mapOfCards) == 5 {
		return 0
	}
	scores := make([]int, 0)
	for cardNum, amount := range mapOfCards {
		amount += numberOfJacks
		if amount >= 4 {
			scores = append(scores, amount+1)
		}
		if amount == 3 {
			scores = append(scores, calculateThreeCardOptions(mapOfCards, cardNum))
		}
		if amount == 2 {
			scores = append(scores, calculateTwoCardOptions(mapOfCards, cardNum))
		}
	}
	if len(scores) == 0 {
		return 0
	}
	return slices.Max(scores)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	hands := getHandsWithStrengthPart2(input)
	hands = orderHands(hands)
	return strconv.Itoa(calculateTotalWinnings(hands))
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
