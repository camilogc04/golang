package files

import (
	"RampUpFolder/devPath/cursogo/golang/ejercicios"
	"fmt"
	"io/ioutil"
	"os"
)

var fileName = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Ejercicio02()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo", err)
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Ejercicio02()
	if !Append(fileName, texto) {
		fmt.Println("Error al concatenar  archivo")
	}

}

func Append(fileName string, texto string) bool {
	archivo, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo" + err.Error())
		return false
	}

	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error al concatenar archivo" + err.Error())
		return false
	}
	archivo.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}
	fmt.Println(string(archivo))
}
