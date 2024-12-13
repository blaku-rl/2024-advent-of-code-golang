package main

import (
	_ "embed"
	"fmt"
	"math/big"
	"strings"
)

type Stone struct {
	stoneStr string
	blinks   int
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func parseInput() []string {
	nums := make([]string, 0)

	for _, numStr := range strings.Fields(strings.TrimSpace(input)) {
		if numStr == "" {
			continue
		}

		nums = append(nums, numStr)
	}

	return nums
}

func getNumStones(numStr string, blinksToGo int, stonesSeen *map[Stone]uint64) uint64 {
	curStone := Stone{numStr, blinksToGo}
	if stones, ok := (*stonesSeen)[curStone]; ok {
		return stones
	}

	if blinksToGo == 0 {
		(*stonesSeen)[curStone] = 1
	} else if numStr == "0" {
		(*stonesSeen)[curStone] = getNumStones("1", blinksToGo-1, stonesSeen)
	} else if len(numStr)%2 == 0 {
		left := getNumStones(numStr[:len(numStr)/2], blinksToGo-1, stonesSeen)
		rightNum := numStr[len(numStr)/2:]
		for len(rightNum) > 0 && rightNum[0] == '0' {
			rightNum = rightNum[1:]
		}
		if len(rightNum) == 0 {
			rightNum = "0"
		}
		right := getNumStones(rightNum, blinksToGo-1, stonesSeen)
		(*stonesSeen)[curStone] = right + left
	} else {
		var mulResult big.Int
		num, _ := new(big.Int).SetString(numStr, 0)
		mulResult.Mul(num, big.NewInt(int64(2024)))

		(*stonesSeen)[curStone] = getNumStones(mulResult.String(), blinksToGo-1, stonesSeen)
	}

	return (*stonesSeen)[curStone]
}

func partone() {
	nums := parseInput()

	totalStones := uint64(0)
	stonesMap := make(map[Stone]uint64)
	for _, num := range nums {
		totalStones += getNumStones(num, 25, &stonesMap)
	}

	fmt.Println("Total number of stones: ", totalStones)
}

func parttwo() {
	nums := parseInput()

	totalStones := uint64(0)
	stonesMap := make(map[Stone]uint64)
	for _, num := range nums {
		totalStones += getNumStones(num, 75, &stonesMap)
	}

	fmt.Println("Total number of stones: ", totalStones)
}
