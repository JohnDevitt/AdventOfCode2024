package utils

import (
	"bufio"
	"os"
)

func ReadLines(filePath string) []string {
	file, _ := os.Open(filePath)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
