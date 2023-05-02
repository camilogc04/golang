package modelos

type Mujer struct {
	Hombre //HERENCIA. LA MUJER HEREDA LAS PROPIEDADES DEL HOMBRE
}

func (m *Mujer) Respirar() {
	m.respirando = true
}

func (m *Mujer) Pensar() {
	m.comiendo = true
}

func (m *Mujer) Comer() {
	m.comiendo = true
}

func (m *Mujer) Sexo() string {
	return "Mujer"
}
