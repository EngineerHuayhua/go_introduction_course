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

}
