package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	aoc "github.com/rafax/aoc2016"
)

type key struct {
	string
}

var keyboard1 = [][]*key{[]*key{nil, nil, nil, nil, nil}, []*key{nil, &key{"1"}, &key{"2"}, &key{"3"}, nil},
	[]*key{nil, &key{"4"}, &key{"5"}, &key{"6"}, nil}, []*key{nil, &key{"7"}, &key{"8"}, &key{"9"}, nil}, []*key{nil, nil, nil, nil, nil}}

var keyboard2 = [][]*key{[]*key{nil, nil, nil, nil, nil, nil, nil},
	[]*key{nil, nil, nil, &key{"1"}, nil, nil, nil},
	[]*key{nil, nil, &key{"2"}, &key{"3"}, &key{"4"}, nil, nil},
	[]*key{nil, &key{"5"}, &key{"6"}, &key{"7"}, &key{"8"}, &key{"9"}, nil},
	[]*key{nil, nil, &key{"A"}, &key{"B"}, &key{"C"}, nil, nil},
	[]*key{nil, nil, nil, &key{"D"}, nil, nil, nil},
	[]*key{nil, nil, nil, nil, nil, nil, nil}}

type Position struct {
	x, y     int
	keyboard [][]*key
}

func (p Position) Move(d aoc.Direction) Position {
	switch d {
	case aoc.U:
		if p.keyboard[p.x][p.y-1] != nil {
			return Position{p.x, p.y - 1, p.keyboard}
		}
	case aoc.R:
		if p.keyboard[p.x+1][p.y] != nil {
			return Position{p.x + 1, p.y, p.keyboard}
		}
	case aoc.D:
		if p.keyboard[p.x][p.y+1] != nil {
			return Position{p.x, p.y + 1, p.keyboard}
		}
	case aoc.L:
		if p.keyboard[p.x-1][p.y] != nil {
			return Position{p.x - 1, p.y, p.keyboard}
		}
	default:
		panic("Unrecognized direction" + string(d))
	}
	return p
}

func (p Position) Key() string {
	return p.keyboard[p.y][p.x].string
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pos := Position{x: 1, y: 3, keyboard: keyboard2}
	for scanner.Scan() {
		dir := scanner.Text()
		for _, in := range dir {
			d := aoc.ParseDirection(in)
			pos = pos.Move(d)
		}
		fmt.Print(pos.Key())
	}
}
