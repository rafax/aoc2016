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

var keyboard = [][]*key{[]*key{nil, nil, nil, nil, nil}, []*key{nil, &key{"1"}, &key{"2"}, &key{"3"}, nil},
	[]*key{nil, &key{"4"}, &key{"5"}, &key{"6"}, nil}, []*key{nil, &key{"7"}, &key{"8"}, &key{"9"}, nil}, []*key{nil, nil, nil, nil, nil}}

type Position struct {
	x, y int
}

func (p Position) Move(d aoc.Direction) Position {
	switch d {
	case aoc.U:
		if keyboard[p.x][p.y-1] != nil {
			return Position{p.x, p.y - 1}
		}
	case aoc.R:
		if keyboard[p.x+1][p.y] != nil {
			return Position{p.x + 1, p.y}
		}
	case aoc.D:
		if keyboard[p.x][p.y+1] != nil {
			return Position{p.x, p.y + 1}
		}
	case aoc.L:
		if keyboard[p.x-1][p.y] != nil {
			return Position{p.x - 1, p.y}
		}
	default:
		panic("Unrecognized direction" + string(d))
	}
	return p
}

func (p Position) Key() string {
	return keyboard[p.y][p.x].string
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pos := Position{x: 2, y: 2}
	for scanner.Scan() {
		dir := scanner.Text()
		for _, in := range dir {
			d := aoc.ParseDirection(in)
			pos = pos.Move(d)
		}
		fmt.Print(pos.Key())
	}
}
