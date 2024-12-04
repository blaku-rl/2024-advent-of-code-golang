package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func parseInput() [][]int {
	reports := make([][]int, 0, 1000)
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}
		report := make([]int, 0, 10)
		for _, num := range strings.Fields(row) {
			level, _ := strconv.Atoi(num)
			report = append(report, level)
		}
		reports = append(reports, report)
	}

	return reports
}

func isReportSafe(report []int) bool {
	isSafe := true
	isPositive := report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		if !isSafeLevelChange(report[i], report[i+1], isPositive) {
			isSafe = false
			break
		}
	}
	return isSafe
}

func isSafeLevelChange(prevLevel int, nextLevel int, isPositive bool) bool {
	if isPositive && (prevLevel >= nextLevel || prevLevel < nextLevel-3) {
		return false
	}
	if !isPositive && (prevLevel <= nextLevel || prevLevel > nextLevel+3) {
		return false
	}
	return true
}

func partone() {
	reports := parseInput()
	safeReports := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		}
	}

	fmt.Println("Number of safe reports: ", safeReports)
}

func parttwo() {
	reports := parseInput()
	safeReports := 0

	for _, report := range reports {
		posAmount := 0
		negAmount := 0
		for i := 0; i < len(report)-1; i++ {
			if report[i] < report[i+1] {
				posAmount++
			} else if report[i] > report[i+1] {
				negAmount++
			}
		}
		var isPositive bool
		if posAmount > negAmount {
			isPositive = true
		} else if posAmount < negAmount {
			isPositive = false
		} else {
			continue
		}
		unsafeValues := 0
		for i := 0; i < len(report)-1; i++ {
			if !isSafeLevelChange(report[i], report[i+1], isPositive) {
				unsafeValues++
				if unsafeValues >= 2 || i == len(report)-2 {
					break
				}
				if !isSafeLevelChange(report[i], report[i+2], isPositive) {
					unsafeValues++
					break
				}
				i++
			}
		}
		if unsafeValues < 2 {
			safeReports++
		}
	}

	fmt.Println("Number of safe reports: ", safeReports)
}
