package main
import "fmt"
import "io"
import "os"
func main(){
	var dd int=17
	var mm int=04
	var yy int=2021
	var str string

	str=fmt.Sprintf("%02d-%02d-%04d\n",dd,mm,yy)
	io.WriteString(os.Stdout,str)
}