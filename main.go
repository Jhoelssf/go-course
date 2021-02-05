package main

import "fmt"

func main() {
	var message string = "Hello World"
	easyMessage := "Hello World using :="
	fmt.Println(message)
	fmt.Println(easyMessage)

	// float numbers
	a := 10.
	var b float64 = 3
	fmt.Println(a / b)

	//integer numbers
	var c int = 10
	d := 3
	fmt.Println(c / d)

	// boolean
	var x bool = true
	y := false
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!x)

	privateFunction()

	PublicFunction()

	printHelloWorld()
}

func privateFunction() {
	fmt.Println("para usarlo en el paquete actual")
}

func PublicFunction() {
	fmt.Println("para poder usarlo en otros paquetes")
}

func printHelloWorld() {
	defer fmt.Println("World!")
	fmt.Println("Hello")
}
