package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const chars string = "123456789"

func main() {
	f, _ := os.Open("./input")
	defer f.Close()
	buff, _ := io.ReadAll(f)
	rows := strings.SplitAfter(string(buff), "\n")
	total := 0
	for i, row := range rows {
		v := parseValue(row)
		fmt.Printf("Row(%d): %q Value: %d\n", i+1, strings.ReplaceAll(row, "\n", ""), v)
		total += v
	}

	fmt.Printf("Total calibration value is:  %d\n", total)
}

func parseValue(input string) int {
	if len(strings.TrimSpace(input)) == 0 {
		return 0
	}
	iFirst := strings.IndexAny(input, chars)
	iLast := strings.LastIndexAny(input, chars)

	first := ""
	last := ""
	if iFirst > -1 {
		first = input[iFirst : iFirst+1]
	}
	if iLast > -1 {
		last = input[iLast : iLast+1]
	}

	num, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, last))

	return num
}
