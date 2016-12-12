package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var res string
	if os.Getenv("PART") == "1" {
		res = part1("ffykfhsq")
	} else {
		res = part2("ffykfhsq")
	}
	fmt.Println(res)
}

func part1(prefix string) string {
	res := []string{}
	for i := 0; len(res) < 8; i++ {
		h := hash(prefix, strconv.Itoa(i))
		if strings.HasPrefix(h, "00000") {
			fmt.Println(h)
			res = append(res, string(h[5]))
		}
	}
	return strings.Join(res, "")
}

func part2(prefix string) string {
	pos := [8]string{}
	set := 0
	for i := 0; set < 8; i++ {
		h := hash(prefix, strconv.Itoa(i))
		if strings.HasPrefix(h, "00000") {
			p, err := strconv.Atoi(string(h[5]))
			if err == nil && p < 8 && pos[p] == "" {
				fmt.Printf("%v: %v (==%v) -> %v\n", h, p, h[5], string(h[6]))
				pos[p] = string(h[6])
				set++
			}
		}
	}
	return strings.Join(pos[:], "")
}

func hash(prefix, guess string) string {
	data := []byte(prefix + guess)
	return fmt.Sprintf("%x", md5.Sum(data))
}
