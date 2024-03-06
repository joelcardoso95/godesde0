package model

type Mulher struct {
	Homem
}

func (m *Mulher) Respirar()    { m.respirando = true }
func (m *Mulher) Comer()       { m.comendo = true }
func (m *Mulher) Pensar()      { m.pensando = true }
func (m *Mulher) Sexo() string { return "Mulher" }
