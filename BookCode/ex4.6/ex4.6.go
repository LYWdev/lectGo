package main

import "fmt"

//package의 전역변수이다.
//main 패키지에 속해있는 변수이다..
var g int = 20

func main()
{
	//go에서 변수의 범위는 {} 중괄호에 따라 결정된다..
	var m int = 20
	{// s변수ㄴ {} 를 벗나면  scope밖이다.
		var s int = 50
		fmt.Println(m,s,g)
	}
	m = s+20//다라서  s는 알 수 없어서 오류이다.
}