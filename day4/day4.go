package main

import (
	_ "embed"
	"fmt"
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

func parseInput() [][]byte {
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	var board [][]byte

	for _, row := range rows {
		board = append(board, []byte(row))
	}
	return board
}

func isWordPresent(board [][]byte, row int, col int, word string, direction string) int {
	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return 0
	}
	if board[row][col] != word[0] {
		return 0
	}
	if len(word) == 1 {
		return 1
	}
	word = word[1:]
	if direction == "N" {
		return isWordPresent(board, row-1, col, word, direction)
	} else if direction == "NE" {
		return isWordPresent(board, row-1, col+1, word, direction)
	} else if direction == "E" {
		return isWordPresent(board, row, col+1, word, direction)
	} else if direction == "SE" {
		return isWordPresent(board, row+1, col+1, word, direction)
	} else if direction == "S" {
		return isWordPresent(board, row+1, col, word, direction)
	} else if direction == "SW" {
		return isWordPresent(board, row+1, col-1, word, direction)
	} else if direction == "W" {
		return isWordPresent(board, row, col-1, word, direction)
	} else if direction == "NW" {
		return isWordPresent(board, row-1, col-1, word, direction)
	}
	return 0
}

func partone() {
	wordSearch := parseInput()
	total := 0

	for row := range wordSearch {
		for col := range wordSearch[row] {
			total += isWordPresent(wordSearch, row, col, "XMAS", "N")
			total += isWordPresent(wordSearch, row, col, "XMAS", "NE")
			total += isWordPresent(wordSearch, row, col, "XMAS", "E")
			total += isWordPresent(wordSearch, row, col, "XMAS", "SE")
			total += isWordPresent(wordSearch, row, col, "XMAS", "S")
			total += isWordPresent(wordSearch, row, col, "XMAS", "SW")
			total += isWordPresent(wordSearch, row, col, "XMAS", "W")
			total += isWordPresent(wordSearch, row, col, "XMAS", "NW")
		}
	}

	fmt.Println("Total xmas hits: ", total)
}

func parttwo() {
	wordSearch := parseInput()
	total := 0

	for row := range wordSearch {
		if row == 0 || row == len(wordSearch)-1 {
			continue
		}

		for col := range wordSearch[row] {
			if col == 0 || col == len(wordSearch[row])-1 {
				continue
			}

			if wordSearch[row][col] != 'A' {
				continue
			}

			topLeft := wordSearch[row-1][col-1]
			topRight := wordSearch[row-1][col+1]
			botLeft := wordSearch[row+1][col-1]
			botRight := wordSearch[row+1][col+1]

			if (topLeft != 'M' || botRight != 'S') && (topLeft != 'S' || botRight != 'M') {
				continue
			}

			if (topRight != 'M' || botLeft != 'S') && (topRight != 'S' || botLeft != 'M') {
				continue
			}

			total++
		}
	}

	fmt.Println("Total x-mas hits: ", total)
}
