package main

import "fmt"

func main() {
	yearsOld := 32

	fmt.Println("Operators")
	fmt.Println(yearsOld > 30)  //false
	fmt.Println(yearsOld >= 30) //true
	fmt.Println(yearsOld < 30)  //false
	fmt.Println(yearsOld <= 30) //false
	fmt.Println(yearsOld >= 30) //true
	fmt.Println(yearsOld == 32) //true

	fmt.Println()
	fmt.Println("Logical Operators OR")
	fmt.Println(yearsOld > 30 || yearsOld < 20) //true OR false = true
	fmt.Println(yearsOld < 30 || yearsOld > 20) //true OR true = true
	fmt.Println(yearsOld < 30 || yearsOld < 20) //false OR false = false

	fmt.Println()
	fmt.Println("Logical Operators AND")
	fmt.Println(yearsOld > 30 && yearsOld < 20) //true AND false = false
	fmt.Println(yearsOld < 30 && yearsOld > 20) //true AND true = true
	fmt.Println(yearsOld < 30 && yearsOld < 20) //false AND false = false

	fmt.Println()
	fmt.Println("Logical Operators negacion !")
	fmt.Println(yearsOld > 30)    //true
	fmt.Println(!(yearsOld > 30)) //false

	fmt.Println()
	fmt.Println("Multiples operadores")
	fmt.Println(yearsOld > 30 || yearsOld < 20 && yearsOld == 32)   //true OR false AND true = true
	fmt.Println(yearsOld > 40 && yearsOld < 20 || yearsOld == 28)   //false AND false OR false = false
	fmt.Println(yearsOld > 40 && (yearsOld < 20 || yearsOld == 28)) //false AND false OR false = false

	fmt.Println()
	yearsOld = 28
	fmt.Println("Condicionales")
	if yearsOld < 30 {
		fmt.Println("You are not older than 30 years old")
	}

	fmt.Println()
	fmt.Println("If")
	boolVal := true
	if !boolVal {
		fmt.Println("The value is true")
	} else {
		fmt.Println("The value is false")
	}

	if value := true; value {
		fmt.Println("The value is true")
	} else {
		fmt.Println("The value is false")
	}

	fmt.Println()
	fmt.Println("If else")
	if value := 2; value == 1 {
		fmt.Println("The value is one")
	} else if value == 2 {
		fmt.Println("The value is two")
	} else {
		fmt.Println("The value is not one or two")
	}

	fmt.Println()
	fmt.Println("Switch")
	yearsOld = 28
	switch yearsOld {
	case 18:
		fmt.Println("You are 18 years old")
	case 28:
		fmt.Println("You are 28 years old")
	default:
		fmt.Println("You are not 18 or 28 years old")
	}

	switch {
	case yearsOld < 18:
		fmt.Println("You are not 18 or 28 years old")
	case yearsOld == 18:
		fmt.Println("You are 18 years old")
	case yearsOld == 28:
		fmt.Println("You are 28 years old")
	default:
		fmt.Println("You are not 18 or 28 years old")
	}
}
