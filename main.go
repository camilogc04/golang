package main

import (
	/* "RampUpFolder/devPath/cursogo/golang/variables"
	"fmt"
	"runtime" */
	"RampUpFolder/devPath/cursogo/golang/ejercicios"
	"fmt"
)

func main() {
	/* fmt.Println("Hello world!")
	estado, nombre := variables.ConviertoATexto(1543)
	fmt.Println(estado)
	fmt.Println(nombre)

	if os := runtime.GOOS; os == "windows" {
		fmt.Println("Estamos en Windows")
	} else if os == "darwin" {
		fmt.Println("Estamos en MAC OS X")
	} else {
		fmt.Println("Estamos en Linux")
	} */
	numero, mensaje := ejercicios.Ejercicio01("S50")

	fmt.Println("el numero ", numero, mensaje)

}
