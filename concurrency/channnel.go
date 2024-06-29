package concurrency

import (
	"fmt"
	"sync"
)

func Channel() {
	// channel
	// ch := make(chan <type>)
	// ch := make(chan <type>, <buffer_size>)
	// ch <- <value> sends a value to the channel
	// <-ch receives a value from the channel
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	value := <-ch
	fmt.Println(value)

	// select
	// select is used to wait on multiple channel operations
	// select {
	// case <variable_name> = <-<channel_name>:
	// 	// do something
	// case <channel_name> <- <value>:
	// 	// do something
	// }
	select {
	case value := <-ch:
		fmt.Println(value)

	case ch <- 20:
		fmt.Println("Sent value to channel")

	default:
		fmt.Println("Default case")
	}

	// close
	// close is used to close a channel
	close(ch)

	// range over channel
	// for <variable_name> := range <channel_name> {
	// 	// do something
	// }
	for value := range ch {
		fmt.Println(value)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2")
	}()

	wg.Wait()

	// buffered channel
	// ch := make(chan <type>, <buffer_size>)
	// buffer_size is the number of values that can be stored in the channel
	// if buffer_size is not provided, the channel is unbuffered
	// an unbuffered channel has a buffer_size of 0
	// sending a value to an unbuffered channel blocks until another goroutine receives the value
	// sending a value to a buffered channel blocks only when the buffer is full
	// receiving a value from a buffered channel blocks only when the buffer is empty
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	// sync.Mutex
	// sync.Mutex is used to provide a locking mechanism to ensure that only one goroutine is running the critical section of code
	// m := sync.Mutex{}
	// m.Lock() locks the mutex
	// m.Unlock() unlocks the mutex
	// m.RLock() locks the mutex for reading
	// m.RUnlock() unlocks the mutex after reading
	m := sync.Mutex{}
	m.Lock()
	fmt.Println("Locked")
	m.Unlock()
	fmt.Println("Unlocked")
}
