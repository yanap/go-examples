package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1 + 1)
	// strings
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello," + "World")
	// booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)
	fmt.Println(!true)
	fmt.Println(!false)
}
