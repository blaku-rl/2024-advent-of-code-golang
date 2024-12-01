package main

import (
  "fmt"
  "log"
  "os"
  "bufio"
  "strings"
  "strconv"
  "sort"
)

func main() {
  fmt.Println("Running part one")
  partone()
  fmt.Println("Running part two")
  parttwo()
}

func readFile() ([]int, []int) {
  f, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }

  defer f.Close()

  scanner := bufio.NewScanner(f)

  firstList := make([]int, 0, 1000)
  secondList := make([]int, 0, 1000)

  for scanner.Scan() {
    splits := strings.Fields(scanner.Text())
    i, err1 := strconv.Atoi(splits[0])
    if err1 != nil {
      panic(err1)
    }
    firstList = append(firstList, i)
    j, err2 := strconv.Atoi(splits[1])
    if err2 != nil {
      panic(err2)
    }
    secondList = append(secondList, j)
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return firstList, secondList
}

func partone() {
  firstList, secondList := readFile()
  sort.Ints(firstList)
  sort.Ints(secondList)

  totalDiff := 0

  for i, val := range firstList {
    curDiff := val - secondList[i]
    if curDiff < 0 {
      curDiff *= -1
    }
    totalDiff += curDiff
  }

  fmt.Println("The total difference is: ", totalDiff)
}

func parttwo() {
  firstList, secondList := readFile()
  similarities := make(map[int]int)

  for _, num := range secondList {
    similarities[num]++
  }

  totalScore := 0

  for _, num := range firstList {
    if val, ok := similarities[num]; !ok {
      continue
    } else {
      totalScore += val * num
    }
  }

  fmt.Println("Total score is: ", totalScore)
}
