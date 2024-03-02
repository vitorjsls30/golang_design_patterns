// Abstract Factory
// Let's you build families of related products without having to specify their concrete classes

// Why?
// Sometimes you need a way to specify "family of related objects" that may appear in a number of variations.
// Let's take a furniture store as an example: You sell chair, sofas and tables and for each of these family of items,
// you have 3 possible variations: Modern, Classical and Post-Modern.
// The abstract pattern allows you to tackle this kind of problemn.

// How?
// Basically, for each family of related items, you declare an interface. Ex: IChair, ISofa and ITable.
// Each concrete family of product must implement this base interface: ModernSofa implements ISofa, ClassicalChair
// implements IChair, and so on.
// We now, declare an abstract factory interface, that must be implemented by the concrete items factories.
// Finally, the client code specifies which kind of factory it's needed based on some kind of configuration
// and must only interact with family items based on their abstract interfaces, reducing this way any unnecessary
// hard coupling.

package main

import (
	"fmt"
)

// Our Product Families Interfaces
type IChair interface {
	hasLegs() bool
	sitDown()
	getName() string
	getYear() int16
	getNumberOfLegs() int16
}

type ISofa interface {
	hasLegs() bool
	sitDown()
	getName() string
	getYear() int16
	getNumberOfLegs() int16
}

// Abstract Products implementing the products base interfaces...
type chair struct {
	Name         string
	LaunchYear   int16
	NumberOfLegs int16
}

func (c *chair) getName() string {
	return c.Name
}

func (c *chair) getYear() int16 {
	return c.LaunchYear
}

func (c *chair) getNumberOfLegs() int16 {
	return c.NumberOfLegs
}

func (c *chair) hasLegs() bool {
	return true
}

func (c *chair) sitDown() {
	fmt.Println("chair::sitDown: sitting...")
}

type sofa struct {
	Name         string
	LaunchYear   int16
	NumberOfLegs int16
}

func (c *sofa) getName() string {
	return c.Name
}

func (c *sofa) getYear() int16 {
	return c.LaunchYear
}

func (c *sofa) getNumberOfLegs() int16 {
	return c.NumberOfLegs
}

func (c *sofa) hasLegs() bool {
	return true
}

func (c *sofa) sitDown() {
	fmt.Println("sofa::sitDown: sitting...")
}

// Our Family related Concrete Produts

// Victorian Products...
type VictorianChair struct {
	chair
}

type VictorianSofa struct {
	sofa
}

// Modern Products...
type ModernChair struct {
	chair
}

type ModernSofa struct {
	sofa
}

// Our abstract Factory
type IFurnitureFactory interface {
	makeChair() IChair
	makeSofa() ISofa
}

// Our concrete factories
type VictorianProductsFactory struct{}
type ModernProductsFactory struct{}

// implementing the abstract factory...
func (f *VictorianProductsFactory) makeChair() IChair {
	return &VictorianChair{
		chair{
			Name:         "Victorian Chair",
			LaunchYear:   1998,
			NumberOfLegs: 4,
		},
	}
}

func (f *VictorianProductsFactory) makeSofa() ISofa {
	return &VictorianSofa{
		sofa{
			Name:         "Victorian Sofa",
			LaunchYear:   1996,
			NumberOfLegs: 4,
		},
	}
}

func (f *ModernProductsFactory) makeChair() IChair {
	return &ModernChair{
		chair{
			Name:         "Modern Chair",
			LaunchYear:   2020,
			NumberOfLegs: 3,
		},
	}
}

func (f *ModernProductsFactory) makeSofa() ISofa {
	return &ModernSofa{
		sofa{
			Name:         "Modern Sofa",
			LaunchYear:   2024,
			NumberOfLegs: 3,
		},
	}
}

func main() {
	// the client code should be unaware of the type of chair or sofa returned
	// instead, it should only work with these products respective interfaces...
	furnitureFactory := func(furnitureType string) IFurnitureFactory {
		if furnitureType == "modern" {
			return &ModernProductsFactory{}
		}
		if furnitureType == "victorian" {
			return &VictorianProductsFactory{}
		}
		return nil
	}

	describeChair := func(product IChair) {
		fmt.Printf("describeChair:: Describing %s\n", product.getName())
		fmt.Printf("describeChair:: has Legs? %t\n", product.hasLegs())
		fmt.Printf("describeChair:: number of legs: %d\n", product.getNumberOfLegs())
		fmt.Println("describeChair:: let's sit down:")
		product.sitDown()
	}

	describeSofa := func(product ISofa) {
		fmt.Printf("describeSofa:: Describing %s\n", product.getName())
		fmt.Printf("describeSofa:: has Legs? %t\n", product.hasLegs())
		fmt.Printf("describeChair:: number of legs: %d\n", product.getNumberOfLegs())
		fmt.Println("describeSofa:: let's sit down:")
		product.sitDown()
	}

	// let's say that the chosen furniture type,in out App UI was initially 'Modern', so...
	factory := furnitureFactory("modern")

	// now, we request the desired victorian familty products
	producta := factory.makeSofa()
	productb := factory.makeChair()

	fmt.Println("Modern Products:")
	describeSofa(producta)
	describeChair(productb)

	//now, we change to modern family prpducts
	factory = furnitureFactory("victorian")

	productc := factory.makeSofa()
	productd := factory.makeChair()

	fmt.Println("Victorian Products:")
	describeSofa(productc)
	describeChair(productd)
}
