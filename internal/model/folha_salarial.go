package model

type FolhaSalarial interface {
	AumentarSalario() float32
	DiminuirSalario() float32
}

func AumentarSalario(f FolhaSalarial) {
	f.AumentarSalario()
}
