package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Vector struct {
	x int64
	y int64
}

type Robot struct {
	location Vector
	velocity Vector
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func parseInput() []Robot {
	robots := make([]Robot, 0)
	robotReg := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		match := robotReg.FindAllStringSubmatch(line, -1)
		robot := Robot{}

		robot.location.x, _ = strconv.ParseInt(match[0][1], 10, 64)
		robot.location.y, _ = strconv.ParseInt(match[0][2], 10, 64)
		robot.velocity.x, _ = strconv.ParseInt(match[0][3], 10, 64)
		robot.velocity.y, _ = strconv.ParseInt(match[0][4], 10, 64)

		robots = append(robots, robot)
	}

	return robots
}

func partone() {
	robots := parseInput()
	boardSize := Vector{
		x: 101,
		y: 103,
	}
	seconds := 100
	quadAmounts := make([]int, 4)

	for _, robot := range robots {
		finalLoc := Vector{
			x: (robot.location.x + (robot.velocity.x * int64(seconds))) % boardSize.x,
			y: (robot.location.y + (robot.velocity.y * int64(seconds))) % boardSize.y,
		}

		if finalLoc.x < 0 {
			finalLoc.x += boardSize.x
		}

		if finalLoc.y < 0 {
			finalLoc.y += boardSize.y
		}

		if finalLoc.x < boardSize.x/2 && finalLoc.y < boardSize.y/2 {
			quadAmounts[0]++
		}
		if finalLoc.x > boardSize.x/2 && finalLoc.y < boardSize.y/2 {
			quadAmounts[1]++
		}
		if finalLoc.x < boardSize.x/2 && finalLoc.y > boardSize.y/2 {
			quadAmounts[2]++
		}
		if finalLoc.x > boardSize.x/2 && finalLoc.y > boardSize.y/2 {
			quadAmounts[3]++
		}
	}

	totalScore := uint64(1)
	for _, quad := range quadAmounts {
		totalScore *= uint64(quad)
	}

	fmt.Println("Safety score: ", totalScore)
}

func parttwo() {

}
