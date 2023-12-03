package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./input")
	defer f.Close()
	buff, _ := io.ReadAll(f)
	rows := strings.SplitAfter(string(buff), "\n")
	for i := 0; i < len(rows); i++ {
		rows[i] = strings.TrimSpace(rows[i])
	}
	doc := document{Rows: rows}
	doc.scanForSymbols()
	for _, s := range doc.Symbols {
		fmt.Printf("Found symbol: %+v\n", s)
	}
	doc.findPartNumbersNearSymbols()
	total := 0
	for _, pn := range doc.PartNumbers {
		fmt.Printf("Found part number: %+v\n", pn)
		total += pn.number
	}
	fmt.Printf("\nTotal sum of all part numbers is: %d\n", total)
}

type position struct {
	Row    int
	Column int
}

type symbol struct {
	position
	character string
}

type document struct {
	Rows        []string
	Symbols     []symbol
	PartNumbers []partNumber
}

func (d *document) scanForSymbols() []symbol {
	d.Symbols = make([]symbol, 0)
	for x, row := range d.Rows {
		for y := 0; y < len(row); y++ {
			c := row[y : y+1]
			if isSymbol(c) {
				d.Symbols = append(d.Symbols, symbol{
					position: position{
						Row:    x,
						Column: y,
					},
					character: c,
				})
			}
		}
	}
	return d.Symbols
}

func (d *document) findPartNumbersNearSymbols() []partNumber {
	d.PartNumbers = make([]partNumber, 0)
	for _, symbol := range d.Symbols {
		// check left
		if p, err := d.peekLeftPartNumber(symbol.position); err == nil {
			p.parseLeftOf(d)
			d.PartNumbers = append(d.PartNumbers, p)
		}
		// check right
		if p, err := d.peekRightPartNumber(symbol.position); err == nil {
			p.parseRightOf(d)
			d.PartNumbers = append(d.PartNumbers, p)
		}

		// check up
		if p, err := d.peekUpPartNumber(symbol.position); err == nil {
			p.parseLeftOf(d)
			p.parseRightOf(d)
			d.PartNumbers = append(d.PartNumbers, p)
		} else {
			// check up left
			if p, err := d.peekUpLeftPartNumber(symbol.position); err == nil {
				p.parseLeftOf(d)
				d.PartNumbers = append(d.PartNumbers, p)
			}
			// check up right
			if p, err := d.peekUpRightPartNumber(symbol.position); err == nil {
				p.parseRightOf(d)
				d.PartNumbers = append(d.PartNumbers, p)
			}
		}

		// check down
		if p, err := d.peekDownPartNumber(symbol.position); err == nil {
			p.parseLeftOf(d)
			p.parseRightOf(d)
			d.PartNumbers = append(d.PartNumbers, p)
		} else {
			// check down left
			if p, err := d.peekDownLeftPartNumber(symbol.position); err == nil {
				p.parseLeftOf(d)
				d.PartNumbers = append(d.PartNumbers, p)
			}
			// check down right
			if p, err := d.peekDownRightPartNumber(symbol.position); err == nil {
				p.parseRightOf(d)
				d.PartNumbers = append(d.PartNumbers, p)
			}
		}
	}
	return d.PartNumbers
}

func (d *document) isPositionInBounds(pos position) bool {
	if pos.Row < 0 || pos.Row > len(d.Rows) {
		return false
	}
	if pos.Column < 0 || pos.Column > len(d.Rows[pos.Row])-1 {
		return false
	}
	return true
}

func (d *document) getCharacterAt(pos position) string {
	if !d.isPositionInBounds(pos) {
		return ""
	}
	return d.Rows[pos.Row][pos.Column : pos.Column+1]
}

func (d *document) peekPartNumberAt(pos position) (partNumber, error) {
	if !d.isPositionInBounds(pos) {
		return partNumber{}, errors.New("out of bounds")
	}

	c := d.getCharacterAt(pos)
	if isDigit(c) {
		n, _ := strconv.Atoi(c) // we will parse the full value later
		return partNumber{
			position:  pos,
			strNumber: c,
			number:    n,
		}, nil
	}
	return partNumber{}, errors.New("no part number found here")
}
func (d *document) peekLeftPartNumber(currPos position) (partNumber, error) {
	newPos := position{
		Row:    currPos.Row,
		Column: currPos.Column - 1,
	}
	return d.peekPartNumberAt(newPos)
}

func (d *document) peekRightPartNumber(currPos position) (partNumber, error) {
	newPos := position{
		Row:    currPos.Row,
		Column: currPos.Column + 1,
	}
	return d.peekPartNumberAt(newPos)
}

func (d *document) peekUpPartNumber(currPos position) (partNumber, error) {
	newPos := position{
		Row:    currPos.Row - 1,
		Column: currPos.Column,
	}
	return d.peekPartNumberAt(newPos)
}

func (d *document) peekUpLeftPartNumber(currPos position) (partNumber, error) {
	newPos := position{
		Row:    currPos.Row - 1,
		Column: currPos.Column - 1,
	}
	return d.peekPartNumberAt(newPos)
}

func (d *document) peekUpRightPartNumber(currPos position) (partNumber, error) {
	newPos := position{
		Row:    currPos.Row - 1,
		Column: currPos.Column + 1,
	}
	return d.peekPartNumberAt(newPos)
}

func (d *document) peekDownPartNumber(currPos position) (partNumber, error) {
	newPos := position{
		Row:    currPos.Row + 1,
		Column: currPos.Column,
	}
	return d.peekPartNumberAt(newPos)
}

func (d *document) peekDownLeftPartNumber(currPos position) (partNumber, error) {
	newPos := position{
		Row:    currPos.Row + 1,
		Column: currPos.Column - 1,
	}
	return d.peekPartNumberAt(newPos)
}

func (d *document) peekDownRightPartNumber(currPos position) (partNumber, error) {
	newPos := position{
		Row:    currPos.Row + 1,
		Column: currPos.Column + 1,
	}
	return d.peekPartNumberAt(newPos)
}

func isSymbol(char string) bool {
	if len(char) != 1 {
		panic(fmt.Sprintf("expected only 1 character but got %q instead", char))
	}
	return !isDigit(char) && !isPeriod(char)
}

func isDigit(char string) bool {
	if len(char) != 1 {
		panic(fmt.Sprintf("expected only 1 character but got %q instead", char))
	}
	return strings.ContainsAny(char, "0123456789")
}

func isPeriod(char string) bool {
	if len(char) != 1 {
		panic(fmt.Sprintf("expected only 1 character but got %q instead", char))
	}
	return char == "."
}

type partNumber struct {
	position
	strNumber string
	number    int
}

func (pn *partNumber) parseLeftOf(doc *document) {
	for i := pn.Column; i > 0; i-- {
		c := doc.getCharacterAt(position{Row: pn.Row, Column: i - 1})
		if len(c) > 0 && isDigit(c) {
			pn.strNumber = c + pn.strNumber
			pn.number, _ = strconv.Atoi(pn.strNumber)
		} else {
			break
		}
	}
}

func (pn *partNumber) parseRightOf(doc *document) {
	for i := pn.Column; i > 0; i++ {
		c := doc.getCharacterAt(position{Row: pn.Row, Column: i + 1})
		if len(c) > 0 && isDigit(c) {
			pn.strNumber = pn.strNumber + c
			pn.number, _ = strconv.Atoi(pn.strNumber)
		} else {
			break
		}
	}
}
