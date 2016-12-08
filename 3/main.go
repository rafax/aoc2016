package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	var cnt int
	if os.Getenv("PART") == "1" {
		cnt = part1(scanner)
	} else {
		cnt = part2(scanner)
	}
	fmt.Println(cnt)
}

func part1(scanner *bufio.Scanner) int {
	cnt := 0
	for scanner.Scan() {
		in := scanner.Text()
		sstr := strings.Split(in, " ")
		sides := []int{}
		for _, s := range sstr {
			side, _ := strconv.Atoi(s)
			sides = append(sides, side)
		}
		sort.Ints(sides)
		if sides[0]+sides[1] > sides[2] {
			cnt++
		}
	}
	return cnt
}
func part2(scanner *bufio.Scanner) int {
	cnt := 0
	for scanner.Scan() {
		in := scanner.Text()
		sstr := strings.Split(in, " ")
		sides := []int{}
		for _, s := range sstr {
			side, _ := strconv.Atoi(s)
			sides = append(sides, side)
		}
		sort.Ints(sides)
		if sides[0]+sides[1] > sides[2] {
			cnt++
		}
	}
	return cnt
}
