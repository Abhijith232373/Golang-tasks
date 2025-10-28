package main 
 import "fmt"
func main(){
	var n int
	fmt.Print("Enter the number : ")
	fmt.Scan(&n)
	 if  n%2==0{
		fmt.Printf("%v is even",n)
	 }else{
		fmt.Printf("%v is odd",n)
	 }
	
}