package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
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
			rules = append(rules, strings.TrimSpace(row))
		} else if strings.Contains(row, ",") {
			updates = append(updates, strings.TrimSpace(row))
		}
	}

	return rules, updates
}

func generateRegexs(rules []string) []*regexp.Regexp {
	var regRules []*regexp.Regexp

	for _, rule := range rules {
		nums := strings.Split(rule, "|")
		regStr := nums[1] + ",.*" + nums[0]
		reg, err := regexp.Compile(regStr)
		if err == nil {
			regRules = append(regRules, reg)
		}
	}

	return regRules
}

func isUpdateValid(update string, regRules []*regexp.Regexp) bool {
	for _, reg := range regRules {
		if reg.MatchString(update) {
			return false
		}
	}
	return true
}

func getMidValue(update string) int {
	nums := strings.Split(update, ",")
	num, err := strconv.Atoi(nums[len(nums)/2])
	if err == nil {
		return num
	}
	return 0
}

func sortUpdate(update string, rulesMap map[string]bool) string {
	nums := strings.Split(update, ",")

	sort.Slice(nums, func(i, j int) bool {
		entry := nums[j] + "|" + nums[i]
		_, ok := rulesMap[entry]
		return !ok
	})

	var newStr string
	for _, num := range nums {
		newStr += num + ","
	}

	return newStr[:len(newStr) - 1]
}

func partone() {
	rules, updates := parseInput()
	regRules := generateRegexs(rules)
	total := 0

	for _, update := range updates {
		if isUpdateValid(update, regRules) {
			total += getMidValue(update)
		}
	}

	fmt.Println("Total is: ", total)
}

func parttwo() {
	rules, updates := parseInput()
	regRules := generateRegexs(rules)
	rulesMap := make(map[string]bool)
	total := 0

	for _, rule := range rules {
		rulesMap[rule] = true
	}

	for _, update := range updates {
		if isUpdateValid(update, regRules) {
			continue
		}

		sortedUpdate := sortUpdate(update, rulesMap)
		total += getMidValue(sortedUpdate)
	}

	fmt.Println("Total is: ", total)
}
