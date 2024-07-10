package main
import "fmt"

// function call

// func main(){
// 	addNumber()
// }
// func addNumber(){
// 	n1:=10
// 	n2:=45
// 	sum:=n1+n2
// 	fmt.Printf("sum:%d",sum)
// }



// func addNumber(n1 int,n2 int){
// 	sum:=n1+n2
// 	fmt.Printf("sum:%d",sum)
// }
// func main(){
// 	var x,y int
// 	fmt.Printf("enter two values:")
// 	fmt.Scanf("%d %d",&x,&y)
// 	addNumber(x,y)
// }



// func addNumber(n1 int,n2 int) int{ // int when you want to use return function
// 	sum:=n1+n2
// 	return sum

// 	fmt.Printf("after return")
// }
// func main(){
// 	fmt.Printf("sum:%d",addNumber(34,55))
// }


// //return multiple values
// func addNumber(n1 int,n2 int) (int,int){
// 	sum:=n1+n2
// 	diff:=n1-n2
// 	return sum,diff
// }
// func main(){
// 	sum,diffrence:=addNumber(23,12)
// 	sum1,diff1:=addNumber(74,55)
// 	sum2,diff2:=addNumber(56,22)
// 	fmt.Printf("sum:%d\n diffrence:%d\n",sum,diffrence)
// 	fmt.Printf("sum:%d\n diffrence:%d\n",sum1,diff1)
// 	fmt.Printf("sum:%d\n diffrence:%d\n",sum2,diff2)
// }



// // local variable
// func main(){
// sum:=addNumber()
// fmt.Printf("sum:%d",sum)
// }
// func addNumber() int{
// 	var sum int
// 	sum=4+5
// 	return sum
// }


func countdown(number int){
	fmt.Println(number)
	countdown(number-1)
}
func main(){
	fmt.Println("countdown start")
	countdown(3)
}
