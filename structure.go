package main
import "fmt"
// type books struct{
// 	title string
// 	author string
// 	subject string
// 	id int

// }
// func main(){
// 	var book1 books //declare book1 of type book
// 	var book2 books // declare book2 of type book
// 	book1.title="go programming"
// 	book1.author="mahesh kumar"
// 	book1.subject="go programming tutorials"
// 	book1.id=4543

// 	book2.title="telecom billing"
// 	book1.author="rajesh kumar"
// 	book1.subject="telecom billing tutorials"
// 	book1.id=7474738
// 	printbook(book1)
// 	printbook(book2)
// }
// 	func printbook(book books){
// 	fmt.Printf("book 1 title:%s\n",book.title)
// 	fmt.Printf("book 1 author:%s\n",book.author)
// 	fmt.Printf("book 1 subject:%s\n",book.subject)
// 	fmt.Printf("book 1 id:%d\n",book.id)
// 	}
type rectangle func(int,int) int
type rectpara struct{
	len int
	bre int
	color string
	rect rectangle
}
func main(){
	result:=rectpara{
		len: 10,
		bre: 84,
		color: "red",
		rect: func(len int, bre int) int{
			return len * bre
		},
	}
	fmt.Println("color:",result.color)
	fmt.Println("area of rect:",result.rect(result.len,result.bre))
}