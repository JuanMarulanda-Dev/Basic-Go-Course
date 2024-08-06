package main

import "fmt"

func main() {
	// Make se utiliza para crear diferentes tipos de variables
	//map[tipo_llave]tipo_valor
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer un mapa
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value := m["Jose"]
	fmt.Println(value)

	// Si no existe el valor (esta segunda variable nos dice si existe o no el valor por medio de un bool)
	value1, ok := m["Josep"]
	fmt.Println(value1, ok)

	// Nota:
	/*
		Al recorrer un map con range hay que tener en cuenta que no siempre vamos a obtener los valores
		en el mismo orden que fueron ingresados, como la insercion ocurre de forma concurrente el orden se puede ver de forma aleatorea
		Si necesitas tener los valores en un orden especifico es mejor utilizar los slices
	*/
}
