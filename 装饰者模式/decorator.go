package main

import (
	"fmt"
)

func main() {
	ramen := Ramen{name: "ramen", price: 10}

	fmt.Println(ramen.Description())
	fmt.Println(ramen.Price())

	egg := Egg{noddles: ramen, name: "egg", price: 2}

	fmt.Println(egg.Description())
	fmt.Println(egg.Price())

	sausage := Sausage{noddles: egg, name: "sausage", price: 3}
	fmt.Println(sausage.Description())
	fmt.Println(sausage.Price())

	egg2 := Egg{noddles: egg, name: "egg", price: 2}

	fmt.Println(egg2.Description())
	fmt.Println(egg2.Price())
}

type Noddles interface {
	Description() string
	Price() float32
}

type Ramen struct {
	name  string
	price float32
}

func (p Ramen) Description() string {
	return p.name
}

func (p Ramen) Price() float32 {
	return p.price
}

type Egg struct {
	noddles Noddles
	name    string
	price   float32
}

func (p Egg) SetNoddles(noddles Noddles) {
	p.noddles = noddles
}

func (p Egg) Description() string {
	return p.noddles.Description() + "+" + p.name
}

func (p Egg) Price() float32 {
	return p.noddles.Price() + p.price
}

type Sausage struct {
	noddles Noddles
	name    string
	price   float32
}

func (p Sausage) SetNoddles(noddles Noddles) {
	p.noddles = noddles
}

func (p Sausage) Description() string {
	return p.noddles.Description() + "+" + p.name
}

func (p Sausage) Price() float32 {
	return p.noddles.Price() + p.price
}
