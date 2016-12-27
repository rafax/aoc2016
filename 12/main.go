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
	p := []assembunny.Instruction{}
	for scanner.Scan() {
		p = append(p, assembunny.ParseInstruction(scanner.Text()))
	}
	if os.Getenv("PART") != "1" {
		fmt.Println(solve1(p))
	} else {
		fmt.Println(solve2(p))
	}
}

func solve1(p assembunny.Program) int {
	c := assembunny.NewComputer()
	c.Execute(p)
	return c.GetRegister("a")
}

func solve2(p assembunny.Program) int {
	c := assembunny.NewComputer()
	c.SetRegister("c", 1)
	c.Execute(p)
	return c.GetRegister("a")
}
