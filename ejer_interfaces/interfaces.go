package ejer_interfaces

import (
	"RampUpFolder/devPath/cursogo/golang/interfaces"
	"fmt"
)

func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s y estoy respirado", hu.Sexo())
}
