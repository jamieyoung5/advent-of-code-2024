package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadColumnsFromFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Fields(line)

		for i := 0; i < len(words); i++ {
			if len(result) < len(words) {
				result = append(result, make([]string, 0))
			}
			result[i] = append(result[i], words[i])
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func ReadRowsFromFileAsInteger(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Fields(line)

		var intWords []int
		for i := 0; i < len(words); i++ {
			intWord, _ := strconv.Atoi(words[i])
			intWords = append(intWords, intWord)
		}

		result = append(result, intWords)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
