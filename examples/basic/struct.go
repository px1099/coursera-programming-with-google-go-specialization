package basic

type Person struct {
	name	string
	address string
	phone	string
}

//lint:ignore U1000 (example)
func struct_example() {
	p1 := new(Person)
	p2 := Person{name: "joe", address: "SG", phone: "123"}
	p2.name = "max"
	_, _ = p1, p2
	// can also be declared for 1 time usage
	dog := struct {
		name	string
		isGood	bool
	}{"Rex", true}
	_ = dog
}
