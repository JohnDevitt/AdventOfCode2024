package main

import (
	"adventOfCode/utils"
	"slices"
	"strings"
)

func readReports() [][]int {
	lines := utils.ReadLines("day_two_red_nosed_reports/input.txt")
	var reports [][]int

	for _, line := range lines {
		reports = append(reports, utils.AListToIList(strings.Fields(line)))
	}

	return reports
}

func isLevelSafe(first, second int, isDescending bool) bool {
	if utils.AbsoluteDifference(first, second) > 3 || utils.AbsoluteDifference(first, second) == 0 {
		return false
	}
	if isDescending && first < second {
		return false
	}
	if !isDescending && second < first {
		return false
	}
	return true
}

func isReportSafe(report []int) bool {
	isDescending := report[0] > report[len(report)-1]
	for i := range report[:len(report)-1] {
		if !isLevelSafe(report[i], report[i+1], isDescending) {
			return false
		}
	}
	return true
}

func partOne() int {
	reports := readReports()

	count := 0
	for _, report := range reports {
		if isReportSafe(report) {
			count++
		}
	}

	return count
}

func isProblemDampenedReportSafe(report []int) bool {
	for index := range report {
		var reportClone = slices.Clone(report)
		var subReport = slices.Delete(reportClone, index, index+1)
		if isReportSafe(subReport) {
			return true
		}
	}
	return false
}

func partTwo() int {
	reports := readReports()

	count := 0
	for _, report := range reports {
		if isProblemDampenedReportSafe(report) {
			count++
		}
	}

	return count
}

func main() {
	utils.PrintTitle("Day 2: Red-Nosed Reports")
	utils.TimeAndFormat(partOne, "part one")
	utils.TimeAndFormat(partTwo, "part two")
}
