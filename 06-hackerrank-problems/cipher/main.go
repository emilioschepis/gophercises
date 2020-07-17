package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	ret := ""
	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetLower, ch) != -1:
			ret = ret + string(rotate(ch, delta, alphabetLower))
		case strings.IndexRune(alphabetUpper, ch) != -1:
			ret = ret + string(rotate(ch, delta, alphabetUpper))
		default:
			ret = ret + string(ch)
		}
	}

	fmt.Println(ret)
}

func rotate(s rune, delta int, key string) rune {
	idx := strings.IndexRune(key, s)
	if idx == -1 {
		panic("rune not found")
	}
	idx = (idx + delta) % len(key)
	return []rune(key)[idx]
}
