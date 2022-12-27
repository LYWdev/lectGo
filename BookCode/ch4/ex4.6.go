package main

import "fmt"

//go에서는 {}중괄호를 기준으로 모든 변수의 범위가 결정된다

//전역변수
var g int = 10

func main() {
	//지역변수
	var m int = 20
	{
		//지역변수
		var s int =50
		fmt.Println(m,s,g)
	}

}
