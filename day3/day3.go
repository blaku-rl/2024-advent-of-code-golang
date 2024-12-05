package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Operation struct {
	op    string
	pos   int
	value int
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func partone() {
	var exp = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	total := 0
	matches := exp.FindAllSubmatch([]byte(input), -1)
	for _, match := range matches {
		num1, err1 := strconv.Atoi(string(match[1]))
		num2, err2 := strconv.Atoi(string(match[2]))
		if err1 == nil && err2 == nil {
			total += num1 * num2
		}
	}

	fmt.Println("Our total is: ", total)
}

func parttwo() {
	var mulExp = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	var doExp = regexp.MustCompile(`do\(\)`)
	var dontExp = regexp.MustCompile(`don't\(\)`)
	var operations []Operation

	indicies := mulExp.FindAllIndex([]byte(input), -1)
	matches := mulExp.FindAllSubmatch([]byte(input), -1)
	for i, match := range matches {
		num1, err1 := strconv.Atoi(string(match[1]))
		num2, err2 := strconv.Atoi(string(match[2]))
		if err1 == nil && err2 == nil {
			operations = append(operations, Operation{"mul", indicies[i][0], num1 * num2})
		}
	}

	indicies = doExp.FindAllIndex([]byte(input), -1)
	for i := range indicies {
		operations = append(operations, Operation{"do", indicies[i][0], 0})
	}

	indicies = dontExp.FindAllIndex([]byte(input), -1)
	for i := range indicies {
		operations = append(operations, Operation{"dont", indicies[i][0], 0})
	}

	sort.Slice(operations, func(i, j int) bool {
		return operations[i].pos < operations[j].pos
	})

	isValid := true
	total := 0

	for _, op := range operations {
		if op.op == "do" {
			isValid = true
		} else if op.op == "dont" {
			isValid = false
		} else if op.op == "mul" && isValid {
			total += op.value
		}
	}

	fmt.Println("Our total is: ", total)
}
