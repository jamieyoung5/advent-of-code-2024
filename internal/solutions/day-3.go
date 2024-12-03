package solutions

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func Day3Part1(inputPath string) int {
	inputBytes, _ := os.ReadFile(filepath.Join(inputPath, "day3.txt"))
	input := string(inputBytes)

	multiplierRegex := regexp.MustCompile(`mul\((-?\d+),(-?\d+)\)`)
	stmts := multiplierRegex.FindAllStringSubmatch(input, -1)

	var result int
	for _, stmt := range stmts {
		num1, _ := strconv.Atoi(stmt[1])
		num2, _ := strconv.Atoi(stmt[2])

		result += num1 * num2
	}

	return result
}

func Day3Part2(inputPath string) int {
	inputBytes, _ := os.ReadFile(filepath.Join(inputPath, "day3.txt"))
	input := string(inputBytes)

	multiplierRegex := regexp.MustCompile(`mul\((-?\d+),(-?\d+)\)|do\(\)|don't\(\)`)
	stmts := multiplierRegex.FindAllStringSubmatch(input, -1)

	multiplyEnabled := true

	var result int
	for _, stmt := range stmts {
		switch stmt[0] {
		case "do()":
			multiplyEnabled = true
			continue
		case "don't()":
			multiplyEnabled = false
			continue
		default:
			if multiplyEnabled {
				num1, _ := strconv.Atoi(stmt[1])
				num2, _ := strconv.Atoi(stmt[2])

				result += num1 * num2
			}
		}
	}

	return result
}
