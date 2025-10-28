package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number : ")
	fmt.Scan(&n)

	a,b:=0,1

	fmt.Println("Series : ")
	for i :=0;i<n;i++{
		fmt.Print(a," ")
		a,b =b,a+b
	}
}