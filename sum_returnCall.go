package main
import "fmt"
func sum(number int) int{
	if number==0{
		fmt.Printf("please enter number")
		return 0
	}else{
		return number+sum(number-1)
	}
	
}
func main () {
	var num=9939
	var result=sum(num)
	fmt.Println(result)
}