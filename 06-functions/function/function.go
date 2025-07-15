package function

import (
	"errors"
	"fmt"
)

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

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y mustn't be zero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	default:
		return 0, errors.New("has been an error")
	}
}

func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return
}

func MSum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("there aren't values")
	}

	result := values[0]

	for _, v := range values[1:] {
		var err error
		result, err = Calc(op, result, v)
		if err != nil {
			return 0, err
		}
	}
	return result, nil

}
