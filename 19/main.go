package main

import (
	"container/list"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Create("p2.prof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	var res string
	if os.Getenv("PART") == "1" {
		res = part1(3001330)
	} else {
		res = part2(3001330)
	}
	fmt.Println(res)
}

type elf struct {
	index, value int
}

func (e *elf) SetValue(v int) {
	e.value = v
}

func part1(cnt int) string {
	own := list.New()
	for i := 0; i < cnt; i++ {
		own.PushBack(elf{index: i, value: 1})
	}
	curr := own.Front()
	for {
		old := curr
		next := WrapNext(curr, own)
		updated := elf{value: curr.Value.(elf).value + next.Value.(elf).value, index: curr.Value.(elf).index}
		own.Remove(old)
		curr = own.InsertBefore(updated, next)
		own.Remove(next)
		if curr.Value.(elf).value == cnt {
			return fmt.Sprint(curr.Value.(elf).index + 1)
		}
		curr = WrapNext(curr, own)
	}
}

func part2(cnt int) string {
	own := list.New()
	for i := 0; i < cnt; i++ {
		own.PushBack(elf{index: i, value: 1})
	}
	curr := own.Front()
	next := curr
	for i := 0; i < own.Len()/2; i++ {
		next = WrapNext(next, own)
	}
	for i := 0; ; i++ {
		old := curr
		updated := elf{value: curr.Value.(elf).value + next.Value.(elf).value, index: curr.Value.(elf).index}
		// fmt.Printf("%+v stealing from %v\n", updated, next.Value)
		insertBefore := WrapNext(curr, own)
		curr = own.InsertBefore(updated, insertBefore)
		nnext := WrapNext(next, own)
		own.Remove(old)
		own.Remove(next)
		if own.Len()%2 == 0 {
			nnext = WrapNext(nnext, own)
		}
		if curr.Value.(elf).value == cnt {
			return fmt.Sprint(curr.Value.(elf).index + 1)
		}
		curr = WrapNext(curr, own)
		next = nnext
	}
}

func WrapNext(e *list.Element, l *list.List) *list.Element {
	if e.Next() == nil {
		return l.Front()
	}
	return e.Next()
}
