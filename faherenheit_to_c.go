package main
import "fmt"
func main(){
	var ftemp float64=0
	var ctemp float64=0

	fmt.Printf("enter faherenheit:%f\n",ftemp)
	fmt.Scanf("%f",&ftemp)
	ctemp=(ftemp-32)/1.8
	fmt.Printf("temreture:%f",ctemp)
}