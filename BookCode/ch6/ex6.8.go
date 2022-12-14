package main
import (
	"fmt"
	"math"
)

func equal(a,b float64) bool {
	return math.Nextafter(a,b) == b
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3
	x,y := 1,2

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b,equal(a+b,c))

    a = 0.00000000004
    b = 0.00000000002
    c = 0.00000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b,c))

	//swap
	fmt.Printf(" BEFORE SWAP  x: %d, y: %d\n", x, y)
	x,y = y,x
	fmt.Printf(" AFTER SWAP  x: %d, y: %d\n", x, y)


}
