package main

import (
	"fmt"
	"os"
)

func main() {
	var res string
	if os.Getenv("PART") == "1" {
		res = solve(3001330)
	} else {

	}
	fmt.Println(res)
}
func solve(cnt int) string {
	own := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		own[i] = 1
	}
	i := 0
	for {
		next := (i + 1) % len(own)
		if own[i] > 0 {
			stealFrom := next
			for {
				if own[stealFrom] > 0 {
					own[i] += own[stealFrom]
					own[stealFrom] = 0
					break
				}

				stealFrom = (stealFrom + 1) % len(own)
			}
		}
		if own[i] == cnt {
			return fmt.Sprint(i + 1)
		}
		i = next
	}
}
