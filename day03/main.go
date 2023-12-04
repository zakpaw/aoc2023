package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isNumberAdjacent(lineAbove, currentLine, lineBelow string, symbolCol int) []int {
	var adjacentNumbers []int
	checkLine := func(line string, rowOffset int) {
		for _, match := range regexp.MustCompile(`\d+`).FindAllStringIndex(line, -1) {
			start, end := match[0], match[1]
			if (symbolCol >= start && symbolCol < end) ||
				(symbolCol-1 >= start && symbolCol-1 < end) ||
				(symbolCol+1 >= start && symbolCol+1 < end) {
				if num, err := strconv.Atoi(line[start:end]); err == nil {
					adjacentNumbers = append(adjacentNumbers, num)
				}
			}
		}
	}
	checkLine(lineAbove, -1)
	checkLine(currentLine, 0)
	checkLine(lineBelow, 1)
	return adjacentNumbers
}

func processLine(lineAbove, currentLine, lineBelow string, rowNum int) (allNumbers []int, gearNumbers []int) {
	for col, char := range currentLine {
		if char != '.' && (char < '0' || char > '9') {
			adjacentNumbers := isNumberAdjacent(lineAbove, currentLine, lineBelow, col)
			allNumbers = append(allNumbers, adjacentNumbers...)
			if char == '*' && len(adjacentNumbers) == 2 {
				gearNumbers = append(gearNumbers, adjacentNumbers[0] * adjacentNumbers[1])
			}
		}
	}
	return allNumbers, gearNumbers
}

func main() {
	var lineAbove, currentLine string
	var row, resultP1, resultP2 int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineBelow := scanner.Text()
		allNumbers, gearNumbers := processLine(lineAbove, currentLine, lineBelow, row-1)
		for _, num := range allNumbers {
			resultP1 += num
		}
		for _, num := range gearNumbers {
			resultP2 += num
		}
		lineAbove, currentLine = currentLine, lineBelow
		row++
	}
    fmt.Printf("part1: %d\npart2: %d\n", resultP1, resultP2)
}
