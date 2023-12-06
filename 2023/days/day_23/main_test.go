package day_23_test

import (
	"testing"

	"github.com/mchieco/advent-of-code/2023/days/day_23"
	"github.com/mchieco/advent-of-code/internal"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	input := internal.LoadInputLines("input_1_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_1.txt")
	result := day_23.Part1(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	input := internal.LoadInputLines("input_2_test.txt")
	expectedResult := internal.LoadFirstInputLine("solution_2.txt")
	result := day_23.Part2(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}