package main
import "fmt"
// func calculate() func() int{
// 	num:=1
// 	return func() int{
// 		num=num+2
// 		return num
// 	}
// }
// func main(){
// 	odd:=calculate()
// 	fmt.Println(odd())
// 	fmt.Println((odd()))
// 	fmt.Println(odd())

// 	odd2:=calculate()
// 	fmt.Println(odd2())
// }


func greet() func() string{
	name:="Jatin"
	return func() string{
		name="hi "+ name
		return name
	}
} 
// closure   is a function which returns another function and  the  
// returned function has access to variables from the enclosing  function
func main(){
	message:=greet()
	fmt.Println(message())
}
