package main

import "fmt"

func main() {
	var n1 int
	var operation string
	var n2 int

	fmt.Print("Enter the first number : ")
	fmt.Scan(&n1)
	fmt.Print("Operation : ")
	fmt.Scan(&operation)
	fmt.Print("Enter Second number : ")
	fmt.Scan(&n2)

	switch operation {
	case "+":
		fmt.Printf("%v + %v = %v",n1,n2,n1+n2)
	case "-":
		fmt.Printf("%v - %v = %v",n1,n2,n1-n2)
	case "*":
		fmt.Printf("%v * %v = %v ",n1,n2,n1*n2)
	}
}