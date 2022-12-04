package day02

import (
	"fmt"
	"strings"

	"github.com/jdrake/aoc2022/internal/util"
)

func Part1() {
	letters := make(map[string]int)
	for i, char := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		letters[string(char)] = i + 1
	}

	sum := 0
	handleLine := func(line string) {
		if strings.TrimSpace(line) == "" {
			return
		}
		compartment1 := make(map[string]bool)
		compartment2 := make(map[string]bool)
		chars := strings.Split(line, "")
		for i, char := range chars {
			if i < (len(chars) / 2) {
				compartment1[char] = true
			} else {
				compartment2[char] = true
			}
		}
		for char, _ := range compartment1 {
			_, found := compartment2[char]
			if found {
				sum += letters[char]
			}
		}
	}
	util.ScanFile("/Users/jon.drake/repos/aoc2022/internal/day03/input.txt", handleLine)
	fmt.Println(sum)

}
