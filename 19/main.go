package main

import (
	"fmt"
	"os"
)

func main() {
	var res string
	if os.Getenv("PART") == "1" {
		res = solve("11111")
	} else {

	}
	fmt.Println(res)
}
func solve(in string) string {
	return ""
}
