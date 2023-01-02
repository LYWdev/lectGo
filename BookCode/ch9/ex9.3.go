package main

import "fmt"

func main() {

	var age = 17
	if age >= 10 && age <=15{
		fmt.Println("10보다 크고 20보다 작음.")
	} else if age > 30 || age < 20{
		fmt.Println("you are not 20's")
	}else {
		fmt.Println("Best Age")
	}
}