package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type Direction byte

const (
	North Direction = 0
	East Direction = 1
	South Direction = 2
	West Direction = 3
)

func (dir Direction) PerpendicularDirections() []Direction {
	if dir == North || dir == South {
		return []Direction {
			East,
			West,
		}
	}

	return []Direction {
		North,
		South,
	}
}

func GetAllDirections() []Direction {
	return []Direction{
		North,
		East,
		South,
		West,
	}
}

type Position struct {
	row int
	col int
}

func (pos Position) MoveDirection(dir Direction) Position {
	row := pos.row
	col := pos.col

	switch dir {
	case North: row -= 1
	case South: row += 1
	case East: col += 1
	case West: col -= 1
	}

	return Position{
		row: row,
		col: col,
	}
}

type Orientation struct {
	pos Position
	dir Direction
}

func (pos Position) GetAdjacentPositions() []Orientation {
	orientations := make([]Orientation, 0, 4)

	for _, dir := range GetAllDirections() {
		or := Orientation{
			pos: pos.MoveDirection(dir),
			dir: dir,
		}

		orientations = append(orientations, or)
	}

	return orientations
}

type Plot struct {
	pos       Position
	plant     byte
	perimeter map[Direction]bool
	checked   bool
}

type Region struct {
	plant byte
	plots map[Position]*Plot
}

func (reg Region) Area() int {
	return len(reg.plots)
}

func (reg Region) Perimeter() int {
	totalPerimeter := 0
	for _, plot := range reg.plots {
		totalPerimeter += int(len(plot.perimeter))
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
				perimeter: make(map[Direction]bool),
				checked:   false,
			}
			plotRow = append(plotRow, plot)
		}
		plots = append(plots, plotRow)
	}

	return plots
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

	for _, orientation := range curPlot.pos.GetAdjacentPositions() {
		pos := orientation.pos
		if isNextPlotConnected(pos, curPlot.plant, plots) {
			makeRegion(region, plots, &((*plots)[pos.row][pos.col]))
		} else {
			curPlot.perimeter[orientation.dir] = true
		}
	}

	region.plots[curPlot.pos] = curPlot
}

func removeSide(region *Region, curPlot *Plot, fenceDir Direction) {
	if _, hasConnectedSide := curPlot.perimeter[fenceDir]; !hasConnectedSide {
		return
	}

	delete(curPlot.perimeter, fenceDir)

	for _, dir := range fenceDir.PerpendicularDirections() {
		nextPos := curPlot.pos.MoveDirection(dir)
		if nextPlot, valid := region.plots[nextPos]; valid {
			removeSide(region, nextPlot, fenceDir)
		}
	}
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
				plots: make(map[Position]*Plot),
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
	plots := parseInput()
	regions := make([]Region, 0)

	for row := range plots {
		for _, plot := range plots[row] {
			if plot.checked {
				continue
			}
			region := Region{
				plant: plot.plant,
				plots: make(map[Position]*Plot),
			}

			makeRegion(&region, &plots, &plot)
			regions = append(regions, region)
		}
	}

	totalPrice := 0
	for _, region := range regions {
		totalSides := 0
		for _, plot := range region.plots {
			for len(plot.perimeter) > 0 {
				var fenceDir Direction
				for dir := range plot.perimeter {
					fenceDir = dir
					break
				}

				removeSide(&region, plot, fenceDir)
				totalSides++
			}
		}

		totalPrice += region.Area() * totalSides
	}

	fmt.Println("Total fence price is: ", totalPrice)
}
