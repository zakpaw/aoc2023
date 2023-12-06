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

	resultP1 := math.MaxInt
	for _, seed := range seeds {
		result, _ := strconv.Atoi(seed)
		for _, converter := range converters {
			result = converter.getCorespondingValue(result)
		}
		if result < resultP1 {
			resultP1 = result
		}
	}
	fmt.Printf("part1: %d\npart2: %d\n", resultP1, -1)
}
