package assembunny

type Instruction interface {
	execute(c *Computer, i *int)
}

type Computer struct {
	registers map[string]int
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
	target string
	jumpBy int
}

func (i inc) execute(computer *Computer, _ *int) {
	computer.registers[i.target]++
}

func (d dec) execute(computer *Computer, _ *int) {
	computer.registers[d.target]--
}

func (j jnz) execute(computer *Computer, jumpBy *int) {
	if computer.registers[j.target] == 0 {
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

func ParseInstruction(in string) Instruction {
	return cpy{}
}
