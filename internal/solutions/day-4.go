package solutions

import (
	"github.com/jamieyoung5/advent-of-code-2024/internal/util"
	"github.com/jamieyoung5/go-strc-yourself/pkg/algorithms/dfs"
	"github.com/jamieyoung5/go-strc-yourself/pkg/sliceutil"
	"path/filepath"
	"slices"
)

const (
	p1keyword = "XMAS"
	p2keyword = "MAS"
)

var (
	diagonals = []dfs.Direction{dfs.UpLeft, dfs.UpRight, dfs.DownLeft, dfs.DownRight}
)

func Day4Part1(inputPath string) int {
	grid, _ := util.ReadCharsFromFile(filepath.Join(inputPath, "day4.txt"))
	pattern := []rune(p1keyword)

	return len(dfs.SearchGrid(grid, pattern)) // from one of my own go libraries
}

func Day4Part2(inputPath string) int {
	grid, _ := util.ReadCharsFromFile(filepath.Join(inputPath, "day4.txt"))
	pattern := []rune(p2keyword)

	occurrences := dfs.SearchGrid(grid, pattern)
	var middleParts []dfs.Point
	for _, occurrence := range occurrences {
		if slices.Contains(diagonals, occurrence.Direction) {
			middleParts = append(middleParts, occurrence.Points['A'])
		}
	}

	return sliceutil.CountDuplicates(middleParts)
}
