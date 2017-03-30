package main

import "fmt"

type retangulo struct {
	largura, altura int
}

func (r *retangulo) area() int {
	return r.largura * r.altura
}

func main() {
	r := retangulo{largura: 10, altura: 5}

	fmt.Println("a:", r.area()) //50

	rc := r
	rc.altura = 20

	fmt.Println("a:", r.area()) //50

	rp := &r
	rp.altura = 20

	fmt.Println("a:", r.area()) //200
}