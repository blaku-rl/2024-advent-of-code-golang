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

func removeIndexFromSlice(curSlice []int, index int) []int {
	newSlice := make([]int, 0, len(curSlice)-1)
	newSlice = append(newSlice, curSlice[:index]...)
	newSlice = append(newSlice, curSlice[index+1:]...)
	return newSlice
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
		if isReportSafe(report) {
			safeReports++
			continue
		}
		for i := range report {
			newReport := removeIndexFromSlice(report, i)
			if isReportSafe(newReport) {
				safeReports++
				break
			}
		}
	}

	fmt.Println("Number of safe reports: ", safeReports)
}
