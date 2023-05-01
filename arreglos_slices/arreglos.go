package arreglos_slices

import "fmt"

var tablaVector [10]int = [10]int{10, 5} // El vector tiene dimension
var tablaMatriz [10][10]int

func MostrarArreglos() {
	tablaVector[2] = 4
	tablaVector[4] = 6

	for i := 0; i < len(tablaVector); i++ {
		fmt.Println(tablaVector[i])
	}

	tablaMatriz[7][5] = 5
	fmt.Println(tablaMatriz)
}
