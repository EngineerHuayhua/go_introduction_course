package main

import (
	"fmt"
	"go_introduction_course/06-functions/function"
)

func main() {

	println("The result is:", function.Add(3, 4))

	function.RepeatString("Hello", 3)

	fmt.Println()
	v, err := function.Calc(function.SUM, 3, 6)
	fmt.Println("Sum:", v, "- Error:", err)

	fmt.Println()
	v, err = function.Calc(function.DIV, 3, 0)
	fmt.Println("Div:", v, "- Error:", err)
}
