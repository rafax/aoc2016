package utils

type Direction int

const (
	T Direction = iota
	R Direction = 1
	D Direction = 2
	L Direction = 3
)

func (d Direction) TurnRight() Direction {
	if d == L {
		return T
	}
	return d + 1
}

func (d Direction) TurnLeft() Direction {
	if d == T {
		return L
	}
	return d - 1
}
