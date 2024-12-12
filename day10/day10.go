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
}

type PointOnGrid struct {
	pos Position
	height byte
	checked bool
	rating int
	connectedPeaks map[Position]bool
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func parseInput() [][]PointOnGrid {
	topoGrid := make([][]PointOnGrid, 0)

	for row, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		rowOfGrid := make([]PointOnGrid, 0, len(line))
		for col := 0; col < len(line); col++ {
			height, err := strconv.Atoi(line[col:col+1])
			if err == nil {
				point := PointOnGrid {
					pos: Position{
						row: row,
						col: col,
					},
					height: byte(height),
					checked: false,
					rating: 0,
					connectedPeaks: make(map[Position]bool),
				}
				rowOfGrid = append(rowOfGrid, point)
			}
		}

		topoGrid = append(topoGrid, rowOfGrid)
	}

	return topoGrid
}

func possibleConnectedPaths(topoGrid *[][]PointOnGrid, point *PointOnGrid) []*PointOnGrid {
	points := make([]*PointOnGrid, 0)
	row := point.pos.row
	col := point.pos.col

	if point.pos.row > 0 {
		points = append(points, &((*topoGrid)[row-1][col]))
	}

	if point.pos.row < len((*topoGrid)) - 1 {
		points = append(points, &((*topoGrid)[row+1][col]))
	}

	if point.pos.col > 0 {
		points = append(points, &((*topoGrid)[row][col-1]))
	}

	if point.pos.col < len((*topoGrid)[row]) - 1 {
		points = append(points, &((*topoGrid)[row][col+1]))
	}

	return points
}

func calculatePaths(topoGrid *[][]PointOnGrid, point *PointOnGrid) (map[Position]bool, int) {
	if point.checked {
		return point.connectedPeaks, point.rating
	}

	if point.height == 9 {
		point.connectedPeaks[point.pos] = true
		point.checked = true
		point.rating = 1
		return point.connectedPeaks, point.rating
	}

	point.checked = true
	for _, nextPoint := range possibleConnectedPaths(topoGrid, point) {
		if nextPoint.height != point.height + 1 {
			continue
		}

		peaks, rating := calculatePaths(topoGrid, nextPoint)
		point.rating += rating

		for peak := range peaks {
			point.connectedPeaks[peak] = true
		}
	}

	return point.connectedPeaks, point.rating
}

func partone() {
	topoGrid := parseInput()

	for row := range topoGrid {
		for col := range topoGrid[row] {
			if topoGrid[row][col].checked {
				continue
			}

			calculatePaths(&topoGrid, &(topoGrid[row][col]))
		}
	}

	totalPaths := 0
	for row := range topoGrid {
		for col := range topoGrid[row] {
			if topoGrid[row][col].height != 0 {
				continue
			}

			totalPaths += len(topoGrid[row][col].connectedPeaks)
		}
	}

	fmt.Println("Total paths: ", totalPaths)
}

func parttwo() {
	topoGrid := parseInput()

	for row := range topoGrid {
		for col := range topoGrid[row] {
			if topoGrid[row][col].checked {
				continue
			}

			calculatePaths(&topoGrid, &(topoGrid[row][col]))
		}
	}

	totalRating := 0
	for row := range topoGrid {
		for col := range topoGrid[row] {
			if topoGrid[row][col].height != 0 {
				continue
			}

			totalRating += topoGrid[row][col].rating
		}
	}

	fmt.Println("Total rating: ", totalRating)
}
