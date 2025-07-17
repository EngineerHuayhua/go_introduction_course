package main

import (
	"fmt"

	"github.com/EngineerHuayhua/go_introduction_course/06-functions/function"
)

func main() {

	println("The result is:", function.Add(3, 4))

	function.RepeatString("Hello", 3)

	fmt.Println()
	v, err := function.Calc(function.SUM, 3, 6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Sum:", v)
	}

	v, err = function.Calc(function.DIV, 3, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Div:", v)
	}

	fmt.Println()
	x, y := function.Split(20)
	fmt.Println("value 1:", x, "value 2:", y)

	fmt.Println()
	v = function.MSum(23, 45, 67, 89, 12, 34, 56, 78, 90, 11, 22, 33, 44, 55, 66, 77, 88, 99)
	fmt.Println("multy sum:", v)

	v, err = function.MOperations(function.SUM, 2, 7, 1)
	fmt.Println("multy sum: ", v, "error: ", err)

	v, err = function.MOperations(function.SUB, 10, 5, 2, 1, 3)
	fmt.Println("multy sub: ", v, "error: ", err)

	v, err = function.MOperations(function.MUL, 10, 5, 2, 1, 3)
	fmt.Println("multy mul: ", v, "error: ", err)

	v, err = function.MOperations(function.DIV, 10, 5, 2, 1, 3)
	fmt.Println("multy div: ", v, "error: ", err)

	fmt.Println()
	fn := function.FactoryOperation(function.SUM)
	v = fn(10, 5)
	fmt.Println("Factory sum:", v)
	fn = function.FactoryOperation(function.SUB)
	v = fn(10, 5)
	fmt.Println("Factory sub:", v)
	fn = function.FactoryOperation(function.MUL)
	v = fn(10, 5)
	fmt.Println("Factory mul:", v)
	fn = function.FactoryOperation(function.DIV)
	v = fn(10, 5)
	fmt.Println("Factory div:", v)
}
