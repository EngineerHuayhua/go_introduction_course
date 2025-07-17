package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int
	myIntVar = -12
	fmt.Println("mi variables es: ", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println("mi variables Uint es: ", myUintVar)

	var myStringVar string
	myStringVar = "Hola, soy una variable de tipo string"
	fmt.Println("mi variables String es: ", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("mi variables Bool es: ", myBoolVar)

	// Para mostrar la dirección de memoria de una variable, usamos el operador &
	fmt.Println("Dirección de memoria de myIntVar: ", &myIntVar)

	myIntVar2 := 10
	fmt.Println("mi variables es: ", myIntVar2)

	// Constantes
	const myConstantVar = "he464"
	fmt.Println("mi constante es: ", myConstantVar)

	const myConstantIntVar int = 10
	fmt.Println("mi constante int es: ", myConstantIntVar)
	//myConstantIntVar = 20                                  // Esto no es válido, ya que las constantes no pueden ser reasignadas
	fmt.Println("mi constante int es: ", myConstantIntVar) // Esto no se ejecutará debido al error anterior

	const myStringConst string = "Hola, soy una constante de tipo string"
	fmt.Println("mi constante String es: ", myStringConst)

	fmt.Println()
	var my8BitIntVar int8
	fmt.Printf("Type: %T\t, bytes: %d\t, bits: %d\n", my8BitIntVar, unsafe.Sizeof(my8BitIntVar), unsafe.Sizeof(my8BitIntVar)*8)

	fmt.Println()
	var my16BitIntVar int16
	fmt.Printf("Type: %T\t, bytes: %d\t, bits: %d\n", my16BitIntVar, unsafe.Sizeof(my16BitIntVar), unsafe.Sizeof(my16BitIntVar)*8)

	fmt.Println()
	var my32BitIntVar int32
	fmt.Printf("Type: %T\t, bytes: %d\t, bits: %d\n", my32BitIntVar, unsafe.Sizeof(my32BitIntVar), unsafe.Sizeof(my32BitIntVar)*8)

	fmt.Println()
	var my64BitIntVar int64
	fmt.Printf("Type: %T\t, bytes: %d\t, bits: %d\n", my64BitIntVar, unsafe.Sizeof(my64BitIntVar), unsafe.Sizeof(my64BitIntVar)*8)

	fmt.Println()
	var my8BitsUnitVar uint8
	fmt.Printf("Type: %T\t, bytes: %d\t, bits: %d\n", my8BitsUnitVar, unsafe.Sizeof(my8BitsUnitVar), unsafe.Sizeof(my8BitsUnitVar)*8)

	fmt.Println()
	var my16BitsUnitVar uint16
	fmt.Printf("Type: %T\t, bytes: %d\t, bits: %d\n", my16BitsUnitVar, unsafe.Sizeof(my16BitsUnitVar), unsafe.Sizeof(my16BitsUnitVar)*8)

	fmt.Println()
	var my32BitsUnitVar uint32
	fmt.Printf("Type: %T\t, bytes: %d\t, bits: %d\n", my32BitsUnitVar, unsafe.Sizeof(my32BitsUnitVar), unsafe.Sizeof(my32BitsUnitVar)*8)

	fmt.Println()
	var my64BitsUnitVar uint64
	fmt.Printf("Type: %T\t, bytes: %d\t, bits: %d\n", my64BitsUnitVar, unsafe.Sizeof(my64BitsUnitVar), unsafe.Sizeof(my64BitsUnitVar)*8)

	fmt.Println()
	var myFloat32Var float32
	fmt.Printf("Float default value: %f\n", myFloat32Var)
	fmt.Printf("Type: %T\t, bytes: %d\t, bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64
	fmt.Printf("Type: %T\t, bytes: %d\t, bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	fmt.Println()
	var myStringVar2 string
	fmt.Printf("String default value: %s\n", myStringVar2)

	myStringVar3 := `Hola, soy una variable de tipo string
	con comillas invertidas, lo que me permite`
	fmt.Printf("String with backticks (alt gr + }): %s\n", myStringVar3)

	fmt.Println("conversion de variables")
	{
		floatVar := 33.14
		fmt.Printf("Type: %T\t, value: %f\n", floatVar, floatVar)
		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("Type: %T\t, value: %s\n", floatStrVar, floatStrVar)

		intVar := 22
		fmt.Printf("Type: %T\t, value: %d\n", intVar, intVar)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("Type: %T\t, value: %s\n", intStrVar, intStrVar)

		intVal, err := strconv.ParseInt("1353", 10, 64)
		fmt.Println(err)
		fmt.Printf("Type: %T\t, value: %d\n", intVal, intVal)

		intVal1, err := strconv.ParseInt("1353ge", 10, 64)
		fmt.Println(err)
		fmt.Printf("Type: %T\t, value: %d\n", intVal1, intVal1)

		fmt.Println("conversion de string a float")
		floatVal1, err := strconv.ParseFloat("-11.96", 64)
		fmt.Println(err)
		fmt.Printf("Type: %T\t, value: %f\n", floatVal1, floatVal1)
	}

}
