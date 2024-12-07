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

func isPositionEnd(board *[][]byte, pos Position) bool {
	b := *board
	return pos.row < 0 || pos.row >= len(b) || pos.col < 0 || pos.col >= len(b[0])
}

func isPositionOpen(board *[][]byte, pos Position) bool {
	b := *board
	return b[pos.row][pos.col] == '.'
}

func movePosition(board *[][]byte, pos Position) Position {
	nextPos := nextPosition(pos) 
	if isPositionEnd(board, nextPos) || isPositionOpen(board, nextPos) {
		return nextPos
	}
	return rotatePosition(pos)
}

func isLoopSpot(board *[][]byte, pos Position) bool {
	visited := make(map[string]bool)

	for !isPositionEnd(board, pos) {
		entry := longPositionString(pos)
		if _, match := visited[entry]; match {
			return true
		}
		visited[entry] = true
		pos = movePosition(board, pos)
	}

	return false
}

func partone() {
	board, pos := parseInput()
	spotsVisited := make(map[string]bool)

	for !isPositionEnd(&board, pos) {
		spotsVisited[shortPositionString(pos)] = true
		pos = movePosition(&board, pos)
	}

	fmt.Println("Spaces visited: ", len(spotsVisited))
}

func parttwo() {
	board, pos := parseInput()
	spotsVisited := make(map[string]bool)
	obstructions := make(map[string]bool)
	startPos := Position{pos.row, pos.col, pos.direction}

	for !isPositionEnd(&board, pos) {
		obsSpot := nextPosition(pos)
		_, obsHit := obstructions[shortPositionString(obsSpot)]

		if !isPositionEnd(&board, obsSpot) && isPositionOpen(&board, obsSpot) && !obsHit && (obsSpot.row != startPos.row || obsSpot.col != startPos.col) {
			board[obsSpot.row][obsSpot.col] = '#'
			if isLoopSpot(&board, startPos) {
				obstructions[shortPositionString(obsSpot)] = true
			}
			board[obsSpot.row][obsSpot.col] = '.'
		}

		spotsVisited[longPositionString(pos)] = true
		pos = movePosition(&board, pos)
	}

	fmt.Println("Obstructions to be placed: ", len(obstructions))
}
