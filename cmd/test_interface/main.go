package main

import (
	"fmt"
	"golang-hello-world-api/internal/model"
)

func main() {
	c := model.Colaborador{Nome: "Matheus", Salario: 1000}
	fmt.Println(c)
	model.AumentarSalario(&c)
	fmt.Println(c)

	a := model.Analista{Nome: "Ana", Salario: 500}
	fmt.Println(a)
	model.AumentarSalario(&a)
	fmt.Println(a)

	d := model.Developer{Nome: "João", Salario: 500}
	fmt.Println(d)
	model.AumentarSalario(&d)
	fmt.Println(d)

}
