package day_03

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

type number struct {
	Value int
	Start int
	Stop  int
	Row   int
}

type element struct {
	Number   number
	IsDigit  bool
	IsSymbol bool
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
	board := make([][]element, 0)
	for y, line := range input {
		row := make([]element, 0)
		fullNumber := ""
		for i, c := range line {
			element := element{}
			v := string(c)
			if v == "." {
				if fullNumber != "" {
					if fullNumberInt, err := strconv.Atoi(fullNumber); err == nil {
						for j := 0; j < len(fullNumber); j++ {
							row[i-1-j].Number = number{Value: fullNumberInt, Start: i - len(fullNumber), Stop: i, Row: y}
						}
						fullNumber = ""
					} else {
						log.Fatal(err)
					}
				}
				row = append(row, element)
				continue
			} else if unicode.IsDigit(c) {
				fullNumber += string(c)
				if fullNumber != "" && i == len(line)-1 {
					if fullNumberInt, err := strconv.Atoi(fullNumber); err == nil {
						for j := 0; j < len(fullNumber); j++ {
							row[i-1-j].Number = number{Value: fullNumberInt, Start: i - len(fullNumber), Stop: i, Row: y}
						}
						fullNumber = ""
					} else {
						log.Fatal(err)
					}
				}
				element.IsDigit = true
				row = append(row, element)
				continue
			} else {
				if fullNumber != "" {
					if fullNumberInt, err := strconv.Atoi(fullNumber); err == nil {
						for j := 0; j < len(fullNumber); j++ {
							row[i-1-j].Number = number{Value: fullNumberInt, Start: i - len(fullNumber), Stop: i, Row: y}
						}
						fullNumber = ""
					} else {
						log.Fatal(err)
					}
				}
				element.IsSymbol = true
				row = append(row, element)
				continue
			}
		}
		board = append(board, row)
	}
	seen := make([]number, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j].IsDigit {
				isAdjacentSymbol := isAdjacentSymbol(board, i, j)
				if isAdjacentSymbol {
					wasAdded := false
					for _, n := range seen {
						if n.Value == board[i][j].Number.Value && n.Row == board[i][j].Number.Row && n.Start == board[i][j].Number.Start && n.Stop == board[i][j].Number.Stop {
							wasAdded = true
						}
					}
					if !wasAdded {
						ans += board[i][j].Number.Value
						seen = append(seen, board[i][j].Number)
					}
				}
			}
		}
	}
	return strconv.Itoa(ans)
}

func unique(numbers []number) []number {
	var unique []number
numberLoop:
	for _, v := range numbers {
		for i, u := range unique {
			if v.Value == u.Value && v.Start == u.Start && v.Stop == u.Stop && v.Row == u.Row {
				unique[i] = v
				continue numberLoop
			}
		}
		unique = append(unique, v)
	}
	return unique
}

func isAdjacentSymbol(board [][]element, i int, j int) bool {
	// up
	if i > 0 && board[i-1][j].IsSymbol {
		return true
	}
	// up left
	if i > 0 && j > 0 && board[i-1][j-1].IsSymbol {
		return true
	}
	// up right
	if i > 0 && j < len(board[i])-1 && board[i-1][j+1].IsSymbol {
		return true
	}
	// down
	if i < len(board)-1 && board[i+1][j].IsSymbol {
		return true
	}
	// down left
	if i < len(board)-1 && j > 0 && board[i+1][j-1].IsSymbol {
		return true
	}
	// down right
	if i < len(board)-1 && j < len(board[i])-1 && board[i+1][j+1].IsSymbol {
		return true
	}
	// left
	if j > 0 && board[i][j-1].IsSymbol {
		return true
	}
	// right
	if j < len(board[i])-1 && board[i][j+1].IsSymbol {
		return true
	}
	return false
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	ans := 0
	board := make([][]element, 0)
	for y, line := range input {
		row := make([]element, 0)
		fullNumber := ""
		for i, c := range line {
			element := element{}
			v := string(c)
			if v == "." {
				if fullNumber != "" {
					if fullNumberInt, err := strconv.Atoi(fullNumber); err == nil {
						for j := 0; j < len(fullNumber); j++ {
							row[i-1-j].Number = number{Value: fullNumberInt, Start: i - len(fullNumber), Stop: i, Row: y}
						}
						fullNumber = ""
					} else {
						log.Fatal(err)
					}
				}
				row = append(row, element)
				continue
			} else if unicode.IsDigit(c) {
				fullNumber += string(c)
				if fullNumber != "" && i == len(line)-1 {
					if fullNumberInt, err := strconv.Atoi(fullNumber); err == nil {
						for j := 0; j < len(fullNumber); j++ {
							row[i-1-j].Number = number{Value: fullNumberInt, Start: i - len(fullNumber), Stop: i, Row: y}
						}
						fullNumber = ""
					} else {
						log.Fatal(err)
					}
				}
				element.IsDigit = true
				row = append(row, element)
				continue
			} else {
				if fullNumber != "" {
					if fullNumberInt, err := strconv.Atoi(fullNumber); err == nil {
						for j := 0; j < len(fullNumber); j++ {
							row[i-1-j].Number = number{Value: fullNumberInt, Start: i - len(fullNumber), Stop: i, Row: y}
						}
						fullNumber = ""
					} else {
						log.Fatal(err)
					}
				}
				element.IsSymbol = true
				row = append(row, element)
				continue
			}
		}
		board = append(board, row)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j].IsSymbol {
				adjacentNumbers := getAdjacentNumbers(board, i, j)
				if len(adjacentNumbers) == 2 {
					ans += adjacentNumbers[0].Value * adjacentNumbers[1].Value
				}
			}
		}
	}
	return strconv.Itoa(ans)
}

func getAdjacentNumbers(board [][]element, i int, j int) []number {
	numbers := make([]number, 0)
	// up
	if i > 0 && board[i-1][j].IsDigit {
		numbers = append(numbers, board[i-1][j].Number)
	}
	// up left
	if i > 0 && j > 0 && board[i-1][j-1].IsDigit {
		numbers = append(numbers, board[i-1][j-1].Number)
	}
	// up right
	if i > 0 && j < len(board[i])-1 && board[i-1][j+1].IsDigit {
		numbers = append(numbers, board[i-1][j+1].Number)
	}
	// down
	if i < len(board)-1 && board[i+1][j].IsDigit {
		numbers = append(numbers, board[i+1][j].Number)
	}
	// down left
	if i < len(board)-1 && j > 0 && board[i+1][j-1].IsDigit {
		numbers = append(numbers, board[i+1][j-1].Number)
	}
	// down right
	if i < len(board)-1 && j < len(board[i])-1 && board[i+1][j+1].IsDigit {
		numbers = append(numbers, board[i+1][j+1].Number)
	}
	// left
	if j > 0 && board[i][j-1].IsDigit {
		numbers = append(numbers, board[i][j-1].Number)
	}
	// right
	if j < len(board[i])-1 && board[i][j+1].IsDigit {
		numbers = append(numbers, board[i][j+1].Number)
	}

	return unique(numbers)
}
