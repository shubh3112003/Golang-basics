package main
import "fmt"
import "math"
func main(){
	// integer:=30
	// fmt.Println(integer,&integer)
	// fmt.Printf("%T %T",integer,&integer)


	// number,floating:=238,124.543345533
	// fmt.Printf("\ndefault:%f",floating)
	// fmt.Printf("\n1.2f:%5.2f",floating)
	// fmt.Printf("\n.4f:%.4f",floating)
	// fmt.Printf("\nScientific:%e",floating)
	// fmt.Printf("\ndecimal:%d",number)
	// fmt.Printf("\nbinary:%b",number)
	// fmt.Printf("\noctal:%o",number)
	// fmt.Printf("\nhexadecimal:%X",number)

	// fmt.Printf("%*s\n",40,"text")
	// fmt.Printf("%040s\n","rajesh")

	// var val byte='A'
	// fmt.Printf("character value:%c\n",val)
	// fmt.Printf("ASCII value:%d\n",val)
	// fmt.Printf("type:%T\n",val)

	var val float64=-74.4244
	fmt.Printf("absolute value of %f is %f", val,math.Abs(val))



}