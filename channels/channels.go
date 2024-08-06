package main

import "fmt"

// una buena practica es expecificar si el channel va a ser solo de entrada o solo de salida
// se hace de la siguiente forma
// c chan<- string = el channel solo va a permitir ingresar valores tipo string.
// c <-chan string = el chanel solo va a permitir la lectura del contenido.
func say(text string, c chan string) {
	c <- text // Así se el pasa info al channel
}

func message(text string, c chan<- string) {
	c <- text
}

func main() {

	//chan = se usa para especificar que se va a crear un channl (kernel)
	//type = se especifica el tipo de datos que va a manejar el channel
	//cantidad= cantidad de go routines que se va a manejar por este channel
	c := make(chan string)

	fmt.Println("Hello")

	go say("Bey", c)

	fmt.Println(<-c) // Sacar información del channel

	// Manejo de Range, Close y Select
	c0 := make(chan string, 2)
	c0 <- "Mensaje 1"
	c0 <- "Mensaej 2"

	// El len() en los channeles me muestra cuandos elementros hay dentro de ese channel y cap() me dice la capacidad de ese channel
	fmt.Println(len(c0), cap(c0))

	c1 := make(chan string, 2)
	c2 := make(chan string, 2)

	message("Hola", c1)
	message("Mundo", c2)

	// Range y Close
	// De esta forma le decimos al runtime de GO que cierre el channel, eso quiere que no va a recibir ningun dato adicional
	// a pesar que el channel tengo espacio para almacenar más datos
	// Es recomendable que los channels se cierren una vez no se necesiten utilizar

	close(c1)

	// Range
	// Basicamente es la forma de iterar cada uno de los elementos que se encuentran en un channel.
	for message := range c {
		fmt.Println(message)
	}

	// Select
	// Esta es la forma de ver en un grupo de channels cual de lso channels responde primero
	// Para utilizar select se debe tener en cuenta cual va a ser la cantidad de datos que se va a manejar por channel

	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje 1", email1)
	go message("mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email1:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
