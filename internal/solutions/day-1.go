package solutions

import (
	"github.com/jamieyoung5/advent-of-code-2024/internal/util"
	"github.com/jamieyoung5/go-strc-yourself/pkg/sliceutil"
	"math"
	"path/filepath"
	"slices"
	"strconv"
)

func Day1Part2(inputsPath string) (int, error) {
	input, err := util.ReadColumnsFromFile(filepath.Join(inputsPath, "day1.txt"))
	if err != nil {
		return -1, err
	}

	occurences := sliceutil.Counts(input[1])

	var similarityScore int
	for _, val := range input[0] {
		valNum, _ := strconv.Atoi(val)
		similarityScore += occurences[val] * valNum
	}

	return similarityScore, nil
}

func Day1Part1(inputsPath string) (int, error) {
	input, err := util.ReadColumnsFromFile(filepath.Join(inputsPath, "day1.txt"))
	if err != nil {
		return -1, err
	}

	column1 := sortAscending(input[0])
	column2 := sortAscending(input[1])

	var total float64
	for i := 0; i < len(input[0]); i++ {
		num1, _ := strconv.ParseFloat(column1[i], 64)
		num2, _ := strconv.ParseFloat(column2[i], 64)
		total += math.Abs(num1 - num2)
	}

	return int(total), nil
}

func sortAscending(column []string) []string {
	slices.SortFunc(column, func(a, b string) int {
		numA, _ := strconv.Atoi(a)
		numB, _ := strconv.Atoi(b)

		if numA < numB {
			return -1
		}
		return 1
	})

	return column
}
