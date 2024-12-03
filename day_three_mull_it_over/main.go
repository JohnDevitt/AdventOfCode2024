package main

import (
	"adventOfCode/utils"
	"regexp"
	"strings"
)

func readData() string {
	data := utils.ReadLines("day_three_mull_it_over/input.txt")
	return strings.Join(data, "")
}

func parseMulOperation(mulOperation string) [2]int {
	trimmed := strings.TrimPrefix(mulOperation, "mul(")
	trimmed = strings.TrimSuffix(trimmed, ")")
	parts := strings.Split(trimmed, ",")
	partNumbers := utils.AListToIList(parts)
	return [2]int{partNumbers[0], partNumbers[1]}
}

func sumProductOfPairs(pairs [][2]int) int {
	sum := 0
	for _, pair := range pairs {
		sum += pair[0] * pair[1]
	}
	return sum
}

func isThreeDigitPair(pair [2]int) bool {
	return pair[0]/1000 < 1 && pair[1]/1000 < 1
}

func partOne() int {
	data := readData()
	re, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(data, -1)

	var pairs [][2]int
	for _, match := range matches {
		pair := parseMulOperation(match)
		if isThreeDigitPair(pair) {
			pairs = append(pairs, pair)
		}
	}

	return sumProductOfPairs(pairs)
}

func partTwo() int {
	data := readData()
	re, _ := regexp.Compile(`(do\(\))|(don't\(\))|mul\(\d+,\d+\)`)
	matches := re.FindAllString(data, -1)

	mulEnabled := true

	var pairs [][2]int
	for _, match := range matches {
		switch match {
		case "do()":
			mulEnabled = true
		case "don't()":
			mulEnabled = false
		default:
			pair := parseMulOperation(match)
			if mulEnabled && isThreeDigitPair(pair) {
				pairs = append(pairs, pair)
			}
		}
	}

	return sumProductOfPairs(pairs)
}

func main() {
	utils.PrintTitle("Day 3: Mull It Over")
	utils.TimeAndFormat(partOne, "part one")
	utils.TimeAndFormat(partTwo, "part two")
}
