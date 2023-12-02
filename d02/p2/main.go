package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var bag = map[string]int{"red": 12, "green": 13, "blue": 14}

func main() {
	f, _ := os.Open("./input")
	defer f.Close()
	buff, _ := io.ReadAll(f)

	lines := strings.SplitAfter(string(buff), "\n")
	parsedLines := make(map[string][]map[string]int)
	for _, line := range lines {
		game, subsets := parseLine(line)
		parsedLines[game] = subsets
	}

	tpow := 0
	for _, subsets := range parsedLines {
		minSet := findMinimumSet(subsets)
		var numbers []int
		for _, n := range minSet {
			numbers = append(numbers, n)
		}
		pow := numbers[0]
		for i := 1; i < len(numbers); i++ {
			pow = pow * numbers[i]
		}
		tpow += pow
	}
	fmt.Printf("The sum of the power of these sets is: %d\n", tpow)
}

func parseLine(line string) (string, []map[string]int) {
	gamedata := strings.Split(line, ":")
	name := strings.TrimSpace(gamedata[0])
	data := strings.TrimSpace(gamedata[1])
	sets := strings.Split(data, ";")

	var subsets []map[string]int
	for _, set := range sets {
		subs := strings.Split(set, ",")
		subset := make(map[string]int)

		for _, sub := range subs {
			nc := strings.Split(strings.TrimSpace(sub), " ")
			number, _ := strconv.Atoi(nc[0])
			color := strings.TrimSpace(nc[1])
			subset[color] = number
		}
		subsets = append(subsets, subset)
	}
	return name, subsets
}

func findMinimumSet(subsets []map[string]int) map[string]int {
	minSet := make(map[string]int)
	for c, _ := range bag {
		minSet[c] = 0
	}
	for _, subset := range subsets {
		for c, n := range subset {
			if minSet[c] < n {
				minSet[c] = n
			}
		}
	}
	return minSet
}
