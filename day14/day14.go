package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Vector struct {
	x int64
	y int64
}

func (v Vector) getAllAdjacentTiles() []Vector {
	return []Vector{
		{
			x: v.x - 1,
			y: v.y - 1,
		},
		{
			x: v.x,
			y: v.y - 1,
		},
		{
			x: v.x + 1,
			y: v.y - 1,
		},
		{
			x: v.x - 1,
			y: v.y,
		},
		{
			x: v.x + 1,
			y: v.y,
		},
		{
			x: v.x - 1,
			y: v.y + 1,
		},
		{
			x: v.x,
			y: v.y + 1,
		},
		{
			x: v.x + 1,
			y: v.y + 1,
		},
	}
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

func moveRobot(robot *Robot, boardSize *Vector) {
	robot.location.x += robot.velocity.x
	robot.location.y += robot.velocity.y

	if robot.location.x < 0 {
		robot.location.x += boardSize.x
	} else if robot.location.x >= boardSize.x {
		robot.location.x -= boardSize.x
	}

	if robot.location.y < 0 {
		robot.location.y += boardSize.y
	} else if robot.location.y >= boardSize.y {
		robot.location.y -= boardSize.y
	}
}

func areRobotsConnected(robotLocations map[Vector]int) bool {
	totalLocations := len(robotLocations)
	for key := range robotLocations {
		return totalLocations == numConnections(&key, &robotLocations)
	}
	return false
}

func robotLengths(robotLocations map[Vector]int) {
	for key := range robotLocations {
		length := len(robotLocations)
		robotLen := numConnections(&key, &robotLocations)

		fmt.Println("Key is: ", key)
		fmt.Println("Key connections is: ", robotLen)
		fmt.Println("Total len is: ", length)
		return
	}
}

func keyLarge(robotLocations map[Vector]int) bool {
	for key := range robotLocations {
		return numConnections(&key, &robotLocations) > 100
	}
	return false
}

func numConnections(loc *Vector, robotLocations *map[Vector]int) int {
	delete(*robotLocations, *loc)
	connections := 1

	for _, newLoc := range loc.getAllAdjacentTiles() {
		if _, match := (*robotLocations)[newLoc]; match {
			connections += numConnections(&newLoc, robotLocations)
		}
	}

	return connections
}

func printBoard(robotLocations *map[Vector]int, boardSize *Vector) {
	for row := 0; row < int(boardSize.x); row++ {
		rowStr := ""
		for col := 0; col < int(boardSize.y); col++ {
			pos := Vector{
				x: int64(row),
				y: int64(col),
			}
			if _, match := (*robotLocations)[pos]; match {
				rowStr += "#"
			} else {
				rowStr += "."
			}
		}
		fmt.Println(rowStr)
	}
}

func parttwo() {
	robots := parseInput()
	boardSize := Vector{
		x: 101,
		y: 103,
	}

	robotLocations := make(map[Vector]int)

	for _, robot := range robots {
		robotLocations[robot.location]++
	}
	seconds := 0

	for !areRobotsConnected(robotLocations) {
		seconds++
		for index := range robots {
			robotLocations[robots[index].location]--
			if robotLocations[robots[index].location] <= 0 {
				delete(robotLocations, robots[index].location)
			}
			moveRobot(&robots[index], &boardSize)
			robotLocations[robots[index].location]++
		}

		printBoard(&robotLocations, &boardSize)
		robotLengths(robotLocations)
		fmt.Println("Seconds: ", seconds)
		time.Sleep(50 * time.Millisecond)
		if keyLarge(robotLocations) {
			break
		}
	}

	printBoard(&robotLocations, &boardSize)
	fmt.Println("Seconds passed: ", seconds)
}
