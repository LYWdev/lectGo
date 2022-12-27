// 프로그램에서 변수는 메모리의 공간을 의미한다.
package main

import "fmt"

func main(){
// go언어는 
// 변수 선언키워드 변수명 타입 초기값으로 선언한다. 

	var a int = 10
	var msg string = "Hello Variable"

	//var a = 20
	// a redeclared in this block 오류로 실행 못함.
	
	msg = "Good Morning"
	fmt.Println(msg,a)
}