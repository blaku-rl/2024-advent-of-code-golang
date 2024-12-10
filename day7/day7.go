package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Equation struct {
	targetNum uint
	numbers   []uint
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func parseInput() []Equation {
	equations := make([]Equation, 0, 1000)

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		var eq Equation
		tarNumsSplit := strings.Split(line, ":")
		tarNum, err := strconv.ParseUint(tarNumsSplit[0], 10, 64)
		if err == nil {
			eq.targetNum = uint(tarNum)
		}

		for _, numStr := range strings.Fields(strings.TrimSpace(tarNumsSplit[1])) {
			numStr = strings.TrimSpace(numStr)
			if numStr == "" {
				continue
			}

			num, err := strconv.ParseUint(numStr, 10, 32)
			if err == nil {
				eq.numbers = append(eq.numbers, uint(num))
			}
		}

		equations = append(equations, eq)
	}

	return equations
}

func isEquationConstructable(targetNum uint, nums []uint, curTotal uint) bool {
	if len(nums) == 1 {
		return curTotal+nums[0] == targetNum || curTotal*nums[0] == targetNum
	}

	return isEquationConstructable(targetNum, nums[1:], curTotal+nums[0]) || isEquationConstructable(targetNum, nums[1:], curTotal*nums[0])
}

func partone() {
	equations := parseInput()
	var totalValue uint

	for _, eq := range equations {
		if isEquationConstructable(eq.targetNum, eq.numbers[1:], eq.numbers[0]) {
			totalValue += eq.targetNum
		}
	}

	fmt.Println("Total value: ", totalValue)
}

func parttwo() {

}
