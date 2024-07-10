package main
import ("fmt")
func main(){
	// x:=10
	// y:=34

	// if(x>y){
	// 	fmt.Println("x is greater")
	// }else{
	// 	fmt.Println("y is greater")
	// }



	//  odd or even using if else
	var a int
	fmt.Printf("enter number:")
	fmt.Scanf("%d",&a)
	if(a%2==0){
		fmt.Println("even")
	}else{
		fmt.Println("odd")
	}

	if(a<0){
		var b int=a*(-1)
		fmt.Println(b)
	}else{
		fmt.Println(a)
	}
}