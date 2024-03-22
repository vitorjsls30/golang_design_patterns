// Factory Pattern
// Like the name states, the Factory Pattern allow us to create different objects that implement a shared interface, by
// simply telling a factory method what kind of concrete object is desired

// Why?
// With the Factory Pattern we can separate the concrete objects creation, whithin the business logic, from the actual client code

// How?
// We declare a interface that specifies what is common for all the concrete objects and delegates it's creation to a factory method

package main

import "fmt"

// The shared Interface
type ICar interface {
	getName() string
	setName(name string)
	getPlate() string
	setPlate(plate string)
}

// the Concrete Types

// we define a main base struct
type Car struct {
	Name  string
	Plate string
}

// Implementing the ICar interface...
func (c *Car) getName() string {
	return c.Name
}

func (c *Car) setName(name string) {
	c.Name = name
}

func (c *Car) getPlate() string {
	return c.Plate
}

func (c *Car) setPlate(plate string) {
	c.Plate = plate
}

// we indirectly implement ICar through the following concrete structs
type SedanCar struct {
	Car
}

type HatchbackCar struct {
	Car
}

// Concrete Factory
func newKiaK5() ICar {
	return &SedanCar{
		Car{
			Name:  "Kia K5",
			Plate: "ABCD",
		},
	}
}

func newHyundaiI30() ICar {
	return &HatchbackCar{
		Car{
			Name:  "Hyundai i30",
			Plate: "EFGH",
		},
	}
}

// The Factory Method...
func getCar(model string) ICar {
	var car ICar
	if model == "sedan" {
		car = newKiaK5()
	}

	if model == "hatchback" {
		car = newHyundaiI30()
	}

	return car
}

// Our main App...
func main() {
	println("[CarStore] Hi!! Let's build you a nice car!! Please, choose a type:")
	println("[CarStore]\n1 - Sedan\n2 - Hatchback")

	car1 := getCar("sedan")
	car2 := getCar("hatchback")

	println("[CarStore] the following types were created:")
	fmt.Printf("Name: %s - Plate: %s\n", car1.getName(), car1.getPlate())
	fmt.Printf("Name: %s - Plate: %s\n", car2.getName(), car2.getPlate())
}
