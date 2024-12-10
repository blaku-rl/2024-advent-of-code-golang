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

func getAntinodes(first, second Position) (Position, Position) {
	rowDiff := first.row - second.row
	colDiff := first.col - second.col

	pos1 := Position{first.row + rowDiff, first.col + colDiff}
	pos2 := Position{second.row - rowDiff, second.col - colDiff}
	return pos1, pos2
}

func isPositionInGrid(grid [][]byte, pos Position) bool {
	return pos.row >= 0 && pos.row < len(grid) && pos.col >= 0 && pos.col < len(grid[0])
}

func partone() {
	grid := parseInput()
	antennas := make(map[byte][]Position)
	antinodes := make(map[Position]bool)

	for row := range grid {
		for col := range grid {
			antenna := grid[row][col]
			if antenna == '.' {
				continue
			}

			if _, ok := antennas[antenna]; !ok {
				antennas[antenna] = make([]Position, 0)
			}

			antennas[antenna] = append(antennas[antenna], Position{row, col})
		}
	}

	for _, nodes := range antennas {
		for first := range nodes {
			for second := first + 1; second < len(nodes); second++ {
				anti1, anti2 := getAntinodes(nodes[first], nodes[second])
				if isPositionInGrid(grid, anti1) {
					antinodes[anti1] = true
				}
				if isPositionInGrid(grid, anti2) {
					antinodes[anti2] = true
				}
			}
		}
	}

	fmt.Println("Total unique antinodes: ", len(antinodes))
}

func parttwo() {

}
