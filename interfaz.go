package main

import "fmt"

type animal interface {
	move() string
}

type animalType struct {
	name string
	action string
}
func (at animalType) move() string {
	return "I'm a " + at.name + " and i " + at.action
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	d := animalType{
		name: "dog",
		action: "walk",
	}
	moveAnimal(d)
	f := animalType{
		name: "fish",
		action: "swim",
	}
	moveAnimal(f)
	b := animalType{
		name: "bird",
		action: "fly",
	}
	moveAnimal(b)
}