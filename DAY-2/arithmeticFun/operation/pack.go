package operation

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}
func Div(a, b int) (int, error) {
	if b == 0 {
	return 0,fmt.Errorf("cannot divide with zero")
	}
	return  a/b,nil
}