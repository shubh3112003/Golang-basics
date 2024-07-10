package main
import("fmt")
func main(){
	var a,b,c int
	fmt.Printf("enter a,b,c:")
	fmt.Scan(&a,&b,&c)

	if a>b && a>c{
		fmt.Println("a is largest")
	}
	if b>a && b>c{
		fmt.Println("b is largest")
	}
	if c>a && c>b{
		fmt.Println("c is largest")
	}
}