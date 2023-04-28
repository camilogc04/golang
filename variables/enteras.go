package variables

import "fmt"

func MostrarEnteros() {
	var intComun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intComun = ", intComun)
	fmt.Println("int32 = ", intde32)
	fmt.Println("int64 = ", intde64)
}
