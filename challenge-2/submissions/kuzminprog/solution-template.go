package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input from standard input
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()

		// Call the ReverseString function
		output := ReverseString(input)

		// Print the result
		fmt.Println(output)
	}
}

// ReverseString returns the reversed string of s.
func ReverseString(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteRune(rune(s[i]))
	}
	return builder.String()
}
