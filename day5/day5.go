package main

import (
	_ "embed"
	"fmt"
	"regexp"
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

func parseInput() ([]string, []string) {
	var rules []string
	var updates []string
	for _, row := range strings.Split(input, "\n") {
		if row == "" {
			continue
		}

		if strings.Contains(row, "|") {
			rules = append(rules, row)
		} else if strings.Contains(row, ",") {
			updates = append(updates, row)
		}
	}

	return rules, updates
}

func partone() {
	rules, updates := parseInput()
	var regRules []*regexp.Regexp
	total := 0

	for _, rule := range rules {
		nums := strings.Split(rule, "|")
		regStr := strings.TrimSpace(nums[1]) + ",.*" + strings.TrimSpace(nums[0])
		reg, err := regexp.Compile(regStr)
		if err == nil {
			regRules = append(regRules, reg)
		}
	}

	for _, update := range updates {
		isValid := true
		for _, reg := range regRules {
			match := reg.MatchString(update)
			if match {
				isValid = false
				break
			}
		}

		if isValid {
			nums := strings.Split(update, ",")
			num, err := strconv.Atoi(nums[len(nums)/2])
			if err == nil {
				total += num
			}
		}
	}

	fmt.Println("Total is: ", total)
}

func parttwo() {

}
