package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)

	c <- "mensaje 1"
	c <- "mensaje 2"

	fmt.Println(len(c), cap(c))

	// RANGE Y CLOSE
	// The `close(c)` function call is closing the channel `c`. This indicates that no more values will be
	// sent on the channel. It is important to close a channel when you are done sending values on it to
	// avoid deadlock situations. Once a channel is closed, you can still receive values from it until all
	// the values have been received.
	close(c)

	// The code snippet `for message := range c { fmt.Println(message) }` is iterating over the values
	// sent on the channel `c`
	for message := range c {
		fmt.Println(message)
	}

	// SELECT
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	// This code snippet is using a `for` loop to iterate twice (from `i=0` to `i=1`). Within each
	// iteration, it uses a `select` statement to receive messages from two different channels `email1`
	// and `email2`.
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
