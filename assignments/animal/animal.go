package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

var animals map[string]Animal = map[string]Animal{
	"cow": {"grass", "walk", "moo"},
	"bird": {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss "},
}

func processRequest(animal_name, info_requested string) {
	var animal Animal
	animal, exists := animals[animal_name]
	if !exists {
		fmt.Println("animal", animal_name, "not found in database")
		return
	}
	switch info_requested {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("info", info_requested, "is not a valid request")
	}
}

func main() {
	for {
		fmt.Print("> ")
		var animal_name, info_requested string
		fmt.Scan(&animal_name, &info_requested)
		processRequest(animal_name, info_requested)
	}
}
