package day_08

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	Value string
	Left  string
	Right string
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

func parseInput(input []string) (string, []string, map[string]node) {
	nodes := make(map[string]node, 0)
	startNodes := make([]string, 0)
	var directions string
	for i, line := range input {
		if i == 0 {
			directions = line
			continue
		}
		if i == 1 {
			continue
		}
		node := node{}
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		split := strings.Split(line, " = ")
		split2 := strings.Split(split[1], ", ")
		node.Value = split[0]
		node.Left = split2[0]
		node.Right = split2[1]
		nodes[split[0]] = node
		if node.Value[2] == 'A' {
			startNodes = append(startNodes, node.Value)
		}
	}
	return directions, startNodes, nodes
}

func calculateStepsPart1(directions string, nodes map[string]node) int {
	requiredSteps := 0
	currentNode := nodes["AAA"]
	for i := 0; i < len(directions); i++ {
		if directions[i] == 'L' {
			currentNode = nodes[currentNode.Left]
		} else {
			currentNode = nodes[currentNode.Right]
		}
		requiredSteps++
		if currentNode.Value == "ZZZ" {
			return requiredSteps
		}
		if i == len(directions)-1 {
			i = -1
		}
	}
	return requiredSteps
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	directions, _, nodes := parseInput(input)
	requiredSteps := calculateStepsPart1(directions, nodes)
	return strconv.Itoa(requiredSteps)
}

func calculateStepsPart2(directions string, startNodes []string, nodes map[string]node) []int {
	results := []int{}

	for _, startNode := range startNodes {
		currentNode := nodes[startNode]
		requiredSteps := 0
		for i := 0; i < len(directions); i++ {
			if directions[i] == 'L' {
				currentNode = nodes[currentNode.Left]
			} else {
				currentNode = nodes[currentNode.Right]
			}
			requiredSteps++
			if currentNode.Value[len(currentNode.Value)-1:] == "Z" {
				results = append(results, requiredSteps)
				break
			}
			if i == len(directions)-1 {
				i = -1
			}
		}
	}
	return results
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	directions, startingNodes, nodes := parseInput(input)
	results := calculateStepsPart2(directions, startingNodes, nodes)
	lcm := LCM(results[0], results[1], results[2:]...)
	return strconv.Itoa(lcm)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
