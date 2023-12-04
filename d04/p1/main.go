package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	name string
	n1   []int
	n2   []int
}

func main() {
	f, _ := os.Open("./input")
	defer f.Close()
	buff, _ := io.ReadAll(f)

	cards := parseCards(buff)
	sum := 0
	for _, card := range cards {
		points := 0
		for _, n := range card.n1 {
			if slices.Contains(card.n2, n) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		sum += points
	}
	fmt.Printf("Total points: %d\n", sum)
}

func parseCards(input []byte) []card {
	lines := strings.SplitAfter(string(input), "\n")
	var cards []card
	for _, line := range lines {
		l := strings.TrimSpace(line)
		s1 := strings.Split(l, ":")
		crd := card{name: strings.TrimSpace(s1[0])}
		s2 := strings.Split(s1[1], "|")
		n1 := strings.Split(s2[0], " ")
		for _, sn := range n1 {
			if len(strings.TrimSpace(sn)) == 0 {
				continue
			}
			n, _ := strconv.Atoi(sn)
			crd.n1 = append(crd.n1, n)
		}
		n2 := strings.Split(s2[1], " ")
		for _, sn := range n2 {
			if len(strings.TrimSpace(sn)) == 0 {
				continue
			}
			n, _ := strconv.Atoi(sn)
			crd.n2 = append(crd.n2, n)
		}
		cards = append(cards, crd)
	}
	return cards
}
