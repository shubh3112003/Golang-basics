package main
import ("fmt"
"os")
func main(){
	// var a int=20
	// var b int=20
	// if !(a<b){
	// 	fmt.Println("47 is greater than 22")
	// }
	// if true{
	// 	fmt.Println("true")
	// }



	// x:=10;
	// if y:=20;a:=10;y>x{
	// 	fmt.Println("x is smaller than y")
	// }


	fmt.Printf("enter name and age:")
	var name string
	var age int
	if _ , err:=fmt.Scan(&name,&age);err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("name:%s\n",name)
	fmt.Printf("your age:%d\n",age)


}