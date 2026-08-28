package animal

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	Name string
}

func (dog *Dog) Sound() {
	fmt.Println(dog.Name + " makes guau guau")
}

type Cat struct {
	Name string
}

func (cat *Cat) Sound() {
	fmt.Println(cat.Name + " makes miau miau")
}

func MakeSound(animal Animal) {
	animal.Sound()
}
