package main
import "fmt"
func main(){
	//// converaion of int to uint

	// var a int=10
	// var b uint=0
	// b=uint(a)
	// fmt.Printf("a=%d, b=%d\n",a,b)

	// //math funtion of getting square
	// var num int=66
	// var result float32
	// result=float32(math.Sqrt(float64(num)))
	// fmt.Printf("result:%f",result)



	var x int=42
	var y float64=float64(x)
	var z uint=uint(y)

	fmt.Printf("\nnumber:%d of type %T",x,x)
	fmt.Printf("\nnumber:%f of type %T",y,y)
	fmt.Printf("\nnumber:%d of type %T",z,z)




}