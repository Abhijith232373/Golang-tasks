package main

import (
	"fmt"
)

func err(a, b int) (int, error) {
	if b == 0 {
		return 0,fmt.Errorf("its zero change number")
	}
	return  a/b,nil


}
func main() {
	res,er:=err(10,0)
	if er!=nil{
		fmt.Println("Error : ",er)
		return
	}
	fmt.Println("success : ",res)
}