package main

import "fmt"

func main() {
	// a()
	// b()
	// c()
	// d()
	f()
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

func e() {
	elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Flurrine",
		"Ne": "Neon",
	}
	fmt.Println(elements)
}

func f() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name": "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name": "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name": "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name": "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name": "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name": "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name": "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name": "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name": "Flurrine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name": "Neon",
			"state": "gas",
		},
	}

		if el, ok := elements["Li"]; ok {
			fmt.Println(el["name"], el["state"])
		}
}
