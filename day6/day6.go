package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type Position struct {
	row int
	col int
	direction byte
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func parseInput() ([][]byte, Position){
	var board [][]byte
	var startPos Position

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		row := make([]byte, 0, len(line))
		for i := 0; i < len(line); i++ {
			if line[i] == '^' {
				startPos.row = len(board)
				startPos.col = len(row)
				startPos.direction = 'N'
				row = append(row, '.')
			} else {
				row = append(row, line[i])
			}
		}
		board = append(board, row)
	}

	return board, startPos
}

func nextPosition(pos Position) Position {
	switch pos.direction {
	case 'N': pos.row--
	case 'E': pos.col++
	case 'S': pos.row++
	case 'W': pos.col--
	}

	return pos
}

func rotatePosition(pos Position) Position {
	switch pos.direction {
	case 'N': pos.direction = 'E'
	case 'E': pos.direction = 'S'
	case 'S': pos.direction = 'W'
	case 'W': pos.direction = 'N'
	}

	return pos
}

func isPositionEnd(board [][]byte, pos Position) bool {
	return pos.row < 0 || pos.row >= len(board) || pos.col < 0 || pos.col >= len(board[0])
}

func movePosition(board [][]byte, pos Position) Position {
	nextPos := nextPosition(pos) 
	if isPositionEnd(board, nextPos) || board[nextPos.row][nextPos.col] == '.' {
		return nextPos
	}
	return rotatePosition(pos)
}

func partone() {
	board, pos := parseInput()
	spotsVisited := make(map[string]bool)

	for !isPositionEnd(board, pos) {
		spot := string(pos.row) + "," + string(pos.col)
		spotsVisited[spot] = true
		pos = movePosition(board, pos)
	}

	fmt.Println("Spaces visited: ", len(spotsVisited))
}

func parttwo() {

}
