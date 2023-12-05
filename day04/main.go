package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "github.com/alecthomas/participle/v2"
    "github.com/deckarep/golang-set"
)

type Scratchcard struct {
    WinningNumbers  []interface{}   `parser:"@Int+"`
    ElfNumbers      []interface{}   `parser:"'|' @Int+"`
    WinCount        int
}

func (s *Scratchcard) checkNumberOfWins() {
    winningSet := mapset.NewSetFromSlice(s.WinningNumbers)
    elfSet := mapset.NewSetFromSlice(s.ElfNumbers)
    s.WinCount = winningSet.Intersect(elfSet).Cardinality()
}

func main() {
	parser := participle.MustBuild[Scratchcard]()
	scanner := bufio.NewScanner(os.Stdin)

	var resultP1 int
	var cards []*Scratchcard
	cardCopies := make(map[int]int)
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), ":")[1]
		card, _ := parser.ParseString("", numbers)
		card.checkNumberOfWins()
		if card.WinCount > 0 {
			resultP1 += 1 << (card.WinCount - 1)
		}
		cards = append(cards, card)
		cardCopies[len(cards)-1] = 1
	}

	resultP2 := 0
	for i, card := range cards {
		resultP2 += cardCopies[i]
		for winOffset := 1; winOffset <= card.WinCount; winOffset++ {
			if nextCardIndex := i + winOffset; nextCardIndex < len(cards) {
				cardCopies[nextCardIndex] += cardCopies[i]
			}
		}
	}
    fmt.Printf("part1: %d\npart2: %d\n", resultP1, resultP2)
}
