package main

import "fmt"

func main() {
	fmt.Println("For Conditional")
	// For conditional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("For while")
	// For While
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("For forever")
	//For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
	}

}
