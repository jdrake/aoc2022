package day02

import (
	"fmt"
	"strings"

	"github.com/jdrake/aoc2022/internal/util"
)

func Part2() {
	letters := make(map[string]int)
	for i, char := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		letters[string(char)] = i + 1
	}

	sum := 0
	group := 0
	seenChars := [3]map[string]int{
		make(map[string]int),
		make(map[string]int),
		make(map[string]int),
	}
	handleLine := func(line string) {
		if strings.TrimSpace(line) == "" {
			return
		}
		chars := strings.Split(line, "")
		for _, char := range chars {
			seenChars[group][char] = 1
		}
		group += 1
		if group%3 == 0 {
			for char := range seenChars[0] {
				_, ok1 := seenChars[1][char]
				_, ok2 := seenChars[2][char]
				if ok1 && ok2 {
					sum += letters[char]
				}
			}

			group = 0
			seenChars = [3]map[string]int{
				make(map[string]int),
				make(map[string]int),
				make(map[string]int),
			}
		}
	}
	util.ScanFile("/Users/jon.drake/repos/aoc2022/internal/day03/input.txt", handleLine)
	fmt.Println(sum)

}
