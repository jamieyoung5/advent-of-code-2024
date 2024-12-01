package main

import (
	"advent-of-code-2024/internal/solutions"
	"fmt"
)

func main() {

	day1Part1, err := solutions.Day1Part1("C:\\Users\\jamie\\git\\advent-of-code-2024\\inputs")
	if err != nil {
		panic(err)
	}

	day1Part2, err := solutions.Day1Part2("C:\\Users\\jamie\\git\\advent-of-code-2024\\inputs")
	fmt.Printf("Solutions:\n Day1: (P1) %d, (P2) %d", day1Part1, day1Part2)
}
