package main

import "fmt"

func main() {
	var x int8 = 16
	var y int8 = -128
	var z int8 = -1
	var w uint8 = 128

	//오퍼플로우 값은 모두 버려짐
	fmt.Printf("x(16):%08b x <<: %08b x<<2: %d\n", x, x<<2,x<<2)
	fmt.Printf("y(-128):%08b y <<: %08b y<<2: %d\n", uint8(y), uint8(y<<2),uint8(y<<2))
	fmt.Printf("z(-1):%08b z <<: %08b z<<2: %d\n", uint8(z), uint8(z<<2),uint8(z<<2))
	fmt.Printf("w(128):%08b w <<: %08b w<<2: %d\n", uint8(w), uint8(w<<2),uint8(w<<2))
}
