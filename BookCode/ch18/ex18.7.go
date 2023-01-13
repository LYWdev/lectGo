package main

import "fmt"

func main() {
	array := [5]int{1,2,3,4,5}

	fmt.Println("array:",array)

	slice := array[1:4] //slicing
	//cap은 1에서 마지막까지 배열을 사용할 수 있다..
	//따라서 4가 나온다.
	fmt.Println("slice:",slice,len(slice),"cap :",cap(slice))

	array[1] = 10

	fmt.Println("After change second element")
	fmt.Println("array:",array)
	fmt.Println("slice:",slice,len(slice),"cap :",cap(slice))

	slice = append(slice, 500)

	fmt.Println("After append 500")
	fmt.Println("array:",array)
	fmt.Println("slice:",slice,len(slice),"cap :",cap(slice))



}