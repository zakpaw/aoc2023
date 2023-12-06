package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
)

type Map struct {
    DestinationFrom int `parser:"@Int"`
    SourceFrom      int `parser:"@Int"`
    Length          int `parser:"@Int"`
}

type Converter struct {
    From  string `parser:"@Ident '-' 'to' '-'"`
    To    string `parser:"@Ident"`
    Maps  []Map  `parser:"'map' ':' @@+"`
}

func (c *Converter) getCorespondingValue() int {
	return 1
}

func main() {
	var seeds []int
	var resultP1, resultP2 int

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := strings.TrimPrefix(scanner.Text(), "seeds:")
		fmt.Sscan(line, &seeds)
        scanner.Scan()
	}

	parser := participle.MustBuild[Converter]()
	var converters []*Converter
	var converterData strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			converter, _ := parser.ParseString("", converterData.String())
            converters = append(converters, converter)
            fmt.Println(*converter)
			converterData.Reset()
		} else {
			converterData.WriteString(" " + line)
		}
	}
	fmt.Printf("part1: %d\npart2: %d\n", resultP1, resultP2)
}
