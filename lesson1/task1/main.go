package main

import "fmt"

func selfmultiply(x int) int {
	x *= x
	return x
}

func multireturn(x int) (string, int) {
	return "Hello", x
}

func addandsubstract(x, y int) (a, b int) {
	a = x + y
	b = x - y
	return
}

func main() {
	var x int
	x = 10
	fmt.Println(selfmultiply(x))
	fmt.Println(multireturn(x))
	fmt.Println(addandsubstract(x, 5))
}
