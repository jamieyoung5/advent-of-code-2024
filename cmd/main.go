package main

import (
	"advent-of-code-2024/internal/solutions"
	"fmt"
)

const inputsPath = "C:\\Users\\jamie\\git\\advent-of-code-2024\\inputs"

func main() {

	day1Part1, err := solutions.Day1Part1(inputsPath)
	if err != nil {
		panic(err)
	}

	day1Part2, err := solutions.Day1Part2(inputsPath)
	fmt.Printf("Solutions:\n Day1: (P1) %d, (P2) %d\n", day1Part1, day1Part2)

	day2Part1, err := solutions.Day2Part1(inputsPath)
	if err != nil {
		panic(err)
	}

	day2Part2, err := solutions.Day2Part2(inputsPath)
	fmt.Printf("Solutions:\n Day2: (P1) %d, (P2) %d\n", day2Part1, day2Part2)
}
