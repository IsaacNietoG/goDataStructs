package main

import "fmt"

func intMax(a int, b int) int {
	if a > b{
		return a
	}

	return b
}

func main() {
	fmt.Println("El maximo entre 15 y 20 es: ", intMax(15,20))
}
