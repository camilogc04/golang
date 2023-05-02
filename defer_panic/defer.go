package defer_panic

import (
	"fmt"
)

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el ultimo mensaje") // ESTA INSTRUCCION SIEMPRE DE EJECUTA AL FINAL; INDEPENDIENTE DE ERRORES (EJ: DEFER DE CLOSE A BD)
	fmt.Println("Este es un segundo mensaje")
}
