package main
import "fmt"
func main(){
	var radius float32=0
	var area float32=0
	var perimeter float32=0
	const PI=3.14
	fmt.Printf("enter radius:");
	fmt.Scanf("%f",&radius)

	area=PI*radius*radius
	perimeter=2*PI*radius
	fmt.Printf("area=%.2f",area)
	fmt.Printf("perimeter=%f",perimeter)

}