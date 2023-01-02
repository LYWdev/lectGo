package main

import "fmt"

func getMyAGE() int {
	return 22
}

func main() {
	switch age:=getMyAGE(); age{
	case 10:
		fmt.Println("a값은 1")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is ", age)
	}
	//스위치 변수의 범위를 벗어나서 age호출이 안됨
	//fmt.Println("My age is ", age)
}