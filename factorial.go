package main
import "fmt"

func factorial(num int)int{
	if num==0{
		return 0
	}else{
		return num+factorial(num-1)
	}
}
func main(){
	var number=3
	var result=factorial(number)
	fmt.Println(result)
}