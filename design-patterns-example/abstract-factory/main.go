package main

import "fmt"

// like factory but you can create families
// animal is the type for abstract factory
type Animal interface {
	Says()
	LikesWater() bool
}

// create a concrete factory for dogs

// implement the abstract factory for dogs
type Dog struct {
}

func (d *Dog) Says() {
	fmt.Println("Woof")
}
func (d *Dog) LikesWater() bool {
	return true
}

// create a concrete factory for cats
type Cat struct {
}

func (d *Cat) Says() {
	fmt.Println("Meow")
}
func (d *Cat) LikesWater() bool {
	return false
}

// implement the abstract factory for cats

type AnimalFactory interface {
	New() Animal
}
type DogFactory struct{}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	//	create one each of a dogfactory and a cat factory
	DogFactory := DogFactory{}
	CatFactory := CatFactory{}
	dog := DogFactory.New()
	cat := CatFactory.New()
	dog.Says()
	fmt.Println("a dog likes water ? ", dog.LikesWater())
	cat.Says()
	cat.LikesWater()
	fmt.Println("a cat likes water ? ", cat.LikesWater())

}
