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
	ranges := []Range{}
	for scanner.Scan() {
		arr := strings.SplitN(scanner.Text(), "-", 2)
		l, _ := strconv.Atoi(arr[0])
		u, _ := strconv.Atoi(arr[1])
		ranges = append(ranges, Range{int64(l), int64(u)})
	}
	fmt.Println(solve1(ranges))
}

func solve1(in []Range) string {
	l := map[int64]int{}
	u := map[int64]int{}
	for _, r := range in {
		l[r.lower]++
		u[r.upper]++
	}
	sort.Slice(in, func(i, j int) bool {
		return in[i].lower < in[j].lower
	})
	allowed := 0
	index := 0
	for i := int64(0); i <= int64(4294967295); {
		r := in[index]
		if i >= r.lower {
			if i <= r.upper {
				i = r.upper + 1
				continue
			}
			index++
		} else {
			allowed++
			i++
		}

		if i%100000000 == 0 {
			fmt.Println(i)
		}
	}
	return strconv.Itoa(allowed)
}

type Range struct {
	lower, upper int64
}
