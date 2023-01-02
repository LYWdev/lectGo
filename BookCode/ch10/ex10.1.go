package main

import "fmt"

func main() {
	a:=4

	switch a {
	case 1:
		fmt.Println("a == 1")
		fmt.Println("a값은 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fmt.Println("a값은 3")
	default:
		fmt.Println("a는 123이 아니다")
	}
}
