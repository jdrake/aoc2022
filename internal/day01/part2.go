package day01

import (
	"container/heap"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jdrake/aoc2022/internal/util"
)

func Part2() {

	currCalories := 0
	i := 0

	pq := make(util.PriorityQueue, 0)
	heap.Init(&pq)

	handleLine := func(line string) {
		if strings.TrimSpace(line) == "" {
			heap.Push(&pq, &util.Item{
				Value:    fmt.Sprintf("%d", currCalories),
				Priority: currCalories,
				Index:    i,
			})
			currCalories = 0
			i += 1
			return
		}
		v, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		currCalories += v
	}

	util.ScanFile("/Users/jon.drake/repos/aoc2022/internal/day01/input.text", handleLine)

	sum := 0
	for i := 0; i < 3; i++ {
		item := heap.Pop(&pq).(*util.Item)
		sum += item.Priority
	}
	fmt.Println(sum)
}
