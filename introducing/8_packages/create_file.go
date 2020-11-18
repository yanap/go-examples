package main

import (
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.CLose()
	
	file.WriteString("test")
}
