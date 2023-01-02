package main

import (
	"fmt"
)

func main() {
	for i:=1; i<7; i++{
		for j:=1; j<7-i; j++ {
			fmt.Print("*")
		} 
			fmt.Println()
	}
}