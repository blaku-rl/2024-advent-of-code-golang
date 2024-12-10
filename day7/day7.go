package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"math"
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

func twoOperatorCheck(targetNum uint, nums []uint, curTotal uint) bool {
	if len(nums) == 1 {
		return curTotal+nums[0] == targetNum || curTotal*nums[0] == targetNum
	}

	return twoOperatorCheck(targetNum, nums[1:], curTotal+nums[0]) || twoOperatorCheck(targetNum, nums[1:], curTotal*nums[0])
}

func concatNumbers(left, right uint) uint {
	power := int(1)
	for right % (uint(math.Pow10(power))) != right {
		power++
	}

	return (left * (uint(math.Pow10(power)))) + right
}

func threeOperatorCheck(targetNum uint, nums []uint, curTotal uint) bool {
	if curTotal > targetNum {
		return false
	}
	if len(nums) == 1 {
		return curTotal+nums[0] == targetNum || curTotal*nums[0] == targetNum || concatNumbers(curTotal, nums[0]) == targetNum
	}

	if threeOperatorCheck(targetNum, nums[1:], curTotal+nums[0]) {
		return true
	}
	if threeOperatorCheck(targetNum, nums[1:], curTotal*nums[0]) {
		return true
	}

	return threeOperatorCheck(targetNum, nums[1:], concatNumbers(curTotal, nums[0]))
}

func partone() {
	equations := parseInput()
	var totalValue uint

	for _, eq := range equations {
		if twoOperatorCheck(eq.targetNum, eq.numbers[1:], eq.numbers[0]) {
			totalValue += eq.targetNum
		}
	}

	fmt.Println("Total value: ", totalValue)
}

func parttwo() {
	equations := parseInput()
	var totalValue uint

	for _, eq := range equations {
		if threeOperatorCheck(eq.targetNum, eq.numbers[1:], eq.numbers[0]) {
			totalValue += eq.targetNum
		}
	}

	fmt.Println("Total value: ", totalValue)
}
