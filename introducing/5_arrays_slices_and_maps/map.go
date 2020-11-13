package main

import "fmt"

func main() {
	// a()
	// b()
	c()
}

func a() {
	var x map[string]int
	x["key"] = 10
	fmt.Println(x)
}

func b() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
}

func c() {
	x := make(map[int]int)
	x[1] = 10
	fmt.Println(x[1])
	delete(x, 1)
	fmt.Println(x[1])
}
