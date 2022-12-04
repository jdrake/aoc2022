package day02

import (
	"fmt"
	"strings"

	"github.com/jdrake/aoc2022/internal/util"
)

func Part2() {
	count := 0
	handleLine := func(line string) {
		if strings.TrimSpace(line) == "" {
			return
		}
		r := ParseRange(line)
		if r[0][0] <= r[1][0] {
			if r[0][1] >= r[1][0] {
				count += 1
			}
		} else {
			if r[1][0] <= r[0][0] {
				if r[1][1] >= r[0][0] {
					count += 1
				}
			}
		}
	}
	util.ScanFile("/Users/jon.drake/repos/aoc2022/internal/day04/input.txt", handleLine)
	fmt.Println(count)

}
