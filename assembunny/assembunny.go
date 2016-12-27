package assembunny

import (
	"strconv"
	"strings"
)

type Instruction interface {
	execute(c *Computer, i *int)
}

type Computer struct {
	registers map[string]int
}

func (c *Computer) GetRegister(r string) int {
	return c.registers[r]
}

func NewComputer() Computer {
	return Computer{registers: map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}}
}

type Program []Instruction

func (c *Computer) Execute(p Program) {
	for i, diff := 0, 1; i < len(p); diff = 1 {
		p[i].execute(c, &diff)
		i += diff
	}
}

type cpy struct {
	from   string
	val    int
	target string
}

type inc struct {
	target string
}

type dec struct {
	target string
}

type jnz struct {
	cond   string
	jumpBy int
}

func (i inc) execute(computer *Computer, _ *int) {
	computer.registers[i.target]++
}

func (d dec) execute(computer *Computer, _ *int) {
	computer.registers[d.target]--
}

func (j jnz) execute(computer *Computer, jumpBy *int) {
	if (isRegister(j.cond) && computer.registers[j.cond] == 0) || j.cond == "0" {
		return
	}
	*jumpBy = j.jumpBy
}

func (c cpy) execute(computer *Computer, _ *int) {
	v, ok := computer.registers[c.from]
	if !ok {
		v = c.val
	}
	computer.registers[c.target] = v
}

func isRegister(v string) bool {
	return strings.Contains("abcd", v)
}

func ParseInstruction(in string) Instruction {
	parts := strings.SplitN(in, " ", 3)
	switch parts[0] {
	case "cpy":
		if isRegister(parts[1]) {
			return cpy{from: parts[1], target: parts[2]}
		}
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		return cpy{val: v, target: parts[2]}
	case "inc":
		return inc{target: parts[1]}
	case "dec":
		return dec{target: parts[1]}
	case "jnz":
		jb, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		return jnz{cond: parts[1], jumpBy: jb}
	default:
		panic("Unknown instruction " + in)
	}
}
