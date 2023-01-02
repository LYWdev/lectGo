package main

import "fmt"

func main() {
	a:= 2
	switch a {
	case 1 :
		fmt.Println("Go언어는 자동으로 Break")
	case 2 :
		fmt.Println("계속 진행하고 싶을 때 fallthrough",a)
		fallthrough
	case 3 :
		fmt.Println("계속 진행하고 싶을 때 fallthrough",a)


	}
}


