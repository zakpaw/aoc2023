package main

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"os"
	"strconv"
	"strings"
)

type Leaderboard struct {
	Times     []int `parser:"'Time' ':' @Int+"`
	Distances []int `parser:"'Distance' ':' @Int+"`
}

func (l Leaderboard) getRaces() (races []Race) {
	for i := range l.Times {
		races = append(races, Race{l.Times[i], l.Distances[i]})
	}
	return
}

func (l Leaderboard) getP2Race() Race {
	var timeBuilder, distanceBuilder strings.Builder
	for i := range l.Times {
		timeBuilder.WriteString(strconv.Itoa(l.Times[i]))
		distanceBuilder.WriteString(strconv.Itoa(l.Distances[i]))
	}
	time, _ := strconv.Atoi(timeBuilder.String())
	distance, _ := strconv.Atoi(distanceBuilder.String())
	return Race{time, distance}
}

type Race struct {
	Time, Distance int
}

func (r Race) getWaysToWinCount() (count int) {
	for time := 0; time <= r.Time; time++ {
		if time*(r.Time-time) > r.Distance {
			count++
		}
	}
	return
}

func main() {
	data, _ := os.ReadFile("input.txt")
	leaderboard, _ := participle.MustBuild[Leaderboard]().ParseString("", string(data))

	resultP1, resultP2 := 1, 1
	for _, race := range leaderboard.getRaces() {
		resultP1 *= race.getWaysToWinCount()
	}

	resultP2 = leaderboard.getP2Race().getWaysToWinCount()
	fmt.Printf("part1: %d\npart2: %d\n", resultP1, resultP2)
}
