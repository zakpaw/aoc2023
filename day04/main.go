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
    Winning         int
}

func (s *Scratchcard) checkNumberOfWins() int {
    winningSet := mapset.NewSetFromSlice(s.WinningNumbers)
    elfSet := mapset.NewSetFromSlice(s.ElfNumbers)
    return winningSet.Intersect(elfSet).Cardinality()
}

func main() {
    parser := participle.MustBuild[Scratchcard]()
	scanner := bufio.NewScanner(os.Stdin)

    var resultP1 int
	for scanner.Scan() {
        numbers := strings.Split(scanner.Text(), ":")[1]
        card, _ := parser.ParseString("", numbers)

        winsCount := card.checkNumberOfWins()
        if winsCount > 0 {
			resultP1 += 1 << (winsCount - 1)
		}
	}
    fmt.Println(resultP1)
}
