package main

import "fmt"

func main() {
	slice1	:= []int{1,2,3,4,5}
	slice2 	:= make([]int, len(slice1)) 

	for i, d := range slice1 {
		slice2[i]=d
	}
	``
	fmt.Println(slice1)
	fmt.Println(slice2)
}