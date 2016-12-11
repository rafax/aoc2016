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
	// var pos Position
	// if os.Getenv("PART") == "1" {
	// 	pos = Position{x: 2, y: 2, keyboard: keyboard1}
	// } else {
	// 	pos = Position{x: 1, y: 3, keyboard: keyboard2}
	// }
	sum := 0
	for scanner.Scan() {
		res := scanner.Text()
		r := parseRoom(res)
		if r.isReal() {
			sum += r.sectorID
		}
	}
	fmt.Println(sum)
}

type room struct {
	counts   map[rune]int
	sectorID int
	checksum string
}

type pair struct {
	k rune
	v int
}

func (r room) isReal() bool {
	pairs := []pair{}
	for k, v := range r.counts {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].v == pairs[j].v {
			return pairs[i].k < pairs[j].k
		}
		return pairs[i].v > pairs[j].v
	})
	f5 := []string{}
	for _, v := range pairs[:5] {
		f5 = append(f5, string(v.k))
	}
	sum := strings.Join(f5, "")
	fmt.Println(r.checksum, sum)
	return r.checksum == sum
}

func parseRoom(in string) room {
	cnt := map[rune]int{}
	parts := strings.Split(in, "-")
	for _, p := range parts[:len(parts)-1] {
		for _, v := range p {
			cnt[v]++
		}
	}
	end := strings.Split(parts[len(parts)-1], "[")
	id, _ := strconv.Atoi(end[0])
	return room{counts: cnt, sectorID: id, checksum: end[1][:len(end[1])-1]}
}
