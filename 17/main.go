package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := buildGrid(bufio.NewScanner(file))
	s := solution{}
	if os.Getenv("PART") == "1" {
		s.solve1(step{grid{m}, 1, 1, "veumntbg"})
		fmt.Println(s.solution[:len(s.solution)])
	} else {
		in := "veumntbg"
		s.solve2(step{grid{m}, 1, 1, in})
		fmt.Println(len(s.solution[len(in):]))
	}
}

type step struct {
	grid grid
	x, y int
	in   string
}

type grid struct {
	data [][]cell
}

func (g grid) isValid(x, y int) bool {
	return 0 <= x && x < len(g.data[0]) && 0 <= y && y < len(g.data)
}

type pos struct {
	x, y int
}

func (g grid) moves(x int, y int, hash string, in string) []step {
	u, d, l, r := hash[0], hash[1], hash[2], hash[3]
	steps := []step{}
	if g.isValid(x, y+1) && canMove(d, g.data[y+1][x]) {
		diff := calcDiff(g.data[y+1][x])
		steps = append(steps, step{g, x, y + diff, in + "D"})
	}
	if g.isValid(x+1, y) && canMove(r, g.data[y][x+1]) {
		diff := calcDiff(g.data[y][x+1])
		steps = append(steps, step{g, x + diff, y, in + "R"})
	}
	if g.isValid(x-1, y) && canMove(l, g.data[y][x-1]) {
		diff := calcDiff(g.data[y][x-1])
		steps = append(steps, step{g, x - diff, y, in + "L"})
	}
	if g.isValid(x, y-1) && canMove(u, g.data[y-1][x]) {
		diff := calcDiff(g.data[y-1][x])
		steps = append(steps, step{g, x, y - diff, in + "U"})
	}
	return steps
}

type solution struct {
	solution string
}

func (sol *solution) solve1(s step) {
	if s.grid.foundEnd(s.x, s.y) {
		if sol.solution == "" || len(s.in) < len(sol.solution) {
			sol.solution = s.in
			return
		}
	}
	if sol.solution != "" && len(s.in) > len(sol.solution) {
		return
	}
	h := hash(s.in)
	steps := s.grid.moves(s.x, s.y, h, s.in)
	for _, step := range steps {
		sol.solve1(step)
	}
}

func (g grid) foundEnd(x, y int) bool {
	return x >= 7 && y >= 7
}

func (sol *solution) solve2(s step) {
	if s.grid.foundEnd(s.x, s.y) {
		if sol.solution == "" || (len(s.in) > len(sol.solution) && !strings.HasPrefix(s.in, sol.solution)) {
			sol.solution = s.in
		}
		return
	}
	h := hash(s.in)
	steps := s.grid.moves(s.x, s.y, h, s.in)
	for _, step := range steps {
		sol.solve2(step)
	}

}

func calcDiff(c cell) int {
	if c == door {
		return 2
	}
	return 1
}

func canMove(dir byte, c cell) bool {
	return ((c == door && isOpen(dir)) || c == free || c == end)
}

func isOpen(r byte) bool {
	return r == 'b' || r == 'c' || r == 'd' || r == 'e' || r == 'f'
}

func buildGrid(scanner *bufio.Scanner) [][]cell {
	grid := [][]cell{}

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]cell, len(line))
		for i, l := range line {
			switch l {
			case '|':
				row[i] = door
			case '-':
				row[i] = door
			case '#':
				row[i] = wall
			case ' ':
				row[i] = free
			case 'V':
				row[i] = end
			case 'S':
				row[i] = free
			}
		}
		grid = append(grid, row)
	}
	return grid
}

type cell int

const (
	wall cell = iota
	door cell = 1
	free cell = 2
	end  cell = 3
)

func hash(in string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(in)))
}
