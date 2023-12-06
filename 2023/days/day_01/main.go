package day_01

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
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
	combinedValues := 0
	for _, line := range input {
		combinedValue := ""
		for _, c := range line {
			v := string(c)
			if unicode.IsDigit(c) {
				combinedValue += v
				break
			}
		}
		reverse := Reverse(line)
		for _, c := range reverse {
			v := string(c)
			if unicode.IsDigit(c) {
				combinedValue += v
				break
			}
		}
		combinedValueInt, err := strconv.Atoi(combinedValue)
		if err != nil {
			log.Fatal(err)
		}
		combinedValues += combinedValueInt
	}
	return strconv.Itoa(combinedValues)
}

var numberMap = map[string]string{
	"one":   "o1ne",
	"two":   "t2wo",
	"three": "th3ree",
	"four":  "f4our",
	"five":  "f5ive",
	"six":   "s6ix",
	"seven": "s7even",
	"eight": "e8ight",
	"nine":  "n9ine",
}

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	combinedValues := 0
	for _, line := range input {
		combinedValue := ""
		for k, v := range numberMap {
			line = strings.ReplaceAll(line, k, v)
		}
		for _, c := range line {
			v := string(c)
			if unicode.IsDigit(c) {
				combinedValue += v
				break
			}
		}
		reverse := Reverse(line)
		for _, c := range reverse {
			v := string(c)
			if unicode.IsDigit(c) {
				combinedValue += v
				break
			}
		}
		combinedValueInt, err := strconv.Atoi(combinedValue)
		if err != nil {
			log.Fatal(err)
		}
		combinedValues += combinedValueInt
	}
	return strconv.Itoa(combinedValues)
}
