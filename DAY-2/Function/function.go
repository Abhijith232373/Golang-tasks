package main

import "fmt"

func fun(a, b int) (sum, diff, product, quotient int) {
	sum = a + b
	diff = a - b
	product = a * b
	quotient = a / b
	return
}

func main() {
	sum, diff, product, quotient := fun(10, 2)
	fmt.Println("sum is = ",sum)
	fmt.Println("diff is = ",diff)
	fmt.Println("Product is = ",product)
	fmt.Println("quotient is = ",quotient)
}