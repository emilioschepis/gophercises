package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Scanf("%s\n", &input)
	fmt.Println("Input is:", input)

	min, max := 'A', 'Z'

	// The first word is not capitalized, but it counts.
	answer := 1
	for _, ch := range input {
		// Single quotes to represent a single rune.
		if ch >= min && ch <= max {
			answer++
		}

		// Other approach:
		// str := string(ch)
		// if strings.ToUpper(str) == str {
		// 	answer++
		// }
	}

	fmt.Println(answer)
}
