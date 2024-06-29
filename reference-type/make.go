package referencetype

import "fmt"

func MakeForReference() {
	//slice with make
	x1 := make([]int, 5)
	x2 := make([]int, 5, 10)
	x3 := []int{0, 0, 0, 0, 0}
	x4 := [...]int{0, 0, 0, 0, 0}
	var x5 []int
	fmt.Println(x5 == nil)
	fmt.Println(x1, x2, x3, x4)
	fmt.Println(len(x1), cap(x1), len(x2), cap(x2), len(x3), cap(x3), len(x4), cap(x4))

	x1 = append(x1, 1, 2, 3, 4, 5)
	x2 = append(x2, 1, 2, 3, 4, 5)
	x3 = append(x3, 1, 2, 3, 4, 5)

	fmt.Println(len(x1), cap(x1), len(x2), cap(x2), len(x3), cap(x3), len(x4), cap(x4))
	fmt.Println(x1, x2, x3, x4)

	x1 = append(x1, 6, 7)
	x2 = append(x2, 6, 7)
	x3 = append(x3, 6, 7)

	fmt.Println(len(x1), cap(x1), len(x2), cap(x2), len(x3), cap(x3), len(x4), cap(x4))
	fmt.Println(x1, x2, x3, x4)

	//map with make
	y1 := make(map[string]int)
	y2 := map[string]int{}
	var y3 map[string]int

	fmt.Println(y1 == nil, y2 == nil, y3 == nil)

	y1["one"] = 1
	y2["one"] = 1
	// y3["one"] = 1 // panic: assignment to entry in nil map

	fmt.Println(y1, y2, y3)

	//channel with make
	z1 := make(chan int)
	z2 := make(chan int, 5)
	// var z3 chan int

	fmt.Println(z1 == nil, z2 == nil)

	go func() {
		z1 <- 1
	}()
	z2 <- 1
	fmt.Println(<-z1, <-z2)
	// z3 <- 1 // fatal error: all goroutines are asleep - deadlock!

}
