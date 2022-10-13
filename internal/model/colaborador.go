package model

type Colaborador struct {
	Nome    string
	Salario float32
}

func (c *Colaborador) AumentarSalario() float32 {
	c.Salario += 2000
	return c.Salario
}

func (c *Colaborador) DiminuirSalario() float32 {
	c.Salario -= 2000
	return c.Salario
}
