package main

import "fmt"


type myInt int  			// 사용자 지정 별칭 타입

func (num myInt)add(b int)int{ 		//일반  함수
	return int(num) +b
}

func main(){
	var num myInt = 20
	//myInt 타입의 add메소드 호출
	fmt.Println(num.add(30))
	
	var num2 int = 10
	fmt.Println(myInt(num2).add(30))
}