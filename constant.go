package main
import ("fmt")

func main(){
// constant can't be changed 

// const pi=3.14
// pi=4
// 	fmt.Println(pi)

// typed constant

   const A int=1
   fmt.Println(A)


// untyped constant

   const B=2
   fmt.Println(B)

// multiple const declaration
const(
	Q=2
	W int=4
	E="Hi"
)
    fmt.Println(Q)
	fmt.Println(W)
	fmt.Println(E)
   
}
