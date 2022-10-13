package model

type Analista struct {
	Nome      string
	Graduacao string
	Salario   float32
}

func (a *Analista) AumentarSalario() float32 {
	a.Salario += 500
	return a.Salario
}

func (a *Analista) DiminuirSalario() float32 {
	a.Salario -= 500
	return a.Salario
}
