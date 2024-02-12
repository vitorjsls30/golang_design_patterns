// Factory Pattern
// Like the name states, the Factory Pattern allow us to create different objects that implement a shared interface, by
// simply telling a factory method what kind of concrete object is desired

// Why?
// With the Factory Pattern we can separate the concrete objects creation, whithin the business logic, from the actual client code

// How?
// We declare a interface that specifies what is common for all the concrete objects and delegates it's creation to a factory method

package main

// The shared Interface
type ICar interface {
	getName() string
}

// the Concrete Types
type SedanCar struct {
	Name string
}

type HatchbackCar struct {
	Name string
}

// Implementing the ICar interface...
func (s *SedanCar) getName() string {
	return s.Name
}

func (h *HatchbackCar) getName() string {
	return h.Name
}

// The Factory Method...
func getCar(model string) ICar {
	var car ICar
	if model == "sedan" {
		car = &SedanCar{Name: "Kia K5"}
	}

	if model == "hatchback" {
		car = &HatchbackCar{Name: "Hyundai i30"}
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
	println(car1.getName())
	println(car2.getName())
}
