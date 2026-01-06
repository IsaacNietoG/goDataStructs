package main
import "fmt"

func strMax(a string, b string) (string) {
	if a[0] > b[0] {
		return a
	}
	if a[0] < b[0] {
		return b
	}
	return strMax(a[1:], b[1:])
}

func main() {
	fmt.Println("La mayor entre las cadenas Alberto y Enrique es: ", strMax("Alberto", "Enrique"))
}
