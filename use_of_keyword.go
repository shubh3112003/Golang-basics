package main
import ("fmt")
func main(){
	// diffrence between var and :=

	var a,b=6, "hello"
	c,d:=5,"hi"
    fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var(
		q int=12
		e int=34
		r string="hr"
	)
	fmt.Println(q,e,r)
	fmt.Printf("%d\n%d\n%s",q,e,r)

}