/**
 * Author: Rakshita Mathur
 * StudentID: 300215340
 * Course: CSI 2120 Programming Paradigms
 * Assignment 1 Question 3
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// This function generates a stream of random integers between 0 and 999,999,
// multiplies each integer by an integer 'm' and sends it to a channel.
// The function continues generating and sending integers until it receives
// a signal on the 'stop' channel.
func RandomGenerator2(wg *sync.WaitGroup, stop <-chan bool, m int) <-chan int {
	intStream := make(chan int)
	go func() {
		// Decrement the wait group counter when the function completes
		defer func() { wg.Done() }()
		// Close the channel when the function completes
		defer close(intStream)
		for {
			select {
			// If the 'stop' channel receives a signal, stop generating numbers
			case <-stop:
				return
			// Generate a random integer between 0 and 999,999, multiply it by 'm',
			// and send it to the channel
			case intStream <- rand.Intn(1000000) * m:
			}
		}
	}()
	return intStream
}

// This function checks if an integer 'x' is a multiple of an integer 'm',
// and returns true if it is, false otherwise.
func Multiple(x int, m int) bool {
	if x%m == 0 {
		return true
	}
	return false
}

func main() {
	// Create a wait group to synchronize the execution of the goroutines
	var wg sync.WaitGroup
	// Create a channel to send a stop signal to the generator goroutines
	stop := make(chan bool)
	// Add 3 goroutines to the wait group
	wg.Add(3)
	// Create 3 channels to generate and receive random integers
	m5 := RandomGenerator2(&wg, stop, 5)
	m13 := RandomGenerator2(&wg, stop, 13)
	m97 := RandomGenerator2(&wg, stop, 97)
	// Listen to the channels and print out the integers that are multiples of 5, 13, and 97
	for i := 0; i < 100; i++ {
		select {
		// If an integer is received on the 'm5' channel, check if it is a multiple of 5,
		// and print it out if it is
		case x := <-m5:
			if Multiple(x, 5) {
				fmt.Println("Multiple of 5:", x)
			}
		// If an integer is received on the 'm13' channel, check if it is a multiple of 13,
		// and print it out if it is
		case x := <-m13:
			if Multiple(x, 13) {
				fmt.Println("Multiple of 13:", x)
			}
		// If an integer is received on the 'm97' channel, check if it is a multiple of 97,
		// and print it out if it is
		case x := <-m97:
			if Multiple(x, 97) {
				fmt.Println("Multiple of 97:", x)
			}
		}
	}
	// Send a signal on the 'stop' channel to stop the generator goroutines
	close(stop)
	// Wait for all generator goroutines to complete before exiting the program
	wg.Wait()
}
