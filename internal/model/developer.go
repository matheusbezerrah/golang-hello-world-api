package model

type Developer struct {
	Nome    string
	Skills  []Skill
	Salario float32
}

type Skill struct {
	Tecnologia string
	Nome       string
}

func (d *Developer) AumentarSalario() float32 {
	d.Salario += 5000
	return d.Salario
}

func (d *Developer) DiminuirSalario() float32 {
	d.Salario -= 5000
	return d.Salario
}
