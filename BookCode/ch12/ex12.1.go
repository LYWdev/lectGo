package main

import "fmt"

func main() {
	
	//변수 선언 var 변수명 [개수 ]타입
	var t[5] float64 =[5] float64{24.0, 25.9, 27.8, 26.9, 26.2}
	for i :=0; i<5; i++{
		fmt.Println(t[i])
	}
}
