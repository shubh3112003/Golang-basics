package main
import("fmt")

func main() {
	var name, lname string
	fmt.Print("Enter your full name:")
	fmt.Scanf("%s %s",&name, &lname)
	fmt.Println("Your name:",name,lname)

	var class string
	fmt.Print("enter branch:")
	fmt.Scanln(&class)
	fmt.Println("your class:",class)
}