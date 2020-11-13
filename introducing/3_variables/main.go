package main

import "fmt"

var (
	a = 5
	b = 10
	c = 15
)

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
	
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
	
	x = "hello"
	y = "world"
	fmt.Println(x == y)

	x = "hello"
	y = "hello"
	fmt.Println(x == y)
}
