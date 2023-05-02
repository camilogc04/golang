package modelos

import "time"

type Usuario struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Estado    bool
}

// Importante el * porque es lo que indica que haga referencia a la direccion de memoria y no un objeto nuevo diferente
func (us *Usuario) AgregarUsuario(id int, name string, createdAt time.Time, estado bool) {
	us.Id = id
	us.Name = name
	us.CreatedAt = createdAt
	us.Estado = estado
}
