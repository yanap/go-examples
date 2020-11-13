package main

import "fmt"

func main() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	fmt.Println(7)
	fmt.Println(8)
	fmt.Println(9)
	fmt.Println(10)
	
	fmt.Println(`1
	2
	3
	4
	5
	6
	7
	8
	9
	10
	`)
	// The for Statement
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
		if i % 2 == 0 {
			fmt.Println("divisible by 2")
		} else if i % 3 == 0 {
			fmt.Println("divisible by 3")
		} else if i % 4 == 0 {
			fmt.Println("divisible by 4")
		}
	}
	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else if i == 4 {
		fmt.Println("Four")
	} else if i == 5 {
		fmt.Println("Five")
	}
}
