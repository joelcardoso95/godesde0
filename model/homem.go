package model

type Homem struct {
	idade      int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comendo    bool
	vivo       bool
}

func (h *Homem) Respirar()    { h.respirando = true }
func (h *Homem) Comer()       { h.comendo = true }
func (h *Homem) Pensar()      { h.pensando = true }
func (h *Homem) Sexo() string { return "Homem" }
