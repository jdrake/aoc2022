package day02

import (
	"fmt"
	"strings"

	"github.com/jdrake/aoc2022/internal/util"
)

var shapeScore = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var opponentMap = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

func Part1() {
	sum := 0
	handleLine := func(line string) {
		if strings.TrimSpace(line) == "" {
			return
		}
		plays := strings.Split(line, " ")
		opponentPlay := opponentMap[plays[0]]
		myPlay := plays[1]
		score := 0
		if opponentPlay == myPlay {
			score += 3
		} else if (opponentPlay == "X" && myPlay == "Y") ||
			(opponentPlay == "Y" && myPlay == "Z") ||
			(opponentPlay == "Z" && myPlay == "X") {
			score += 6
		}
		sum += score + shapeScore[myPlay]
	}
	util.ScanFile("/Users/jon.drake/repos/aoc2022/internal/day02/input.txt", handleLine)
	fmt.Println(sum)

}
