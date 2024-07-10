package main
import "fmt"
func calculate() func() int{
	number:=0
	return func() int{
		number++
		return number
	}
}
func main(){
	num1:=calculate()
	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(calculate())
	fmt.Println(calculate()())

	num2:=calculate()
	fmt.Println(num2())
	fmt.Println(num2())
}