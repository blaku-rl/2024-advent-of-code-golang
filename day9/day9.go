package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

type DiskItem struct {
	fileId int
	size   byte
}

//go:embed input.txt
var input string

func main() {
	fmt.Println("Running part one")
	partone()
	fmt.Println("Running part two")
	parttwo()
}

func parseInput() []byte {
	var diskLayout []byte
	for i := 0; i < len(input); i++ {
		if num, err := strconv.Atoi(input[i : i+1]); err == nil {
			diskLayout = append(diskLayout, byte(num))
		}
	}

	return diskLayout
}

func partone() {
	diskLayout := parseInput()
	diskMemory := make([]DiskItem, 0, len(diskLayout))

	for index, num := range diskLayout {
		if index%2 == 0 {
			diskMemory = append(diskMemory, DiskItem{index / 2, num})
		} else {
			diskMemory = append(diskMemory, DiskItem{-1, num})
		}
	}

	compressedMemory := make([]DiskItem, 0)
	for len(diskMemory) > 0 {
		fileToProcess := diskMemory[0]
		diskMemory = diskMemory[1:]

		if fileToProcess.fileId >= 0 {
			compressedMemory = append(compressedMemory, fileToProcess)
			continue
		}

		for fileToProcess.size > 0 && len(diskMemory) > 0 {
			lastFile := diskMemory[len(diskMemory)-1]
			diskMemory = diskMemory[:len(diskMemory)-1]

			if lastFile.fileId < 0 {
				continue
			}

			if fileToProcess.size >= lastFile.size {
				fileToProcess.size -= lastFile.size
				compressedMemory = append(compressedMemory, lastFile)
				continue
			}

			compressedMemory = append(compressedMemory, DiskItem{lastFile.fileId, fileToProcess.size})
			lastFile.size -= fileToProcess.size
			fileToProcess.size = 0
			diskMemory = append(diskMemory, lastFile)
		}
	}

	checksum := uint64(0)
	index := 0
	for _, item := range compressedMemory {
		checksum += uint64(item.fileId * ((int(item.size) * (index + (index + int(item.size) - 1))) / 2))
		index += int(item.size)
	}

	fmt.Println("Checksum value is: ", checksum)
}

func parttwo() {

}
