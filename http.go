package  main
import("fmt" 
"math")
type shape interface{
	area() float64
}
type circle struct{
	x,y,radius float64
}
type rect struct{
	width,height float64
}
func  (circle circle) area()float64{
	return  math.Pi * (circle.radius*circle.radius)
}
func  (rect rect) area()float64 {
	return rect.width * rect.height
}
func calculateArea(s shape){
	return  s.area()
}
func main(){
	c := circle{ x:0 , y:-3 , radius :2 }
	r := rect{ width:10	, height:5 }

	fmt.Println("circle: ",calculateArea(c))
	fmt.Println("Rectangle: ",calculateArea(r))
}