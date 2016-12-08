package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// var pos Position
	// if os.Getenv("PART") == "1" {
	// 	pos = Position{x: 2, y: 2, keyboard: keyboard1}
	// } else {
	// 	pos = Position{x: 1, y: 3, keyboard: keyboard2}
	// }
	for scanner.Scan() {

	}
}

type room struct {
	name     []rune
	sectorID int
}

type pair struct {
	k rune
	v int
}

func (r room) isReal() {
	cnt := map[rune]int{}
	for _, v := range r.name {
		cnt[v]++
	}
	pairs := []pair{}
	for k, v := range cnt {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].v == pairs[j].v {
			return pairs[i].k < pairs[j].k
		}
		return pairs[i].v < pairs[j].v
	})
	fmt.Println(pairs)
}
