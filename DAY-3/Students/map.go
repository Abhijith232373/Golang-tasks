package main

import "fmt"

func main() {
	student := make(map[string]string)
	student["abhijith"] = "S"
	student["sanal"] = "A"
	student["anandhu"] = "A"
	student["vishnu"] = "B"
	fmt.Println(student)


	student["vishnu"] = "A"
	fmt.Println(student)

	student["abin"]="A+"
	fmt.Println(student)

	delete(student,"abhijith")
	fmt.Println(student)

	fmt.Print(student["abin"])
}