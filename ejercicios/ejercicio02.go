package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n int
var err error
var texto string

func Ejercicio02() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese un numero entero")
	if scanner.Scan() {
		n, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Dato incorrecto" + err.Error())
		}
		for i := 1; i < 11; i++ {
			texto += fmt.Sprintln(n, "x", i, "=", n*i)
		}

	}

	return texto

}
