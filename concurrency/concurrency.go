package concurrency

import (
	"fmt"
	"time"
)

func Concurrency() {
	// In go, concurrency
	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Task 2")
		}
	}()

	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Task 1")
	}

	//reference
	//https://bwoff.medium.com/the-comprehensive-guide-to-concurrency-in-golang-aaa99f8bccf6
	//https://github.com/luk4z7/go-concurrency-guide
}
