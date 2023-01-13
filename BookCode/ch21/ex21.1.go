package main

import "fmt"

	var functioncall int = 0

func sum(nums ...int) int{
	sum :=0
	functioncall += 1

	fmt.Printf("[functioncall] : %d nums 타입 %T\n",functioncall,nums)

	for i,v := range nums {
		sum +=v
		fmt.Printf("[%d]%d\n",i,v)
	}
	return sum
}

func main() {

	fmt.Println(sum(1,2,3,4,5))
	fmt.Println(sum(10,20))
	fmt.Println(sum())
}