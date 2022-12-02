package day01

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jdrake/aoc2022/internal/util"
)

func Part1() {

	currCalories := 0
	maxCalories := 0

	handleLine := func(line string) {
		if strings.TrimSpace(line) == "" {
			currCalories = 0
			return
		}
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		currCalories += v
		if currCalories > maxCalories {
			maxCalories = currCalories
		}
	}
	util.ScanFile("/Users/jon.drake/repos/aoc2022/internal/day01/part1/input.txt", handleLine)
	fmt.Println(maxCalories)

}
