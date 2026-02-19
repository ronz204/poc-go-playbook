package demos

import "fmt"

func Operate(a, b int, callback func(int, int) int) int {
	fmt.Println("Performing operation on", a, "and", b)
	return callback(a, b)
}

func CallbacksDemo() {
	add := func(x, y int) int {
		return x + y
	}

	deduct := func(x, y int) int {
		return x - y
	}

	multiply := func(x, y int) int {
		return x * y
	}

	result := Operate(5, 3, add)
	println("Result of addition:", result)

	result = Operate(5, 3, deduct)
	println("Result of deduction:", result)

	result = Operate(5, 3, multiply)
	println("Result of multiplication:", result)
}
