package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type Position struct {
	row int
	col int
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func parseInput() [][]byte {
	var grid [][]byte

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		row := make([]byte, 0, len(line))
		for i := 0; i < len(line); i++ {
			row = append(row, line[i])
		}

		grid = append(grid, row)
	}

	return grid
}

func getAntennasMap(grid *[][]byte) map[byte][]Position {
	antennas := make(map[byte][]Position)

	for row := range *grid {
		for col := range (*grid)[row] {
			antenna := (*grid)[row][col]
			if antenna == '.' {
				continue
			}

			if _, ok := antennas[antenna]; !ok {
				antennas[antenna] = make([]Position, 0)
			}

			antennas[antenna] = append(antennas[antenna], Position{row, col})
		}
	}

	return antennas
}

func getAntinodePair(first, second Position) (Position, Position) {
	rowDiff := first.row - second.row
	colDiff := first.col - second.col

	pos1 := Position{first.row + rowDiff, first.col + colDiff}
	pos2 := Position{second.row - rowDiff, second.col - colDiff}
	return pos1, pos2
}

func isPositionInGrid(grid *[][]byte, pos Position) bool {
	return pos.row >= 0 && pos.row < len(*grid) && pos.col >= 0 && pos.col < len((*grid)[0])
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func getAntinodeList(grid *[][]byte, first, second Position) []Position {
	antinodes := make([]Position, 0)
	rowDiff := first.row - second.row
	colDiff := first.col - second.col

	denominator := gcd(rowDiff, colDiff)
	rowDiff /= denominator
	colDiff /= denominator

	antinodes = append(antinodes, first)

	curPos := Position{first.row + rowDiff, first.col + colDiff}
	for isPositionInGrid(grid, curPos) {
		antinodes = append(antinodes, curPos)
		curPos = Position{curPos.row + rowDiff, curPos.col + colDiff}
	}

	curPos = Position{first.row - rowDiff, first.col - colDiff}
	for isPositionInGrid(grid, curPos) {
		antinodes = append(antinodes, curPos)
		curPos = Position{curPos.row - rowDiff, curPos.col - colDiff}
	}

	return antinodes
}

func partone() {
	grid := parseInput()
	antennas := getAntennasMap(&grid)
	antinodes := make(map[Position]bool)

	for _, nodes := range antennas {
		for first := range nodes {
			for second := first + 1; second < len(nodes); second++ {
				anti1, anti2 := getAntinodePair(nodes[first], nodes[second])
				if isPositionInGrid(&grid, anti1) {
					antinodes[anti1] = true
				}
				if isPositionInGrid(&grid, anti2) {
					antinodes[anti2] = true
				}
			}
		}
	}

	fmt.Println("Total unique antinodes: ", len(antinodes))
}

func parttwo() {
	grid := parseInput()
	antennas := getAntennasMap(&grid)
	antinodes := make(map[Position]bool)

	for _, nodes := range antennas {
		for first := range nodes {
			for second := first + 1; second < len(nodes); second++ {
				nodeList := getAntinodeList(&grid, nodes[first], nodes[second])
				for _, antinode := range nodeList {
					antinodes[antinode] = true
				}
			}
		}
	}

	fmt.Println("Total unique antinodes: ", len(antinodes))
}
