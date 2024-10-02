package basic

//lint:ignore U1000 (example)
func sum_two_int(x, y int) int {
	return x + y
}

//lint:ignore U1000 (example)
func plus_minus_int(x int) (int, int) {
	if x < 0 {
		// passed by value
		// => no external modification
		x = -x
	}
	return x, -x
}

//lint:ignore U1000 (example)
func increase_int(x *int) {
	// passed by reference
	// => external modification
	*x = *x + 1
}

//lint:ignore U1000 (example)
func sum_int_array(a *[3]int) int {
	// avoid making copy of array by passing pointer
	sum := 0
	for i := range 3 {
		sum += (*a)[i]
	}
	return sum
}

//lint:ignore U1000 (example)
func sum_int_slice(a []int) int {
	// slice is a pointer => pass by reference
	sum := 0
	for x := range a {
		sum += x
	}
	return sum
}
