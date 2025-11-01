package main

import "fmt"



func point(num *int){
	*num=*num+10
}
func main() {
	x := 5
	fmt.Println("Before modification : ",x)

	point(&x)
	fmt.Println("After modifiaction : ",x)

}