package solutions

import (
	"advent-of-code-2024/internal/util"
	"math"
	"path/filepath"
)

func Day2Part1(inputsPath string) (int, error) {
	input, err := util.ReadRowsFromFileAsInteger(filepath.Join(inputsPath, "day2.txt"))
	if err != nil {
		return -1, err
	}

	var safeReports int
	for _, row := range input {
		if isReportSafe(row) {
			safeReports++
		}
	}

	return safeReports, nil
}

func Day2Part2(inputsPath string) (int, error) {
	input, err := util.ReadRowsFromFileAsInteger(filepath.Join(inputsPath, "day2.txt"))
	if err != nil {
		return -1, err
	}

	var safeReports int
	for _, row := range input {
		if isReportSafe(row) {
			safeReports++
			continue
		}

		if isSafeWithRemoval(row) {
			safeReports++
		}
	}

	return safeReports, nil
}

func isSafeWithRemoval(report []int) bool {
	for i := range report {
		modified := make([]int, 0, len(report)-1)
		modified = append(modified, report[:i]...)
		modified = append(modified, report[i+1:]...)

		if isReportSafe(modified) {
			return true
		}
	}

	return false
}

func isReportSafe(report []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(report)-1; i++ {
		diff := math.Abs(float64(report[i] - report[i+1]))

		if diff < 1 || diff > 3 {
			return false
		}

		if report[i] >= report[i+1] {
			increasing = false
			continue
		}
		decreasing = false
	}

	return increasing || decreasing
}
