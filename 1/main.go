package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	dir := scanner.Text()
	x, y := endPos(dir)
	fmt.Println(x, y)
	fmt.Println(int(math.Abs(float64(x)) + math.Abs(float64(y))))
}

func endPos(in string) (int, int) {
	visited := map[[2]int]bool{}
	moves := strings.Split(in, ", ")
	x, y := 0, 0
	dir := ao
	for _, m := range moves {
		if m[0] == 'R' {
			dir = dir.TurnRight()
		} else {
			dir = dir.TurnLeft()
		}
		steps, _ := strconv.Atoi(string(m[1:]))
		for i := 0; i < steps; i++ {
			switch dir {
			case T:
				y--
			case R:
				x++
			case D:
				y++
			case L:
				x--
			}
			if _, ok := visited[[2]int{x, y}]; !ok {
				visited[[2]int{x, y}] = true
			} else {
				return x, y
			}
		}

	}
	return x, y
}
