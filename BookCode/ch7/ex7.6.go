package main

import "fmt"

//재귀 호출 탈출 조건

func prinNo(n int){
	if n==0{
		return
	}
	fmt.Println(n)
	prinNo(n-1)
	fmt.Println("After",n)
}

func main(){
	prinNo(3)
}
