package main
import "fmt"
func main(){
	// // for loop working as do...while loop
	// number:=1
	// for{
	// 	fmt.Printf("%d\n",number);
	// 	number++
	// 	if number>5{
	// 		break;
	// 	}
	// }

	// multiplication using while loop
	multiplier:=1
	var num int
	fmt.Println("enter number:")
	fmt.Scan(&num)
	for multiplier<=10{
		product:=num*multiplier
		fmt.Printf("%d*%d=%d\n",num,multiplier,product)
		multiplier++
	}
}