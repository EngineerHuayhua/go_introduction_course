package function

import "fmt"

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(a int, b int) int {
	return a + b
}

func RepeatString(value string, increment int) {

	for i := 0; i < increment; i++ {
		fmt.Println(value)
	}
}

func Calc(op Operation, x, y float64) (float64, bool) {
	switch op {
	case SUM:
		return x + y, false
	case SUB:
		return x - y, false
	case DIV:
		if y == 0 {
			return 0, true
		}
		return x / y, false
	case MUL:
		return x * y, false
	default:
		return 0, true
	}
}
