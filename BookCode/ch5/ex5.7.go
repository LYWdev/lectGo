package main

import "fmt"

func main() {
	var a int
	var b int

	n,err:= fmt.Scanln(&a, &b)
	if err != nil{
		fmt.Println(n,err)
	}
	// 컴퓨터 내부 표준 입력 스트립이라는 메모리 공간에 FIFO(First In First Out)으로 저장됨.
	// 읽어온 다음 타입을 판단함. 만약에 int타입으로 저정하였는데 Hello라는 문자가 저장되어 있으면 Error리턴

	else{
		fmt.Println(n,a,b)
	}
}
