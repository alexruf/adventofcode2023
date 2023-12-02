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

	sum := 0
	for game, subsets := range parsedLines {
		id, r := validateGame(game, subsets)
		fmt.Printf("Game %d: %t\n", id, r)
		if r {
			sum += id
		}
	}
	fmt.Printf("Total sum of possible game IDs: %d\n", sum)
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

func validateGame(game string, subsets []map[string]int) (int, bool) {
	id, _ := strconv.Atoi(strings.Split(strings.TrimSpace(game), " ")[1])
	for _, subset := range subsets {
		for c, n := range subset {
			if bag[c] < n {
				return id, false
			}
		}
	}
	return id, true
}
