package day02

import (
	"fmt"
	"strings"

	"github.com/jdrake/aoc2022/internal/util"
)

func Part2() {
	sum := 0
	handleLine := func(line string) {
		if strings.TrimSpace(line) == "" {
			return
		}
		plays := strings.Split(line, " ")
		opponentPlay := plays[0]
		var myPlay string
		outcome := plays[1]
		score := 0
		switch outcome {
		case "X":
			// Lose
			switch opponentPlay {
			case "A":
				myPlay = "Z"
			case "B":
				myPlay = "X"
			case "C":
				myPlay = "Y"
			}
		case "Y":
			// Draw
			score += 3
			switch opponentPlay {
			case "A":
				myPlay = "X"
			case "B":
				myPlay = "Y"
			case "C":
				myPlay = "Z"
			}
		case "Z":
			// Win
			score += 6
			switch opponentPlay {
			case "A":
				myPlay = "Y"
			case "B":
				myPlay = "Z"
			case "C":
				myPlay = "X"
			}
		}
		sum += score + shapeScore[myPlay]
	}
	util.ScanFile("/Users/jon.drake/repos/aoc2022/internal/day02/input.txt", handleLine)
	fmt.Println(sum)

}
