package main
import "fmt"
var sum=0
func find(num int)int{
	square:=num*num
	return square
}
func main(){
	// var greet=func(){
	// 	fmt.Println("hello world")
	// }
	// var welcome=greet
	// welcome()

	// var sum=func(num int,num2 int) int{
	// 	result:=num+num2
	// 	return result
	// }

	// res:=sum(5,6)
	// fmt.Printf("sum:%d",res)

	sum:=func(num1 int,num2 int)int{
		return num1+num2
	}

	result:=find(sum(2,2))
	fmt.Println("square:",result)
}