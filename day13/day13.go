package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Position struct {
	x uint64
	y uint64
}

type ClawMachine struct {
	aButton Position
	bButton Position
	prize   Position
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
	x, _ := strconv.ParseUint(strs[len(strs)-2], 10, 64)
	y, _ := strconv.ParseUint(strs[len(strs)-1], 10, 64)

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

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func positionsAreMultiples(pos1, pos2 *Position) bool {
	det := int64(pos1.x*pos2.y) - int64(pos1.y*pos2.x)
	return det == 0
}

func getOptimalTokenUsage(machine *ClawMachine) uint64 {
	if positionsAreMultiples(&machine.aButton, &machine.bButton) {
		if !positionsAreMultiples(&machine.aButton, &machine.prize) {
			return 0
		}

		fmt.Println("Linearlly dependent")
		return 0
	}

	denom := gcd(machine.aButton.x, machine.aButton.y)
	rowFactor1 := machine.aButton.y / denom
	rowFactor2 := machine.aButton.x / denom

	bPresses := int64(rowFactor2*machine.bButton.y) - int64(rowFactor1*machine.bButton.x)
	bResult := int64(rowFactor2*machine.prize.y) - int64(rowFactor1*machine.prize.x)

	if (bPresses < 0 && bResult > 0) || (bPresses > 0 && bResult < 0) {
		return 0
	}

	if bPresses < 0 {
		bPresses *= -1
		bResult *= -1
	}

	reduced := gcd(uint64(bPresses), uint64(bResult))
	bPresses /= int64(reduced)
	bResult /= int64(reduced)

	if bResult%bPresses != 0 {
		return 0
	}

	bPresses = bResult / bPresses
	top := int64(machine.prize.x) - (int64(machine.bButton.x) * bPresses)

	if top%int64(machine.aButton.x) != 0 {
		return 0
	}

	aPresses := top / int64(machine.aButton.x)
	return uint64(aPresses*3) + uint64(bPresses)
}

func partone() {
	machines := parseInput()
	optimalTokens := uint64(0)
	for _, machine := range machines {
		optimalTokens += getOptimalTokenUsage(&machine)
	}

	fmt.Println("Optimal tokens is: ", optimalTokens)
}

func parttwo() {
	machines := parseInput()
	optimalTokens := uint64(0)
	prizeMovement := uint64(10000000000000)

	for _, machine := range machines {
		machine.prize.x += prizeMovement
		machine.prize.y += prizeMovement
		optimalTokens += getOptimalTokenUsage(&machine)
	}

	fmt.Println("Optimal tokens is: ", optimalTokens)
}
