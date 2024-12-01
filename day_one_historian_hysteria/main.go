package main

import (
	"adventOfCode/utils"
	"sort"
	"strconv"
	"strings"
)

func absoluteDifference(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func buildLists() ([]int, []int) {
	lines := utils.ReadLines("day_one_historian_hysteria/input.txt")

	var left, right []int
	for _, line := range lines {
		numbers := strings.Fields(line)
		start, _ := strconv.Atoi(numbers[0])
		end, _ := strconv.Atoi(numbers[1])
		left = append(left, start)
		right = append(right, end)
	}

	return left, right
}

func partOne() int {
	left, right := buildLists()

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for index, leftNum := range left {
		sum += absoluteDifference(right[index], leftNum)
	}

	return sum
}

func partTwo() int {
	left, right := buildLists()

	sum := 0
	for _, leftNum := range left {
		count := 0
		for _, rightNum := range right {
			if leftNum == rightNum {
				count++
			}
		}
		sum += count * leftNum
	}

	return sum
}

func main() {
	utils.PrintTitle("Day 1: Historian Hysteria")
	utils.TimeAndFormat(partOne, "part one")
	utils.TimeAndFormat(partTwo, "part two")
}
