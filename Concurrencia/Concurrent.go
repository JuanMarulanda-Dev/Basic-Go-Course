package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() // con esto le decimos al waitGroup que el Goroutine ya pudo ser terminada.
	fmt.Println(text)
}

// Una Goroutine es una forma de manejar múltiples procesos de forma concurrente.
// la forma de crear una go routine es solo agregarle la palabra GO

//WaitGroup = espera a qeu una colección de gorotines termine su ejecución, es más complejo mantenerlo

// la funcion main no es más que una goroutine
func main() {
	var wg sync.WaitGroup

	fmt.Println("Esto es un")

	wg.Add(1) // LE decimos al waitGroup que tenemos un goroutine en ejecucion para que la espere cuando ella termine
	go say("Hello", &wg)

	wg.Add(1)
	go say("Wolrd", &wg) // Esta linea se ejecuta de forma concurrente // forma asincrona

	wg.Wait() // No va a terminar el goroutine hasta que todo el waitGroup este terminado
}
