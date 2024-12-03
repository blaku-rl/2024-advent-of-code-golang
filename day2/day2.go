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

func partone() {
	reports := parseInput()
	safeReports := 0

	for _, report := range reports {
		isPositive := report[0] < report[1]
		isSafe := true
		for i := range report {
			if i == len(report)-1 {
				break
			}
			if isPositive && (report[i] >= report[i+1] || report[i] < report[i+1]-3) {
				isSafe = false
				break
			}
			if !isPositive && (report[i] <= report[i+1] || report[i] > report[i+1]+3) {
				isSafe = false
				break
			}
		}
		if isSafe {
			safeReports++
		}
	}

	fmt.Println("Number of safe reports: ", safeReports)
}
