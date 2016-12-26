package assembunny

import "testing"

// func TestParseCopyValue(t *testing.T) {
// 	v, to := 42, "a"
// 	text := fmt.Sprintln("cpy", v, to)
// 	i := ParseInstruction(text)
// 	// if i.(copy) == nil {
// 	// 	t.Fatalf("Instruction should be a copy, got %+v", i)
// 	// }
// 	c := i.(copy)
// 	if c.from != "" {
// 		t.Error()
// 	}
// 	if c.val != v {
// 		t.Errorf("Expected val to equal %v got %v", v, c.val)
// 	}
// 	if c.target != to {
// 		t.Errorf("Expected target to equal %v got %v", to, c.target)
// 	}
// }

func TestCopyValue(t *testing.T) {
	reg := "a"
	expected := 42
	p := []Instruction{cpy{from: "", target: reg, val: expected}}
	computer := NewComputer()
	computer.Execute(p)
	actual := computer.registers[reg]
	if actual != expected {
		t.Errorf("Expected register %v to equal %v, got %v", reg, expected, actual)
	}
}

func TestCopyRegister(t *testing.T) {
	to := "a"
	from := "b"
	expected := 42
	computer := NewComputer()
	computer.registers[from] = expected
	p := []Instruction{cpy{from: from, target: to}}
	computer.Execute(p)
	actual := computer.registers[to]
	if actual != expected {
		t.Errorf("Expected register %v to equal %v, got %v", to, expected, actual)
	}
}

func TestInc(t *testing.T) {
	reg := "a"
	p := []Instruction{inc{target: reg}}
	computer := NewComputer()
	computer.Execute(p)
	actual := computer.registers[reg]
	if computer.registers["a"] != 1 {
		t.Errorf("Expected register %v to equal %v, got %v", reg, 1, actual)
	}
}

func TestDec(t *testing.T) {
	reg := "a"
	p := []Instruction{dec{target: reg}}
	computer := NewComputer()
	computer.Execute(p)
	actual := computer.registers[reg]
	if computer.registers["a"] != -1 {
		t.Errorf("Expected register %v to equal %v, got %v", reg, -1, actual)
	}
}

func TestJnzForZero(t *testing.T) {
	reg := "a"
	j := jnz{target: reg, jumpBy: 2}
	computer := NewComputer()
	i := 0
	j.execute(&computer, &i)
	if i != 0 {
		t.Errorf("Expected no jump so i should equal 0, got %v", i)
	}
}

func TestJnzJumpBack(t *testing.T) {
	reg := "a"
	j := jnz{target: reg, jumpBy: -2}
	computer := NewComputer()
	computer.registers[reg] = 123
	i := 0
	j.execute(&computer, &i)
	if i != j.jumpBy {
		t.Errorf("Expected jump by %v so i should equal %v, got %v", j.jumpBy, j.jumpBy, i)
	}
}

func TestJnzJumpForward(t *testing.T) {
	reg := "a"
	j := jnz{target: reg, jumpBy: 2}
	computer := NewComputer()
	computer.registers[reg] = 123
	i := 0
	j.execute(&computer, &i)
	if i != j.jumpBy {
		t.Errorf("Expected jump by %v so i should equal %v, got %v", j.jumpBy, j.jumpBy, i)
	}
}

func TestJnzJumpToEnd(t *testing.T) {
	reg := "a"
	p := []Instruction{jnz{target: reg, jumpBy: 10}, inc{target: reg}, inc{target: reg}}
	computer := NewComputer()
	computer.registers[reg] = 1

	computer.Execute(p)

	actual := computer.registers[reg]
	if computer.registers[reg] != 1 {
		t.Errorf("Expected register %v to equal %v, got %v", reg, 1, actual)
	}
}

func TestExecutionNoParsing(t *testing.T) {
	reg := "a"
	p := []Instruction{cpy{val: 41, target: reg}, inc{target: reg}, inc{target: reg}, dec{target: reg}, jnz{target: reg, jumpBy: 2}, dec{target: reg}}
	c := NewComputer()
	c.Execute(p)
	if c.registers[reg] != 42 {
		t.Errorf("Expected register %v to equal %v, got %v", reg, 42, c.registers[reg])
	}
}
