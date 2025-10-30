package main

import "fmt"

func main(){
	m:=[]int{10,33,42,55,211,34}

	min:=m[0]
	max:=m[0]

	for _,ans:=range m{
		if ans>max{
			max=ans
		}
		if ans<min{
			min=ans
		}
	}
	fmt.Println("max values is :",max)
	fmt.Println("min values is :",min)
}