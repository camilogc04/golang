package arreglos_slices

import (
	"fmt"
)

// LOS SLICES PUEDEN AUMENTAR SU CAPACIDAD EN TIEMPO DE EJECICION
var tablaSlice []int = []int{1, 2, 3}
var arreglo [15]int = [15]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 200, 300, 400, 500, 600}

func MostrarSlices() {
	//fmt.Println(tablaSlice)

	porcion := arreglo[3:]    //slice creado desde un vector desde la posicion 3 hasta el final
	porcion2 := arreglo[:5]   //slice creado desde la primnera posicion hasta la posicion 5
	porcion3 := arreglo[5:10] //slice creado desde la posicion 5 hasta la posicion 10
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	//          HACER (SLICE,5 elementos, 20 de capacidad)
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d/n", len(nums), cap(nums))
}
