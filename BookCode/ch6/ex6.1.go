package main

import "fmt"

func main(){
	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	// 정수 연산 정수는 정수
	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	// 실수 연산 실수 는 실수다
	fmt.Println("s * t = ", s*t)
	fmt.Println("s / t = ", s/t)
}
