package main

import "fmt"

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	
	someOtherName := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(someOtherName))

	f()

	fmt.Println(f1())

	x, y := mul()
	fmt.Println(x, y)

	fmt.Println(add(1,2,3))

	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1,1))
}

func average(xs []float64) float64 {
	//panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

var x int = 5
func f() {
	fmt.Println(x)
}

func f1() int {
	return f2()
}

func f2() (r int) {
	r = 1
	return
}

func mul() (int, int){
	return 5, 6
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
