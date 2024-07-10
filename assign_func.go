// package main
// import "fmt"
// func main(){
// 	test:= func(x int) int{
// 		return * -1
// 	}(8)
// 	fmt.Println(test)
// }


package main
import "fmt"
func calculate(x,y int)(int,int){
	return x+y,x-y
}
func main(){
	sum,diffrence:=calculate(10,20)
	fmt.Printf("sum is %d and diffrence is %d",sum,diffrence)
}