package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func main() {
	// The `flag` package defines flags that can be provided when running a program through the command line
	// All the flags can be displayed running the program with the -h flag
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	// Takes all the provided flags and parses them
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %v\n", *csvFilename)
		os.Exit(1)
	}

	// This reader is able to parse a CSV file contained in a `Reader` instance (`file` is a `Reader`)
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failed to parse the CSV file: %v\n", err)
	}

	problems := parseLines(lines)
	correct := 0

	for idx, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", idx+1, problem.q)

		var answer string

		// Takes input from the keyboard while the program is running and saves it to the memory location of `answer`
		// Scanf is not always ideal as it trims the input and generally can lead to all sorts of problems
		fmt.Scanf("%s\n", &answer)

		if answer == problem.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	// Declares a slice with the same size of the input
	// This could be done through the `append` function but it would be less efficient
	ret := make([]problem, len(lines))

	for idx, line := range lines {
		ret[idx] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}
