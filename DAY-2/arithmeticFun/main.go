package main

import ("fmt"
"arithmeticFun/operation")

func main() {
	a, b := 20,2
	fmt.Println("addition : ",operation.Add(a,b))
	fmt.Println("subtraction : ",operation.Sub(a,b))
	fmt.Println("multiplication : " ,operation.Mul(a,b))
	res,err:=operation.Div(a,b)
	if err!=nil{
		fmt.Println("its error:",err)
	}
	fmt.Println("Divition :",res)

}