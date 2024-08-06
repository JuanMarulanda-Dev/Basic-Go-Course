package main

import "fmt"

type pc struct {
	ram   int
	disck int
	brand string
}

// Stringer = cuando se imprima el objeto de PC se va a imprimir de esta forma
func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disck, myPC.brand)
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

func main() {

	a := 50
	b := &a // Puntero de la variable

	fmt.Println(b)
	fmt.Println(*b) // Valor de lo que se encuentra en el puntero

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disck: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRam()

	fmt.Println(myPC)
	myPC.duplicateRam()

	fmt.Println(myPC)
}
