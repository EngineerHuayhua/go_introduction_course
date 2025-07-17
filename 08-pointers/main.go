package main

import "fmt"

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("address: %p\n", &m)
	m.Name = name
}
func (m *MyStruct) SetPtr(name string) {
	fmt.Printf("address Ptr: %p\n", m)
	m.Name = name
}

func main() {

	// i is a variable of type int
	var i int
	// iPtr is a pointer to an int
	// It holds the address of the variable i
	var iPtr *int

	i = 34
	iPtr = &i

	// Print the address of i
	// This will print the memory address where i is stored
	fmt.Println(&i)
	// Print the address stored in iPtr
	// This will print the same memory address as above
	fmt.Println(iPtr)
	// Print the value of i
	// This will print the value of i
	fmt.Println(i)
	// Print the value pointed to by iPtr
	// This will dereference iPtr and print the value of i
	fmt.Println(*iPtr)

	*iPtr = 1
	fmt.Printf("value of i: %d, value pointed to by iPtr: %d\n", i, *iPtr)
	fmt.Println()

	myVar := 30
	fmt.Printf("address of myVar: %p, value of myVar: %d\n", &myVar, myVar)
	increment(myVar)
	fmt.Println("myVar after increment:", myVar)
	// el valor se pasa por referencia, por lo que no se modifica el valor original
	incrementPtr(&myVar)
	fmt.Println("myVar after incrementPtr:", myVar)
	fmt.Println()

	// los slice trabajan por referencia
	// Si se pasa un slice a una funcion, se pasa la referencia al slice y la funcion puede modificar el slice original
	var mySlice []int
	mySlice = append(mySlice, 1, 2, 3)
	fmt.Printf("address of mySlice: %p, mySlice: %v\n", mySlice, mySlice)
	fmt.Printf("address 1: %p, address 2: %p, address 3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Println("mySlice after updateSlice:", mySlice)
	fmt.Println()

	myStruct := &MyStruct{
		ID:   1,
		Name: "John",
	}
	fmt.Printf("address of myStruct: %p\n", myStruct)
	fmt.Printf("id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	updateStruct(myStruct)
	fmt.Printf("after updateStruct - id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	fmt.Println()

	fmt.Printf("address of myStruct before Set: %p\n", myStruct)
	myStruct.Set("Julia")
	fmt.Printf("after Set - id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	fmt.Printf("address of myStruct after Set: %p\n", myStruct)
	myStruct.SetPtr("Doe")
	fmt.Printf("after SetPtr - id: %d, name: %s\n", myStruct.ID, myStruct.Name)
	fmt.Printf("address of myStruct after SetPtr: %p\n", myStruct)
}

func increment(val int) {
	fmt.Println(&val)
	val++
}

func incrementPtr(val *int) {
	fmt.Println(val)
	*val++
}

func updateSlice(v []int) {
	fmt.Printf("address of v: %p\n", v)
	fmt.Printf("address 1: %p, address 2: %p, address 3: %p\n", &v[0], &v[1], &v[2])
	v[0] = 12

	// la adicion de mas elementos al slice
	// no modifica la direccion de memoria del slice original
	v = append(v, 4, 5)
}

func updateStruct(v *MyStruct) {
	fmt.Printf("address of v in the function: %p\n", v)
	v.ID = 999
	v.Name = "Jane"

}
