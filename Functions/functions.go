package main

import "fmt"

// Difine arguments
func normalFunction(message string) {
	fmt.Println(message)
}

// function with several arguments
func tripleArguments(a int, b int, c string) {
	fmt.Printf("Priemr valor %d, segundo valor %d, tercer valor %s\n", a, b, c)
}

// Define type of return
func returnValue(a int) int {
	return a * 2
}

// Return double values
func doubleReturn(a int) (c, b int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo")
	tripleArguments(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("value: ", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("Values: ", value1, value2)

	// Doesn't take in count all of the returns of a function
	value3, _ := doubleReturn(3)
	fmt.Println("Value3: ", value3)
}
