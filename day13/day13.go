package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type ClawMachine struct {
	aButton Position
	bButton Position
	prize Position
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func extractPositionFromStringSlice(strs []string) Position {
	x, _ := strconv.Atoi(strs[len(strs) - 2])
	y, _ := strconv.Atoi(strs[len(strs) - 1])

	return Position{
		x: x,
		y: y,
	}
}

func parseInput() []ClawMachine {
	buttonRegex := regexp.MustCompile(`Button ([AB]): X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	machines := make([]ClawMachine, 0)
	claw := ClawMachine{}
	justUpdated := true

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			if !justUpdated {
				machines = append(machines, claw)
				claw = ClawMachine{}
				justUpdated = true
			}
			continue
		}
		justUpdated = false

		if buttonMatch := buttonRegex.FindAllStringSubmatch(line, -1); len(buttonMatch) > 0 {
			pos := extractPositionFromStringSlice(buttonMatch[0])
			if buttonMatch[0][1] == "A" {
				claw.aButton = pos
			} else {
				claw.bButton = pos
			}
		} else if prizeMatch := prizeRegex.FindAllStringSubmatch(line, -1); len(prizeMatch) > 0 {
			claw.prize = extractPositionFromStringSlice(prizeMatch[0])
		}
	}

	return machines
}

func isPositionReachable(curPos, goalPos *Position) bool {
	return curPos.x <= goalPos.x && curPos.y <= goalPos.y
}

func isGoalReached(curPos, goalPos *Position) bool {
	return curPos.x == goalPos.x && curPos.y == goalPos.y
}

func scoreForPresses(aPresses, bPresses int) int {
	return (3 * aPresses) + bPresses
}

func getOptimalPresses(machine *ClawMachine, maxPresses int) int {
	bestScore := math.MaxInt
	curPos := Position{}

	for aPresses := 0; aPresses < maxPresses; aPresses++ {
		for bPresses := 0; bPresses < maxPresses; bPresses++ {
			curPos.x = (machine.aButton.x * aPresses) + (machine.bButton.x * bPresses)
			curPos.y = (machine.aButton.y * aPresses) + (machine.bButton.y * bPresses)

			if !isPositionReachable(&curPos, &machine.prize) {
				break
			}

			if isGoalReached(&curPos, &machine.prize) {
				score := scoreForPresses(aPresses, bPresses)
				if score < bestScore {
					bestScore = score
					break
				}
			}
		}
	}

	if bestScore == math.MaxInt {
		return 0
	}
	return bestScore
}

func partone() {
	machines := parseInput()
	optimalTokens := 0
	for _, machine := range machines {
		optimalTokens += getOptimalPresses(&machine, 100)
	}

	fmt.Println("Optimal tokens is: ", optimalTokens)
}

func parttwo() {

}
