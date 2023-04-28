package ejercicios

import (
	"strconv"
)

func Ejercicio01(valor string) (int, string) {
	var valorConvertido int
	valorConvertido, err := strconv.Atoi(valor)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if valorConvertido > 100 {
		return valorConvertido, "es mayor a 100"
	}
	return valorConvertido, "es menor a 100"
}
