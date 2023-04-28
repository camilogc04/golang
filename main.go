package main

import (
	"RampUpFolder/devPath/cursogo/golang/variables"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	estado, nombre := variables.ConviertoATexto(1543)
	fmt.Println(estado)
	fmt.Println(nombre)
}
