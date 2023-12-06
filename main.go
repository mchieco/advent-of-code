package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/mchieco/advent-of-code/internal"
)

// main entry point
// The --day parameter is required to choose which daily challenge should be executed.
// The --input parameter is also required, and it points to the input file that should be used for the challenge.
// The --mode parameter specifies which part of the challenge should be executed:
// - 1: only the first part
// - 2: only the second part
// - 3 or empty: both parts
func main() {
	y := flag.String("year", "", "year ID to execute")
	d := flag.String("day", "", "day ID to execute")
	i := flag.String("input", "", "input file path")
	m := flag.String("mode", "3", "running mode")
	flag.Parse()

	if y == nil || d == nil || i == nil {
		fmt.Println("missing required input params")
		os.Exit(1)
	}

	day, err := strconv.Atoi(*d)
	if err != nil {
		fmt.Println("couldn't parse day")
		os.Exit(1)
	}

	mode, err := strconv.Atoi(*m)
	if err != nil {
		fmt.Println("incorrect mode")
		os.Exit(1)
	}

	inputPath := *i
	internal.RunChallenge(day, inputPath, mode)
}
