package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/rafax/aoc2016/assembunny"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if os.Getenv("PART") != "1" {
		fmt.Println(solve1(scanner))
	}
}

func solve1(scanner *bufio.Scanner) int {
	p := []assembunny.Instruction{}
	for scanner.Scan() {
		p = append(p, assembunny.ParseInstruction(scanner.Text()))
	}
	c := assembunny.NewComputer()
	c.Execute(p)
	return c.GetRegister("a")
}
