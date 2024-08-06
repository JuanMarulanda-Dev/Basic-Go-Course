package main

import "fmt"

func main() {

	// Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// continue y brake
	for i := 0; i < 10; i++ {

		if i == 2 {
			fmt.Println("Continue")
			continue
		}

		if i == 8 {
			fmt.Println("break")
			break
		}
	}
}
