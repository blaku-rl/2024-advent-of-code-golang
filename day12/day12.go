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

type Plot struct {
	pos       Position
	plant     byte
	perimeter byte
	checked   bool
}

type Region struct {
	plant byte
	plots map[Plot]bool
}

func (reg Region) Area() int {
	return len(reg.plots)
}

func (reg Region) Perimeter() int {
	totalPerimeter := 0
	for plot := range reg.plots {
		totalPerimeter += int(plot.perimeter)
	}
	return totalPerimeter
}

func (reg Region) FencePrice() int {
	return reg.Area() * reg.Perimeter()
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func parseInput() [][]Plot {
	plots := make([][]Plot, 0)

	for row, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		plotRow := make([]Plot, 0, len(line))
		for col := 0; col < len(line); col++ {
			plot := Plot{
				pos: Position{
					row: row,
					col: col,
				},
				plant:     line[col],
				perimeter: 0,
				checked:   false,
			}
			plotRow = append(plotRow, plot)
		}
		plots = append(plots, plotRow)
	}

	return plots
}

func getAdjacentPositions(pos Position) []Position {
	return []Position{
		{
			row: pos.row + 1,
			col: pos.col,
		},
		{
			row: pos.row - 1,
			col: pos.col,
		},
		{
			row: pos.row,
			col: pos.col + 1,
		},
		{
			row: pos.row,
			col: pos.col - 1,
		},
	}
}

func isNextPlotConnected(pos Position, plant byte, plots *[][]Plot) bool {
	row := pos.row
	col := pos.col
	return row >= 0 && row < len((*plots)) && col >= 0 && col < len((*plots)[row]) && (*plots)[row][col].plant == plant
}

func makeRegion(region *Region, plots *[][]Plot, curPlot *Plot) {
	if curPlot.checked {
		return
	}

	curPlot.checked = true

	for _, pos := range getAdjacentPositions(curPlot.pos) {
		if isNextPlotConnected(pos, curPlot.plant, plots) {
			makeRegion(region, plots, &((*plots)[pos.row][pos.col]))
		} else {
			curPlot.perimeter += 1
		}
	}

	region.plots[*curPlot] = true
}

func partone() {
	plots := parseInput()
	regions := make([]Region, 0)

	for row := range plots {
		for _, plot := range plots[row] {
			if plot.checked {
				continue
			}
			region := Region{
				plant: plot.plant,
				plots: make(map[Plot]bool),
			}

			makeRegion(&region, &plots, &plot)
			regions = append(regions, region)
		}
	}

	totalPrice := 0
	for _, region := range regions {
		totalPrice += region.FencePrice()
	}

	fmt.Println("Total fence price is: ", totalPrice)
}

func parttwo() {

}
