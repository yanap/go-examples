package main

import "fmt"

func main() {
	// a()
	// b()
	// c()
	d()
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

func d() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	
	fmt.Println(elements["Un"])
	
	name, ok := elements["Un"]
	fmt.Println(name, ok)
}
