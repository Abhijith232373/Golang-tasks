package main

import "fmt"

type Student struct{
	Name string
	Age int
}
var Students []Student 


func add(name string,age int){
	Students=append(Students, Student{name,age})
	fmt.Println("Student added")
}

func list(){
	if len(Students)==0{
		fmt.Println("not founded")
		return
	}
	
	for i,s:=range Students{
		fmt.Printf("%d.Name:%s-Age:%d",i,s.Name,s.Age)
	}
}
func Update(index int,newName string,newAge int){
	if index <0||index>=len(Students){
		fmt.Println("invaild")
		return
	}
	Students[index].Name=newName
	Students[index].Age=newAge
	fmt.Println("Ubdated")
}
func Delete(index int){
	if index<0||index>=len(Students){
		fmt.Println("invaild")
		return
	}
	Students=append(Students[:index],Students[index+1:]...)
	fmt.Println("Deleted")
}

func main() {
for{
	fmt.Println("\n STUDENT DATAS")
	fmt.Println("\n 1.ADD")
	fmt.Println("\n 2.List")
	fmt.Println("\n 3.Update")
	fmt.Println("\n 4.Delete")
	fmt.Println("\n 5.Exit")
	var Choise int
	fmt.Println("Enter any Choise")
	fmt.Scan(&Choise)

	switch Choise{
	case 1:
		var name string
		var age int
		fmt.Println("Enter the name:")
		fmt.Scan(&name)
		fmt.Println("Enter the age:")
		fmt.Scan(&age)
		add(name,age)

	case 2:
		list()
	case 3:
		var index,age int
		var name string
		list()
		fmt.Println("Enter the index:")
		fmt.Scan(&index)
		fmt.Println("change name:")
		fmt.Scan(&name)
		fmt.Println("change age:")
		fmt.Scan(&age)
		Update(index,name,age)

	case 4:
		var index int
		list()
		fmt.Println("Enter the index:")
		fmt.Scan(&index)
		Delete(index)

	case 5:
		fmt.Println("Exit")
		return
	default:
		fmt.Println("invaild")
		return
	}
}
}