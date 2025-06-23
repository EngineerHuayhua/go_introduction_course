package main

import (
	"fmt"
)

func main() {
	m1 := make(map[int]string) // Create an empty map with int keys and string values
	fmt.Printf("m1: %v, type: %T, size: %d bytes\n", m1, m1, len(m1))
	m1[1] = "A"
	m1[2] = "B"
	m1[3] = "C"
	fmt.Printf("m1 after adding elements: %v, size: %d bytes\n", m1, len(m1))

	delete(m1, 2) // Delete the key 2 from the map
	fmt.Printf("m1 after deleting key 2: %v, size: %d bytes\n", m1, len(m1))

	m1[1] = "A2" // Update the value for key 1
	fmt.Printf("m1 after updating key 1: %v, size: %d bytes\n", m1, len(m1))

	m1[7] = ""
	fmt.Println(m1[7])  // Accessing a key that does not exist returns the zero value for the value type (empty string in this case)
	fmt.Println(m1[99]) // Accessing a non-existent key returns the zero value, which is also an empty string

	v, ok := m1[7] // Check if key 7 exists in the map
	fmt.Println("Key 7 exists:", ok, "Value:", v)
	v, ok = m1[99] // Check if key 99 exists in the map
	fmt.Println("Key 99 exists:", ok, "Value:", v)

	m2 := map[int]string{ // Create and initialize a map with values
		1: "A",
		2: "B",
		3: "Z",
	}
	fmt.Printf("m2: %v, type: %T, size: %d bytes\n", m2, m2, len(m2))
}
