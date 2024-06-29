package referencetype

import (
	"fmt"
	"unicode/utf8"
)

func ReferenceType() {
	// Code for pointer
	// pointer
	// pointer is a variable that stores the memory address of another variable
	// var <variable_name> *<data_type>
	// &<variable_name> returns the memory address of the variable
	// *<variable_name> returns the value stored at the memory address
	p := 10
	pointer := &p
	fmt.Println(pointer)
	fmt.Println(*pointer)

	// array
	// array is a fixed size sequence of elements of the same type
	// var <variable_name> [<size>]<data_type>
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// slice
	// slice is a reference to an array
	// slice is a dynamic array
	slice := []int{1, 2, 3, 4, 5}
	slice2 := []int{}
	slice3 := make([]int, 5)
	slice4 := make([]int, 5, 10)
	slice4 = append(slice4, slice...)
	slice4 = append(slice4, 6, 7, 8, 9, 10)
	fmt.Println(slice, slice2, slice3, slice4)
	fmt.Println(slice)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	k := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(k)

	arrMod := append(slice, 6)
	arrMod2 := append(arr[:], 6)
	fmt.Println(arrMod, len(arrMod), cap(arrMod2))

	name := "เจมส์"
	println(name, len(name))
	println(utf8.RuneCountInString(name))

	// range
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// map
	// map[<key_type>]<value_type>
	m := map[string]int{
		"James": 32,
		"John":  28,
	}

	m["James"] = 33
	m["Jane"] = 30
	delete(m, "John")

	if age, ok := m["James"]; ok {
		fmt.Println(age)
	}

	fmt.Println(m["James"])

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	// Nil slices
	var x []int
	fmt.Println(s, len(s), cap(s))
	if x == nil {
		fmt.Println("nil!")
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
