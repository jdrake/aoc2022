package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jdrake/aoc2022/internal/util"
)

func ParseRange(line string) [2][2]int {
	ranges := strings.Split(line, ",")
	range1String := strings.Split(ranges[0], "-")
	range2String := strings.Split(ranges[1], "-")
	range10, _ := strconv.Atoi(range1String[0])
	range11, _ := strconv.Atoi(range1String[1])
	range20, _ := strconv.Atoi(range2String[0])
	range21, _ := strconv.Atoi(range2String[1])
	return [2][2]int{
		{range10, range11},
		{range20, range21},
	}

}

func Part1() {
	count := 0
	handleLine := func(line string) {
		if strings.TrimSpace(line) == "" {
			return
		}
		r := ParseRange(line)
		if (r[0][0] >= r[1][0] && r[0][1] <= r[1][1]) || (r[1][0] >= r[0][0] && r[1][1] <= r[0][1]) {
			count += 1
		}
	}
	util.ScanFile("/Users/jon.drake/repos/aoc2022/internal/day04/input.txt", handleLine)
	fmt.Println(count)

}
