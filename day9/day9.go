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

func getDiskMemory(diskLayout *[]byte) []DiskItem {
	diskMemory := make([]DiskItem, 0, len(*diskLayout))

	for index, num := range *diskLayout {
		if index%2 == 0 {
			diskMemory = append(diskMemory, DiskItem{index / 2, num})
		} else {
			diskMemory = append(diskMemory, DiskItem{-1, num})
		}
	}

	return diskMemory
}

func fragmentedCompression(diskMemory []DiskItem) []DiskItem {
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

	return compressedMemory
}

func smartCompression(diskMemory []DiskItem) []DiskItem {
	compressedMemory := make([]DiskItem, 0)
	fileMap := make(map[byte][]int)

	//loop in reverse to get items
	for i := len(diskMemory) - 1; i >= 0; i-- {
		file := diskMemory[i]
		if file.fileId == -1 {
			continue
		}

		if _, ok := fileMap[file.size]; !ok {
			fileMap[file.size] = make([]int, 0)
		}

		fileMap[file.size] = append(fileMap[file.size], i)
	}

	for i := 0; i < len(diskMemory); i++ {
		file := diskMemory[i]

		if file.size == 0 {
			continue
		}

		if file.fileId != -1 {
			compressedMemory = append(compressedMemory, file)
			continue
		}

		lastFile := struct {
			index int
			size byte
		} {
			index: i,
			size: 0,
		}

		for j := file.size; j >= 0 && j < 10; j-- {
			if len(fileMap[j]) == 0 {
				continue
			}

			potentialMovedFileIndex := fileMap[j][0]
			if lastFile.index < potentialMovedFileIndex {
				lastFile.index = potentialMovedFileIndex
				lastFile.size = j
			}
		}

		if lastFile.index > i {
			compressedMemory = append(compressedMemory, diskMemory[lastFile.index])
			diskMemory[lastFile.index].fileId = -1
			diskMemory[i].size -= lastFile.size
			fileMap[lastFile.size] = fileMap[lastFile.size][1:]
			i--
		} else {
			compressedMemory = append(compressedMemory, file)
		}
	}

	return compressedMemory
}

func checksumForFile(file DiskItem, curSum uint64, index int) (uint64, int) {
	if file.fileId >= 0 {
		curSum += uint64(file.fileId * ((int(file.size) * (index + (index + int(file.size) - 1))) / 2))
	}
	index += int(file.size)
	return curSum, index
}

func partone() {
	diskLayout := parseInput()
	diskMemory := getDiskMemory(&diskLayout)
	compressedMemory := fragmentedCompression(diskMemory)

	checksum := uint64(0)
	index := 0
	for _, item := range compressedMemory {
		checksum, index = checksumForFile(item, checksum, index)
	}

	fmt.Println("Checksum value is: ", checksum)
}

func parttwo() {
	diskLayout := parseInput()
	diskMemory := getDiskMemory(&diskLayout)
	compressedMemory := smartCompression(diskMemory)

	checksum := uint64(0)
	index := 0
	for _, item := range compressedMemory {
		checksum, index = checksumForFile(item, checksum, index)
	}

	fmt.Println("Checksum value is: ", checksum)
}
