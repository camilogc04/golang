package main

import (
	/* "RampUpFolder/devPath/cursogo/golang/variables"
	"fmt"
	"runtime" */
	/*"RampUpFolder/devPath/cursogo/golang/ejercicios"
	"fmt"*/
	/*"RampUpFolder/devPath/cursogo/golang/files"*/
	//"RampUpFolder/devPath/cursogo/golang/funciones"
	//"RampUpFolder/devPath/cursogo/golang/arreglos_slices"
	//"RampUpFolder/devPath/cursogo/golang/mapas"
	//"RampUpFolder/devPath/cursogo/golang/usuarios"
	/*e "RampUpFolder/devPath/cursogo/golang/ejer_interfaces"
	m "RampUpFolder/devPath/cursogo/golang/modelos"*/
	//df "RampUpFolder/devPath/cursogo/golang/defer_panic"
	//"RampUpFolder/devPath/cursogo/golang/gorutines"
	"RampUpFolder/devPath/cursogo/golang/webserver"
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

	/*numero, mensaje := ejercicios.Ejercicio01("S50")
	JUANJO EL LINDO
		fmt.Println("el numero ", numero, mensaje)*/

	//fmt.Println(ejercicios.Ejercicio02())
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()

	//funciones.Calculos()
	//funciones.Exponencia(2)
	//arreglos_slices.MostrarArreglos()
	//arreglos_slices.MostrarSlices()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//usuarios.AltaUsuario()

	/*Pedro := new(m.Hombre)
	e.HumanoRespirando(Pedro)

	Maria := new(m.Mujer)
	e.HumanoRespirando(Maria)*/

	//df.VemosDefer()
	/*canal1 := make(chan bool)
	go gorutines.MiNombreLento("CAMILO", canal1) // Llamada asincrina gracias al prefijo go
	<-canal1                                     // Esto es un canal que hasta que no retorne un valor, no termina la ejecucion.
	*/
	webserver.MiServidorWeb()
}
