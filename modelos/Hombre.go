package modelos

type Hombre struct {
	edad       int
	altura     float32
	peso       float32
	pensando   bool
	respirando bool
	comiendo   bool
	vivo       bool
}

func (h *Hombre) Respirar() {
	h.respirando = true
}

func (h *Hombre) Pensar() {
	h.comiendo = true
}

func (h *Hombre) Comer() {
	h.comiendo = true
}

func (h *Hombre) Sexo() string {
	return "Hombre"
}
