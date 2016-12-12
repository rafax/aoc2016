package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res string
	if os.Getenv("PART") == "1" {
		res = part1(scanner)
	} else {
		res = part2(scanner)
	}
	fmt.Println(res)
}

func part1(scanner *bufio.Scanner) string {
	counts := []map[rune]int{}
	scanner.Scan()
	line := scanner.Text()
	for i := 0; i < len(line); i++ {
		counts = append(counts, map[rune]int{})
	}
	for i, l := range line {
		counts[i][l]++
	}
	for scanner.Scan() {
		line := scanner.Text()
		for i, l := range line {
			counts[i][l]++
		}
	}
	return mostFrequent(counts)
}

func mostFrequent(counts []map[rune]int) string {
	res := []string{}
	for _, m := range counts {
		max := 0
		maxk := ""
		for k, v := range m {
			if v > max {
				max = v
				maxk = string(k)
			}
		}
		res = append(res, maxk)
	}
	return strings.Join(res, "")
}

func leastFrequent(counts []map[rune]int) string {
	res := []string{}
	for _, m := range counts {
		min := math.MaxInt32
		mink := ""
		for k, v := range m {
			if v < min {
				min = v
				mink = string(k)
			}
		}
		res = append(res, mink)
	}
	return strings.Join(res, "")
}

func part2(scanner *bufio.Scanner) string {
	counts := []map[rune]int{}
	scanner.Scan()
	line := scanner.Text()
	for i := 0; i < len(line); i++ {
		counts = append(counts, map[rune]int{})
	}
	for i, l := range line {
		counts[i][l]++
	}
	for scanner.Scan() {
		line := scanner.Text()
		for i, l := range line {
			counts[i][l]++
		}
	}
	return leastFrequent(counts)
}
