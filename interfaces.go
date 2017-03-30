package main

import ("fmt"; "math")

type figura interface {
	area() float64
}

type retangulo struct {
	largura, altura float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func verArea(f figura) {
	fmt.Println(f.area())
}

func main() {
	r := retangulo{largura: 10, altura: 10}
	c := circulo{raio: 5}

	verArea(r) //100
	verArea(c) //78.53981633974483
}

