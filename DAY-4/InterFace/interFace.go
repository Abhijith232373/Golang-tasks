package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	redius float64
}

func (c Circle) Area()float64 {
	return math.Pi*c.redius*c.redius
}

func (c Circle) Perimeter()float64{
	return 2*math.Pi*c.redius
}

type Reactangle struct{
	lenght,width float64
}

func (r Reactangle)Area()float64  {
	return r.lenght*r.width
	
}
func (r Reactangle)Perimeter()float64{
	return 2*(r.lenght+r.width)
}


func  Results(S shape){
	fmt.Println("Area is:",S.Area())
	fmt.Println("Perimeter is:",S.Perimeter())
	fmt.Println("_________________________-")
}

func main() {
c:=Circle{redius: 5}
r:=Reactangle{lenght: 10,width: 13}

fmt.Println("Circle")
Results(c)
fmt.Println("Reactangle")
Results(r)
}