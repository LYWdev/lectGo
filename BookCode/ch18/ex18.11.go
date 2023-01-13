package main

import "fmt"

func main() {
	slice := []int{1,2,3,4,5,6}

	slice = append(slice,0)

	idx :=2

	for i := len(slice)-2; i>=idx; i--{
		slice[i+1]=slice[i]
	}

	fmt.Println(slice)

	slice[idx]=100

	fmt.Println(slice)

//	slice2 := []int{1,2,3,4,5,6}
//
//	slice2 = append(slice2,0)
//
//	fmt.Println("slice2: " ,slice2)
//	for i := idx+1; i<len(slice2)-1; i++{
//	}
//
//	slice2[idx]=100
//
//	fmt.Println("slice2: " ,slice2)

}