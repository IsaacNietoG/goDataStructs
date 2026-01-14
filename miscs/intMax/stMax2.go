package main


import "fmt"

func strMax(a string, b string) (string) {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("La mayor entre las cadenas Alberto y Enrique es: ", strMax("Alberto", "Enrique"))
}
