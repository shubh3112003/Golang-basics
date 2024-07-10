package main
import "fmt"
type vertex struct{
	lat, Long float64
}
var m=map[string]vertex{
	"Bangalore":vertex{13.0589, 77.5917},
	"Delhi":vertex{22.3221, 46.7474},
}
func main() {
	fmt.Println(m)
}