package main

import "fmt"

type person struct {
	name string
	age  int
}

func (P person) fun() {
	fmt.Printf("my name is %s and im %d years old",P.name,P.age)
}

func main() {
p1:=person{"abhijith",20}

p1.fun()
}