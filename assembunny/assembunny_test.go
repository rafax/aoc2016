package assembunny

import (
	"fmt"
	"testing"
)

func TestParseCopyValue(t *testing.T) {
	v, to := 42, "a"
	text := fmt.Sprintf("cpy %v %v", v, to)
	i := ParseInstruction(text)
	c, ok := i.(cpy)
	if !ok {
		t.Fatalf("Instruction should be a copy, got %+v", i)
	}
	if c.from != "" {
		t.Error()
	}
	if c.val != v {
		t.Errorf("Expected val to equal %v got %v", v, c.val)
	}
	if c.target != to {
		t.Errorf("Expected target to equal %v got %v", to, c.target)
	}
}

func TestParseCopyTarget(t *testing.T) {
	v, to := "b", "a"
	text := fmt.Sprintf("cpy %v %v", v, to)
	i := ParseInstruction(text)
	c, ok := i.(cpy)
	if !ok {
		t.Fatalf("Instruction should be a copy, got %+v", i)
	}
	if c.from != v {
		t.Errorf("Expected from to equal %v got %v", v, c.val)
	}
	if c.target != to {
		t.Errorf("Expected target to equal %v got %v", to, c.target)
	}
}

func TestParseInc(t *testing.T) {
	to := "a"
	text := fmt.Sprintf("inc %v", to)
	i := ParseInstruction(text)
	c, ok := i.(inc)
	if !ok {
		t.Fatalf("Instruction should be a inc, got %+v", i)
	}
	if c.target != to {
		t.Errorf("Expected target to equal %v got %v", to, c.target)
	}
}

func TestParseDec(t *testing.T) {
	to := "a"
	text := fmt.Sprintf("dec %v", to)
	i := ParseInstruction(text)
	c, ok := i.(dec)
	if !ok {
		t.Fatalf("Instruction should be a dec, got %+v", i)
	}
	if c.target != to {
		t.Errorf("Expected target to equal %v got %v", to, c.target)
	}
}

func TestParseJnz(t *testing.T) {
	cond, jmp := "a", -3
	text := fmt.Sprintf("jnz %v %v", cond, jmp)
	i := ParseInstruction(text)
	c, ok := i.(jnz)
	if !ok {
		t.Fatalf("Instruction should be a jnz, got %+v", i)
	}
	if c.cond != cond {
		t.Errorf("Expected target to equal %v got %v", cond, c.cond)
	}
	if c.jumpBy != jmp {
		t.Errorf("Expected target to equal %v got %v", jmp, c.jumpBy)
	}
}

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

func TestJnzForZeroReg(t *testing.T) {
	reg := "a"
	j := jnz{cond: reg, jumpBy: 2}
	computer := NewComputer()
	i := 0
	j.execute(&computer, &i)
	if i != 0 {
		t.Errorf("Expected no jump so i should equal 0, got %v", i)
	}
}

func TestJnzForZeroValue(t *testing.T) {
	j := jnz{cond: "0", jumpBy: 2}
	computer := NewComputer()
	i := 0
	j.execute(&computer, &i)
	if i != 0 {
		t.Errorf("Expected no jump so i should equal 0, got %v", i)
	}
}

func TestJnzJumpBack(t *testing.T) {
	reg := "a"
	j := jnz{cond: reg, jumpBy: -2}
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
	j := jnz{cond: reg, jumpBy: 2}
	computer := NewComputer()
	computer.registers[reg] = 123
	i := 0
	j.execute(&computer, &i)
	if i != j.jumpBy {
		t.Errorf("Expected jump by %v so i should equal %v, got %v", j.jumpBy, j.jumpBy, i)
	}
}

func TestJnzJumpForwardForNonZeroValue(t *testing.T) {
	j := jnz{cond: "1", jumpBy: 2}
	computer := NewComputer()
	i := 0
	j.execute(&computer, &i)
	if i != j.jumpBy {
		t.Errorf("Expected jump by %v so i should equal %v, got %v", j.jumpBy, j.jumpBy, i)
	}
}

func TestJnzJumpToEnd(t *testing.T) {
	reg := "a"
	p := []Instruction{jnz{cond: reg, jumpBy: 10}, inc{target: reg}, inc{target: reg}}
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
	p := []Instruction{cpy{val: 41, target: reg}, inc{target: reg}, inc{target: reg}, dec{target: reg}, jnz{cond: reg, jumpBy: 2}, dec{target: reg}}
	c := NewComputer()
	c.Execute(p)
	if c.registers[reg] != 42 {
		t.Errorf("Expected register %v to equal %v, got %v", reg, 42, c.registers[reg])
	}
}
