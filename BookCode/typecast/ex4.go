package main

import "fmt"

func main (){
	// 
	var num1 float64 = 3.5
	var b int =3
	// print 9  큰 타입이 작은 타입으로 전환되면 버림 연산을 한다. 이것은 매우 중요하다.
	// 컴퓨터는 이진 연산이기 때문에 생각보다 더 큰 오류를 발생시킬 수 있다.
	fmt.Println(int(num1)*b)
}