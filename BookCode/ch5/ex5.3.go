package main

import "fmt"

func main() {
	var a = 324.13455
	var c = 3.14

	fmt.Printf("%08.2f\n",a)//최소 너비 8, 소수점이하 2자리까지 표현, 나머진 0채우기
	fmt.Printf("%08.2g\n",a)//최소 너비 8, 총 숫자 2자리까지 표현, 나머진 0채우기
	fmt.Printf("%8.5g\n",a)//최소 너비 8, 총 숫자 2자리까지 표현
	fmt.Printf("%f",c)//소숫점 이하 6자리까지 출력
}
