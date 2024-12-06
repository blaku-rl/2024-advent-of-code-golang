package main

import (
	_ "embed"
	"fmt"
	"strconv"
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

func shortPositionString(pos Position) string {
	return strconv.Itoa(pos.row) + "," + strconv.Itoa(pos.col)
}

func longPositionString(pos Position) string {
	return shortPositionString(pos) + "," + string(pos.direction)
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

func isPositionOpen(board [][]byte, pos Position) bool {
	return board[pos.row][pos.col] == '.'
}

func movePosition(board [][]byte, pos Position) Position {
	nextPos := nextPosition(pos) 
	if isPositionEnd(board, nextPos) || isPositionOpen(board, nextPos) {
		return nextPos
	}
	return rotatePosition(pos)
}

func isLoopSpot(board [][]byte, pos Position, spotsVisited map[string]bool) bool {
	rotPos := rotatePosition(pos)

	for true {
		if isPositionEnd(board, rotPos) || !isPositionOpen(board, rotPos) {
			return false
		}

		if _, match := spotsVisited[longPositionString(rotPos)]; match {
			return true
		}

		rotPos = nextPosition(rotPos)
	}
	 return false
}

func partone() {
	board, pos := parseInput()
	spotsVisited := make(map[string]bool)

	for !isPositionEnd(board, pos) {
		spotsVisited[shortPositionString(pos)] = true
		pos = movePosition(board, pos)
	}

	fmt.Println("Spaces visited: ", len(spotsVisited))
}

func parttwo() {
	board, pos := parseInput()
	spotsVisited := make(map[string]bool)
	startPos := Position{}
	startPos.row = pos.row
	startPos.col = pos.col
	obstructions := make(map[string]bool)

	for !isPositionEnd(board, pos) {
		if isLoopSpot(board, pos, spotsVisited) {
			obstructSpot := nextPosition(pos)
			if obstructSpot.row != startPos.row || obstructSpot.col != startPos.col {
				obstructions[shortPositionString(obstructSpot)] = true
			}
		}

		spotsVisited[longPositionString(pos)] = true
		pos = movePosition(board, pos)
	}

	fmt.Println("Obstructions to be placed: ", len(obstructions))
}
