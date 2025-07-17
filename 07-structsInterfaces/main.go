package main

import (
	"encoding/json"
	"fmt"

	"github.com/EngineerHuayhua/go_introduction_course/07-structsInterfaces/structsInterface/structs"
	"github.com/EngineerHuayhua/go_introduction_course/07-structsInterfaces/structsInterface/vehicles"
)

func main() {
	// This is the main function where the program starts.
	// You can add your code here to use the structs defined in the package.
	var p1 structs.Product
	p1.ID = 1
	p1.Name = "Laptop"
	p1.Type = structs.Type{}
	p1.Price = 999.99
	fmt.Println(p1)

	p2 := structs.Product{}
	p2.ID = 2
	p2.Name = "Smartphone"
	p2.Type = structs.Type{}
	p2.Price = 499.99
	fmt.Println(p2)

	p3 := structs.Product{
		ID:    3,
		Name:  "Tablet",
		Type:  structs.Type{},
		Price: 299.99,
	}
	fmt.Println(p3)

	p4 := structs.Product{
		ID:   4,
		Name: "Smartwatch",
		Type: structs.Type{
			ID:          1,
			Code:        "SW001",
			Description: "Smartwatch with health tracking features",
		},
		Count: 10,
		Price: 199.99,
		Tags:  []string{"wearable", "sarasa", "freezer", "refrigerator"},
	}
	fmt.Println(p4)

	// Convert p4 to JSON and print it
	v, err := json.Marshal(p4)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println()
	fmt.Println(string(v))

	fmt.Println("Total Price of p4:", p4.TotalPrice())
	fmt.Println()

	fmt.Println(p4)
	p4.SetName("Updated Smartwatch")
	p4.AddTags("newtag1", "newtag2", "newtag3")
	fmt.Println("After updating name:", p4.Name)
	fmt.Println(p4)

	p5 := structs.Product{
		ID:   5,
		Name: "Cortina",
		Type: structs.Type{
			ID:          2,
			Code:        "HW001",
			Description: "Hogar",
		},
		Count: 3,
		Price: 2699.99,
		Tags:  []string{"Hogar", "Cortina"},
	}

	p6 := structs.Product{
		ID:   6,
		Name: "Naranja",
		Type: structs.Type{
			ID:          3,
			Code:        "BG001",
			Description: "Alimento",
		},
		Count: 20,
		Price: 19.99,
		Tags:  []string{"alimento", "verdura"},
	}

	car := structs.NewCar(11212)
	car.AddProduct(p4, p5, p6)

	fmt.Println("PRODUCTS CAR")
	fmt.Println("Total Products:", len(car.Products))
	fmt.Printf("Total Price of Car: $%.2f\n", car.Total())

	fmt.Println()
	fmt.Println("VEHICLES")

	// Create a new Car instance and calculate the distance based on time (dentro de paquete vehicles llamamos a la estructura Car).
	// Here, we assume the car travels at a speed of 100 km/h.
	carV := vehicles.Car{
		Time: 160, // Time in minutes
	}
	fmt.Printf("Distance traveled by car: %.2f km\n", carV.Distance())

	vArray := []string{"CAR", "BIKE", "MOTORCYCLE", "BUS", "TRUCK"}

	var d float64
	for _, v := range vArray {
		fmt.Printf("Vehicle: %s\n", v)
		veh, err := vehicles.New(v, 400)
		if err != nil {
			fmt.Println("Error creating vehicle:", err)
			continue
		}
		distance := veh.Distance()
		fmt.Printf("Distance traveled by %s: %.2f km\n", v, distance)
		fmt.Println()
		d += distance
	}
	fmt.Printf("Total Distance Traveled by all vehicles: %.2f km\n", d)
	fmt.Println()
}
