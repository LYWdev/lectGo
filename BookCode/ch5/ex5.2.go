package main

import "fmt"

func main () {
	// 최소 출력 너비 지정
	var a = 123
	var b = 456
	var c = 123456789
	//서식의 %와 d(문자 타입 지칭)문자 사이에 숫자를 넣음 해당 만큼 칸이 생성됨
	fmt.Printf("%5d, %5d\n", a,b)
	//너비 앞에 0을 붙이면 빈자리를 0으로 채움
	fmt.Printf("%5d, %5d\n", a,b)
	//왼쪽 정렬하기 숫자에 -를 붙이면 왼쪽 정렬됨.
	fmt.Printf("%05d, %05d\n", a,b)

	fmt.Printf("%5d, %5d\n", c,c)
	fmt.Printf("%05d, %05d\n", c,c)
	fmt.Printf("%-5d, %-5d\n", c,c)
}