package flowcontrol

import (
	"fmt"
	"time"
)

func FlowControl() {
	x := 10

	// if else
	if x == 10 {
		fmt.Println("x is 10")
	} else {
		fmt.Println("x is not 10")
	}

	// if statement with a short statement
	if y := 10; y == 10 {
		fmt.Println("y is 10")
	}

	// switch case
	switch x {
	case 10:
		fmt.Println("x is 10")
	case 20:
		fmt.Println("x is 20")
	default:
		fmt.Println("x is not 10 or 20")
	}

	// switch case with evalutaion order
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }
}
