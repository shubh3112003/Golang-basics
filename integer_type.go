package main
import ("fmt")

func main(){
	// range for uint8 is 0 to 255
	// var a uint8=45
	// fmt.Println(a)


	// //inferred and static input
	// var student string="John" //type is static
	// var student2="Jane" //type is inferred
	// x:=2 //type is inferred

	// fmt.Printf("%s\n",student)
	// fmt.Printf("%s\n",student2)
	// fmt.Printf("%d",x)


	// //inference type using %T
	// var x float64=20.0
	// fmt.Printf("type of variable:%T",x)

	// variable declaration without initialization
	var e string
	var t bool
	fmt.Println(e)
    fmt.Println(t)

	var a,b,c=1,2,"abc"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("type is:%T",c)
	

}