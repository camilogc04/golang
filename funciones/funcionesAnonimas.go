package funciones

import (
	"fmt"
)

func Calculos() {

	//FUNCIONES ANONIMAS EN GO DONDE A UNA VARIABLE LE ASIGNO UNA FUNCION MUY SNECILLA EN ESTE CASO
	//QUE ES SUMA DE DOS VALORES (SU OBJETIVO ES ACOTAR FUNCIONES)
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}
	//AQUI VEMOS COMO SUMA ES UNA VARIABLE PERO ES RECONOCIDA COMO UNA FUNCION TAMBIEN
	fmt.Println(calculo(10, 25))

	//SOBREESCRIBIMOS LA FUNCION CALCULOS
	calculo = func(numero1 int, numero2 int) int {
		return numero1 / numero2
	}
	//AQUI VEMOS COMO SUMA ES UNA VARIABLE PERO ES RECONOCIDA COMO UNA FUNCION TAMBIEN
	fmt.Println(calculo(10, 25))
}
