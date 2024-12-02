package main

import (
	_ "embed"
	"fmt"
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

func readFile() ([]int, []int) {
	firstList := make([]int, 0, 1000)
	secondList := make([]int, 0, 1000)

	for _, row := range strings.Split(input, "\n") {
		splits := strings.Fields(row)
		left, err1 := strconv.Atoi(splits[0])
		right, err2 := strconv.Atoi(splits[1])
		if err1 == nil && err2 == nil {
			firstList = append(firstList, left)
			secondList = append(secondList, right)
		}
	}

	return firstList, secondList
}

func partone() {
	firstList, secondList := readFile()
	sort.Ints(firstList)
	sort.Ints(secondList)

	totalDiff := 0

	for i, val := range firstList {
		curDiff := val - secondList[i]
		if curDiff < 0 {
			curDiff *= -1
		}
		totalDiff += curDiff
	}

	fmt.Println("The total difference is: ", totalDiff)
}

func parttwo() {
	firstList, secondList := readFile()
	similarities := make(map[int]int)

	for _, num := range secondList {
		similarities[num]++
	}

	totalScore := 0

	for _, num := range firstList {
		if val, ok := similarities[num]; !ok {
			continue
		} else {
			totalScore += val * num
		}
	}

	fmt.Println("Total score is: ", totalScore)
}
