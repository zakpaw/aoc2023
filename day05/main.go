package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/participle/v2"
)

type Map struct {
	DestinationFrom, SourceFrom, Length int `parser:"@Int"`
}

type Converter struct {
	From string `parser:"@Ident '-' 'to' '-'"`
	To   string `parser:"@Ident"`
	Maps []Map  `parser:"'map' ':' @@+"`
}

func (c *Converter) getCorespondingValue(value int) int {
	for _, m := range c.Maps {
		if value >= m.SourceFrom && value < m.SourceFrom+m.Length {
			offset := value - m.SourceFrom
			return m.DestinationFrom + offset
		}
	}
	return value
}

func parseConvertersStream(scanner *bufio.Scanner) []*Converter {
	parser := participle.MustBuild[Converter]()
	var converters []*Converter
	var converterData strings.Builder

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			converter, _ := parser.ParseString("", converterData.String())
			converters = append(converters, converter)
			converterData.Reset()
		} else {
			converterData.WriteString(line + " ")
		}
	}

	converter, _ := parser.ParseString("", converterData.String())
	return append(converters, converter)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := strings.TrimPrefix(scanner.Text(), "seeds:")
	seeds := strings.Fields(line)

	converters := parseConvertersStream(scanner)

	resultP1, resultP2 := math.MaxInt, math.MaxInt
	for i, seedStr := range seeds {
		seedVal, _ := strconv.Atoi(seedStr)
		for _, converter := range converters {
			seedVal = converter.getCorespondingValue(seedVal)
		}
		if seedVal < resultP1 {
			resultP1 = seedVal
		}
		if i%2 == 1 {
			seedVal, _ = strconv.Atoi(seedStr)
			prevSeedVal, _ := strconv.Atoi(seeds[i-1])
			for j := prevSeedVal; j <= prevSeedVal+seedVal; j++ {
				convertedJ := j
				for _, converter := range converters {
					convertedJ = converter.getCorespondingValue(convertedJ)
				}
				if convertedJ < resultP2 {
					resultP2 = convertedJ
				}
			}
		}
	}
	fmt.Printf("part1: %d\npart2: %d\n", resultP1, resultP2)
}
