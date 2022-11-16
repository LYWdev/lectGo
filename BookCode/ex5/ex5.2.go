package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20 
	var f float64 = 327999438743.3887

	//
	fmt.Print("[1]   print : a:",a, "b:",b, "f",f, "\n")

	//Println은 개행을 삽입하여 출력한다.
	fmt.Println("[2] println : a:",a, "b:",b, "f",f)

	//Printf는 C언어 처럼 출력 서식을 정의하여 출력한다.
	fmt.Printf("[3]  printf : a:,%d, b:,%d, f,%f \n", a,b,f)
}