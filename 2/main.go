package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	aoc "github.com/rafax/aoc2016"
)

var keyboard = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

type Position struct {
	pos int
}

func (p Position) Move(d aoc.Direction) Position {
	switch d {
	case aoc.U:
		if p.pos < 3 {
			return p
		}
		return Position{p.pos - 3}
	case aoc.R:
		if p.pos == 2 || p.pos == 5 || p.pos == 8 {
			return p
		}
		return Position{p.pos + 1}
	case aoc.D:
		if p.pos > 5 {
			return p
		}
		return Position{p.pos + 3}
	case aoc.L:
		if p.pos == 0 || p.pos == 3 || p.pos == 6 {
			return p
		}
		return Position{p.pos - 1}
	default:
		panic("Unrecognized direction" + string(d))
	}
}

func (p Position) Key() string {
	return strconv.Itoa(p.pos + 1)
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pos := Position{pos: 4}
	for scanner.Scan() {
		dir := scanner.Text()
		for _, in := range dir {
			d := aoc.ParseDirection(in)
			pos = pos.Move(d)
		}
		fmt.Print(pos.Key())
	}
}
