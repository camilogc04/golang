package usuarios

import (
	"RampUpFolder/devPath/cursogo/golang/modelos"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(modelos.Usuario)
	u.AgregarUsuario(10, "Camilo", time.Now(), true)
	fmt.Println(u)
}
