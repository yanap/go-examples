package main

import "fmt"

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func max(xs ...int) int {
	var max int
	for i, x := range xs {
		if i == 0 || x < max {
			max = x
		}
	}
	return max
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func f() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("RECOVER:", x)
		}
	}()
	panic("panic!")
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	fmt.Println(half(10))
	fmt.Println(max(1, 2, 3, 4, 5))
	fmt.Println(fibonacci(3))
	makeOddGenerator()
	x := 10
	xx := &x
	y := 20
	yy := &y
	swap(xx, yy)
	f()
}
