package aoc2016

type Direction int

const (
	U Direction = iota
	R Direction = 1
	D Direction = 2
	L Direction = 3
)

func (d Direction) TurnRight() Direction {
	if d == L {
		return U
	}
	return d + 1
}

func (d Direction) TurnLeft() Direction {
	if d == U {
		return L
	}
	return d - 1
}

func ParseDirection(in rune) Direction {
	switch in {
	case 'U':
		return U
	case 'R':
		return R
	case 'D':
		return D
	case 'L':
		return L
	default:
		panic("Unknown direction:" + string(in))
	}
}
