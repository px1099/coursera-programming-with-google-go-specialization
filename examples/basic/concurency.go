package basic

import "fmt"
import "sync"

//lint:ignore U1000 (example)
func concurrent_task(msg string, wg *sync.WaitGroup) {
	defer wg.Done()  // Reduce counter by one
	fmt.Println(msg)
}

//lint:ignore U1000 (example)
func concurrent_waitgroup_example() {
	var wg sync.WaitGroup
	wg.Add(2)  // Increase counter by two
	go concurrent_task("A", &wg)
	go concurrent_task("B", &wg)
	wg.Wait()  // Wait until counter down to zero
	fmt.Println("All task completed")
}

//lint:ignore U1000 (example)
func prod(v1, v2 int, ch chan int) {
	// Channel communication is synchronous (send and receive at same time)
	// Unbuffered channel: sending will block until receive 
	// Buffered channel: sending will block if channel buffer is full
	ch <- v1 * v2
}

//lint:ignore U1000 (example)
func concurrent_channel_example() {
	ch := make(chan int)  // Unbuffered channel (cannot hold data)
	_ = make(chan int, 2)  // Buffered channel (Hold max 2 data)
	go prod(1, 2, ch)
	go prod(3, 4, ch)
	a := <- ch  // Unbuffered channel: blocks if no value is sent
	b := <- ch  // Buffered channel: blocks if buffer is empty
	fmt.Println(a*b)
}
