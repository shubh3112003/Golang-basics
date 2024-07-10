package main
import "fmt"

// // recursive function call to infinity
// func countdown(number int){
// 	fmt.Println(number)
// 	countdown(number-1)
// }
// func main(){
// 	fmt.Println("countdown start")
// 	countdown(3)
// }

func countdown(number int){
	if number>0{
		fmt.Println(number)
		countdown(number-1)
	}else{
		fmt.Println("countdown stops")
	}

}
func main(){
	fmt.Println("countdown start")
	countdown(3)
}
