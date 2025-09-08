package main

import "fmt"

func main() {
	b := 20
	c := 30

	for c != 0 {
		b, c = c, b%c
	}
	fmt.Print("le + grand diviseur commun est : ", b)

}
