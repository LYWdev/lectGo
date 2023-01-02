package main

import (
	"fmt"
)

func main() {
	dan := 2
	b := 2
	for {
		for {
			fmt.Printf("%d * %d =%d \n", dan, b, dan*b)
			b++
			if b == 10 {
				break
			}
		}
		b=1
		if  dan == 10 {
			break
		}
		dan++
}
	fmt.Println("for exit")
}