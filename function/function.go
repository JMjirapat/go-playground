package function

// function
// func <function_name>(<parameters>) <return_type> {}
func Add(a int, b int) int {
	return a + b
}

func AddSub(a, b int) (add int, sub int) {
	add = a + b
	sub = a - b
	return
}

func Add3(a, b, c int) int {
	return a + b + c
}

func AddArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func AddVariadic(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
