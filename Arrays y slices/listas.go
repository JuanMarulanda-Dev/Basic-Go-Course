package main

import "fmt"

func main() {
	// Arrays
	var array [4]int
	array[0] = 1
	array[1] = 1

	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Slice Methods [incluyente: excluyente]
	// inclutenye quiere decir que si va a incluir ese indice,
	// excluyente quiere decir que no va agregar ese indice si no el aterior
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Agregar un valor nuevo
	slice = append(slice, 7)
	fmt.Println(slice)

	// Agregar una nuevo slice a un slice ya existente
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)

}
