package main

import (
	"bufio"
	"fmt"
	"github.com/alecthomas/participle/v2"
	"os"
)

type Color struct {
	Count int    `parser:"@Int"`
	Name  string `parser:"@('red' | 'green' | 'blue')"`
}

type Game struct {
	ID   int      `parser:"'Game' @Int ':'"`
	Sets []*Color `parser:"@@ (',' @@ | ';' @@)*"`
}

var maxColors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	parser := participle.MustBuild[Game]()
	scanner := bufio.NewScanner(os.Stdin)

	var resultP1, resultP2 int
	for scanner.Scan() {
		game, _ := parser.ParseString("", scanner.Text())

		isValid := true
		minRequired := make(map[string]int)
		for _, color := range game.Sets {
			if color.Count > maxColors[color.Name] {
				isValid = false
			}
			if color.Count > minRequired[color.Name] {
				minRequired[color.Name] = color.Count
			}
		}
		if isValid {
			resultP1 += game.ID
		}
		resultP2 += minRequired["red"] * minRequired["green"] * minRequired["blue"]
	}
	fmt.Printf("part1: %d\npart2: %d\n", resultP1, resultP2)
}
