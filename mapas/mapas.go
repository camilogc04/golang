package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F."
	paises["Colombia"] = "Bogota"
	//fmt.Println(paises)
	//fmt.Println(paises["Colombia"])

	campeonato := map[string]int{
		"Millonarios": 1,
		"Nacional":    2,
		"DIM":         3,
		"Santa Fe":    4,
		"Cali":        5,
		"America":     6,
		"Junior":      7,
		"Once Caldas": 8,
	}

	fmt.Println(campeonato)

	/*for equipo, posicion := range campeonato {
		fmt.Printf("Equipo: %s esta en la posicion %d \n", equipo, posicion)
	}*/
	delete(campeonato, "Once Caldas")
	fmt.Println(campeonato)

	posicion, existe := campeonato["Once Caldas"]
	fmt.Printf("La posicion capturada es %d, y el equipo existe = %t \n", posicion, existe)
}
